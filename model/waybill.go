package model


type Waybill struct {
	// 订单编号
	ID string `gorm:"type:varchar(32);primary_key;uniqueIndex:index_uid_oid;uniqueIndex:index_touid_oid;"`
	// UnixTime
	// 创建时间/下单时间
	Created  int    `gorm:"type:int(10);not null"`
	Updated  int    `gorm:"type:int(10);default:0;autoUpdateTime;"`
	Deleted  int    `gorm:"type:int(10);default:0;autoCreateTime"`
	Handlers string `gorm:"type:varchar(32);default:''"`
	// 配送外卖员ID
	TakeOutUserID string `gorm:"type:varchar(32);uniqueIndex:index_touid_oid;not null"`
	// 配送外卖员名称
	TakeOutUserName string `gorm:"type:varchar(32);not null"`
	// 备注
	Comment string `gorm:"type:varchar(128);not null"`
	// 接单时间
	OrderReceiveTime int `gorm:"type:int(10);not null"`
	// 预计到达时间
	ExpectArriveTime int `gorm:"type:int(10);not null"`
	// 送达时间
	DeliveryTime int `gorm:"type:int(10);not null"`
	// 是否被删除 1：否 2：是
	FakeDelete int `gorm:"type:tinyint(1);default:1"`
	// ...
}