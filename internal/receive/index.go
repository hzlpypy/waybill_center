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

func Receive(conn *amqp.Connection, queue string, l *logrus.Logger, db *gorm.DB) {
	ch, err := conn.Channel()
	if err != nil {
		l.Errorf("make conn chan error, err=%v", err)
		return
	}
	defer ch.Close()
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
		return
	}
	for msg := range msgs {
		res := &oc.OrderMsg{}
		err := json.Unmarshal(msg.Body, &res)
		if err != nil {
			l.Errorf("Unmarshal error,err=%v", err)
			continue
		}
		// 外卖员ID
		takeOutUserID := utils.NewUUID()
		nT := int(time.Now().Unix())
		wb := &wc.Waybill{
			ID:               res.Order.ID,
			Created:          nT,
			TakeOutUserID:    takeOutUserID,
			TakeOutUserName:  "小明",
			Comment:          "test",
			OrderReceiveTime: nT,
			// 预计一个小时候到达，实际根据算法平台给出时间
			ExpectArriveTime: nT + 60*60*1,
		}
		err = db.Model(wc.Waybill{}).Create(wb).Error
		if err != nil {
			e := err.(*mysql.MySQLError)
			// 若数据已经存在
			if e.Number == 1062 {
				l.Errorf("order is already, ID=%s", wb.ID)
				err = msg.Ack(false)
			} else {
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
