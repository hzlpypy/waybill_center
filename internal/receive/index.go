package receive

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	oc "github.com/hzlpypy/common/model/order_center"
	wc "github.com/hzlpypy/common/model/waybill_center"
	"github.com/hzlpypy/common/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
	"time"
)

var sendCountMap = make(map[string]int)

func getMsgs(queue string, conn *amqp.Connection, l *logrus.Logger) (*amqp.Channel, <-chan amqp.Delivery, error) {
	ch, err := conn.Channel()
	if err != nil {
		l.Errorf("make conn chan error, err=%v", err)
		return nil, nil, err
	}
	msgs, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		l.Errorf("QueueBind error,err=%v", err)
		return nil, nil, err
	}
	return ch, msgs, nil
}
func ReceiveConsumer(conn *amqp.Connection, queue string, l *logrus.Logger, db *gorm.DB) {
	ch, msgs, err := getMsgs(queue, conn, l)
	if err != nil {
		l.Errorf("make conn chan error, err=%v", err)
		return
	}
	defer ch.Close()
	for msg := range msgs {
		res := &oc.OrderMsg{}
		err := json.Unmarshal(msg.Body, &res)
		if err != nil {
			l.Errorf("Unmarshal error,err=%v", err)
			_ = msg.Nack(false, false)
			continue
		}
		errC := make(chan error)
		orderID := res.Order.ID
		go checkErr(errC, msg, orderID)
		// 外卖员ID
		takeOutUserID := utils.NewUUID()
		nT := int(time.Now().Unix())
		wb := &wc.Waybill{
			ID:               orderID,
			Created:          nT,
			TakeOutUserID:    takeOutUserID,
			TakeOutUserName:  "小明",
			Comment:          "test",
			OrderReceiveTime: nT,
			// 预计一个小时候到达，实际根据算法平台给出时间
			ExpectArriveTime: nT + 60*60*1,
		}
		err = db.Model(wc.Waybill{}).Create(wb).Error
		// test
		//err = &mysql.MySQLError{Number: 123}
		if err != nil {
			e := err.(*mysql.MySQLError)
			// 若数据已经存在
			if e.Number == 1062 {
				l.Errorf("order is already, ID=%s", wb.ID)
				err = msg.Ack(false)
			} else {
				errC <- err
				l.Errorf("Create Waybill error,err=%v", err)
			}
			continue
		}
		err = msg.Ack(false)
		if err != nil {
			l.Errorf("ack error,err=%v", err)
			continue
		}
		log.Print("get message ok")
	}
}

// checkErr:校验错误，若相同订单重复发送5次及以上，仍报错，则转入死信队列，等待处理
func checkErr(c chan error, msg amqp.Delivery, orderID string) {
	select {
	case e := <-c:
		if e == nil {
			return
		}
		if _, ok := sendCountMap[orderID]; !ok {
			sendCountMap[orderID] = 1
		} else {
			if sendCountMap[orderID] >= 5 {
				_ = msg.Nack(false, false)
				delete(sendCountMap, orderID)
				log.Println("通知死信队列")
				return
			}
			sendCountMap[orderID] += 1
		}
		// 通知rabbitmq重发该消息
		_ = msg.Nack(true, true)
		return
	}
}

func ReceiveDead(conn *amqp.Connection, queue string, l *logrus.Logger, db *gorm.DB) {
	ch, msgs, err := getMsgs(queue, conn, l)
	if err != nil {
		l.Errorf("make conn chan error, err=%v", err)
		return
	}
	defer ch.Close()
	for msg := range msgs {
		log.Println("死信队列接受到消息")
		_ = msg.Ack(false)
		// todo notify admin
		continue
	}
}
