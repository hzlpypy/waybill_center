package router

import (
	"context"
	"fmt"
	"github.com/hzlpypy/common/rabbitmq/topic"
	"github.com/hzlpypy/waybill_center/init_service"
	"github.com/hzlpypy/waybill_center/internal/receive"
	protos "github.com/hzlpypy/waybill_center/proto_info/protos"
	"github.com/hzlpypy/waybill_center/waybill_server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"net"
	"strconv"
)

type initServiceConfig struct {
	grpcPort         int
	initService      init_service.InitService
	topicReqConsumer *topic.TopicReq
	topicReqDead     *topic.TopicReq
	l                *logrus.Logger
	db               *gorm.DB
	ctx              context.Context
}

func NewServiceConfig(grpcPort int, i init_service.InitService, topicReqConsumer, topicReqDead *topic.TopicReq, db *gorm.DB, ctx context.Context, l *logrus.Logger) *initServiceConfig {
	return &initServiceConfig{
		grpcPort:         grpcPort,
		initService:      i,
		topicReqConsumer: topicReqConsumer,
		topicReqDead:     topicReqDead,
		l:                l,
		db:               db,
		ctx:              ctx,
	}
}

//RunGrpcServer creates a gRPC server which has no service registered and has not
func (isc *initServiceConfig) RunGrpcServer() error {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))

	protos.RegisterWaybillCenterServer(grpcServer, &waybill_server.WaybillServer{})
	topicReqConsumer := isc.topicReqConsumer
	topicReqDead := isc.topicReqDead
	// init service
	isc.initService.InitVCTable()
	// 创建死信队列
	err := isc.initService.CreateExchangeAndBindQueue(isc.ctx, topicReqDead)
	if err != nil {
		return err
	}
	// 创建订单接受交换机和队列并绑定死信队列
	err = isc.initService.CreateExchangeAndBindQueue(isc.ctx, topicReqConsumer)
	if err != nil {
		return err
	}
	// 开启监听队列
	go receive.ReceiveConsumer(topicReqConsumer.Conn, topicReqConsumer.Queue.QueueName, isc.l, isc.db)
	go receive.ReceiveDead(topicReqDead.Conn, topicReqDead.Queue.QueueName, isc.l, isc.db)
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(isc.grpcPort))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("User center service has started=> grpc:127.0.0.1:%d", isc.grpcPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("failed to grpc-serve:%v", err)
		return err
	}
	return nil
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(500, fmt.Sprintf("panic: %v", r))
			log.Printf("%v", err)
		}
	}()
	logReqUri(ctx, info.FullMethod)
	return handler(ctx, req)
}

func logReqUri(ctx context.Context, reqUri string) {
	var ip string
	md, _ := metadata.FromIncomingContext(ctx)
	contentType := md.Get("content-type")[0]
	if val, ok := md[":authority"]; ok {
		ip = val[0]
		p, ok := peer.FromContext(ctx)
		if ok {
			log.Printf("%v -> %v%v [%v]", p.Addr.String(), ip, reqUri, contentType)
		} else {
			log.Printf("ip is none: %v", ctx)
		}
	} else {
		log.Printf("ip is none: %v", ctx)
	}

}
