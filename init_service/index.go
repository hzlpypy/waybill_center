package init_service

import (
	"context"
	"github.com/hzlpypy/common/rabbitmq/topic"
	"github.com/hzlpypy/waybill_center/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Init struct {
	db *gorm.DB
	l  *logrus.Logger
}

func NewInit(db *gorm.DB, l *logrus.Logger) (*Init, error) {
	// init conf
	return &Init{
		db: db,
		l:  l,
	}, nil
}

type InitService interface {
	InitVCTable()
	CreateExchangeAndBindQueue(ctx context.Context, req *topic.TopicReq) error
}

var _ InitService = (*Init)(nil)

func (i Init) InitVCTable() {
	tables := []interface{}{
		&model.Waybill{},
	}
	var createChan = make(chan error, len(tables))
	for _, table := range tables {
		go i.checkAndCreateTable(table, createChan)
	}
	var errors []error
	for range tables {
		err := <-createChan
		if err != nil {
			errors = append(errors, err)
		}
	}
	close(createChan)
	if len(errors) != 0 {
		i.l.Error(errors)
	}
}

func (i *Init) checkAndCreateTable(table interface{}, createChan chan error) {
	isExist := i.db.Migrator().HasTable(table)
	if !isExist {
		log.Printf("create table: %v", reflect.TypeOf(table))
		err := i.db.Migrator().CreateTable(table)
		if err != nil {
			createChan <- err
		} else {
			createChan <- nil
		}
		return
	}
	createChan <- nil
	return

}

// TODO：迁移到name_service
func (i *Init) CreateExchangeAndBindQueue(c context.Context, req *topic.TopicReq) error {
	ti, err := topic.NewTopicReq(c, req)
	if err != nil {
		i.l.WithField("sendMsg", "amqp NewTopicReq error").Error(err)
		return err
	}
	err = ti.CreateExchange()
	if err != nil {
		i.l.WithField("sendMsg", "amqp CreateExchange error").Error(err)
		return err
	}
	err = ti.QueueDeclareAndBindRoutingKey()
	if err != nil {
		i.l.WithField("sendMsg", "amqp QueueDeclareAndBindRoutingKey error").Error(err)
		return err
	}
	return nil
}
