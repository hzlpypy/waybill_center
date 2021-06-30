package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hzlpypy/common/clog"
	hm "github.com/hzlpypy/common/databases/mysql"
	v2 "github.com/hzlpypy/common/databases/mysql/v2"
	"github.com/hzlpypy/common/name_service"
	"github.com/hzlpypy/common/rabbitmq/topic"
	"github.com/hzlpypy/common/utils"
	"github.com/hzlpypy/grpc-lb/common"
	"github.com/hzlpypy/grpc-lb/registry"
	etcd3 "github.com/hzlpypy/grpc-lb/registry/etcd3"
	"github.com/hzlpypy/waybill_center/cmd/config"
	"github.com/hzlpypy/waybill_center/init_service"
	"github.com/hzlpypy/waybill_center/router"
	"github.com/ozonru/etcd/v3/clientv3"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var nodeID = flag.String("node", "node1", "node ID")

func main() {
	flag.Parse()
	cf := config.NewConfig()
	// local ip
	ip, _ := utils.ExternalIP()
	// register to etcd
	nameServiceMap, err := name_service.GetNameService()
	if err != nil {
		log.Fatal(err)
	}
	serviceName := nameServiceMap[cf.Server.Name]
	lbConfig := &etcd3.Config{
		EtcdConfig: clientv3.Config{
			Endpoints:   []string{fmt.Sprintf("http://%s:%d", cf.Etcd.Ip, cf.Etcd.Port)},
			DialTimeout: 5 * time.Second,
		},
		RegistryDir: cf.Server.Name,
		Ttl:         time.Duration(cf.Etcd.Ttl) * time.Second,
	}
	register, _ := etcd3.NewRegistrar(lbConfig)
	service := &registry.ServiceInfo{
		InstanceId: *nodeID,
		Name:       serviceName,
		Version:    "1.0",
		Address:    fmt.Sprintf("%s:%d", ip.String(), cf.Server.Port),
		Metadata:   metadata.Pairs(common.WeightKey, "1"),
	}
	err = register.Register(service)
	if err != nil {
		log.Fatal(err)
	}
	// log
	logPathMap := map[string]string{"access": cf.Log.AccessPath, "error": cf.Log.ErrorPath}
	logCfg := &clog.Cfg{}
	for name, path := range logPathMap {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			file, err = os.Create(path)
			if err != nil {
				log.Fatal(err)
			}
		}
		defer file.Close()
		logCfg.CfgFiles = append(logCfg.CfgFiles, &clog.CfgFile{
			Name: name,
			File: file,
		})
	}
	l, err := clog.Init(logCfg)
	if err != nil {
		log.Fatal(err)
	}
	//accessLog := *l.Access()
	errorLog := *l.Error()

	// conn rabbitmq || mysql
	w := etcd3.NewWatcher(cf.Rabbitmq.Name, register.Etcd3Client)
	addressList := w.GetAllAddresses()

	mysqlAddress := etcd3.NewWatcher(cf.Mysql.Name, register.Etcd3Client).GetAllAddresses()

	// rabbitMq 连接池
	rbConnPools := []*amqp.Connection{}
	for _, address := range addressList {
		addr := address.Addr
		conn, err := amqp.Dial(fmt.Sprintf("amqp://admin:admin@%s", addr))
		if err != nil {
			errorLog.Errorf("Failed to connect to RabbitMQ, err=%v", err)
			continue
		}
		defer conn.Close()
		rbConnPools = append(rbConnPools, conn)
	}
	// db
	mysqlCf := cf.Mysql
	var debug bool
	logLevel := logger.Error
	if cf.Model == "debug" {
		debug = true
		logLevel = logger.Info
	}
	// todo 暂时只取第一个地址
	spMA := strings.Split(mysqlAddress[0].Addr, ":")
	host := spMA[0]
	port, _ := strconv.Atoi(spMA[1])
	db, err := v2.NewDbConnection(&hm.Config{
		Username:                                 mysqlCf.User,
		Password:                                 mysqlCf.Pwd,
		DBName:                                   mysqlCf.DBName,
		Host:                                     host,
		Port:                                     port,
		Charset:                                  mysqlCf.Charset,
		Debug:                                    debug,
		ConnMaxLifetime:                          time.Duration(mysqlCf.ConnMaxLifetime),
		MaxIdleConns:                             mysqlCf.MaxIdleConns,
		MaxOpenConns:                             mysqlCf.MaxOpenConns,
		LogLevel:                                 logLevel,
		DisableForeignKeyConstraintWhenMigrating: mysqlCf.DisableForeignKeyConstraintWhenMigrating,
	})
	if err != nil {
		log.Fatal(err)
	}
	// grpc
	i, err := init_service.NewInit(db, &errorLog)
	if err != nil {
		log.Fatal(err)
	}
	conn := rbConnPools[0]
	ch, _ := conn.Channel()
	occ := cf.OrderCenter.Consumer
	ocd := cf.OrderCenter.Dead
	ctx := context.Background()
	defer ctx.Done()
	err = router.NewServiceConfig(cf.Server.Port, i, &topic.TopicReq{
		Conn:         rbConnPools[0],
		Ch:           ch,
		ExchangeName: occ.Exchange,
		ExchangeType: occ.ExchangeType,
		Durable:      true,
		ContentType:  occ.ContentType,
		RoutingKey:   occ.RoutingKey,
		Queue: &topic.Queue{
			QueueName: occ.Queue,
			QueueDeclareMap: map[string]interface{}{
				"x-max-length":              10,
				"x-dead-letter-exchange":    ocd.Exchange,
				"x-dead-letter-routing-key": ocd.RoutingKey,
			},
		},
	}, &topic.TopicReq{
		Conn:         rbConnPools[0],
		Ch:           ch,
		ExchangeName: ocd.Exchange,
		ExchangeType: ocd.ExchangeType,
		Durable:      true,
		ContentType:  ocd.ContentType,
		RoutingKey:   ocd.RoutingKey,
		Queue: &topic.Queue{
			QueueName:       ocd.Queue,
			QueueDeclareMap: map[string]interface{}{"x-max-length": 10},
		},
	}, db, ctx, &errorLog).RunGrpcServer()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
