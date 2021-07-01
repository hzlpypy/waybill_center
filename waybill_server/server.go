package waybill_server

import (
	"context"
	cm "github.com/hzlpypy/common/databases/mysql"
	wc "github.com/hzlpypy/common/model/waybill_center"
	protos "github.com/hzlpypy/waybill_center/proto_info/protos"
	"gorm.io/gorm"
)

type WaybillServer struct {
	*protos.UnimplementedWaybillCenterServer
	Db *gorm.DB
}

func (w *WaybillServer) ListWaybill(ctx context.Context, req *protos.ListWaybillReq) (*protos.ListWaybillRes, error) {
	db := w.Db.Model(&wc.Waybill{})
	if len(req.OrderIds) > 0 {
		db = db.Where("id in ?", req.OrderIds)
	}
	wbs := []*wc.Waybill{}
	if !req.FindAll {
		limit, offect := cm.GetLimitOffset(req.Page, req.PageSize)
		db = db.Limit(int(limit)).Offset(int(offect))
	}
	err := db.Find(&wbs).Error
	if err != nil {
		return nil, err
	}
	pWbs := make([]*protos.Waybill, len(wbs))
	for i, wb := range wbs {
		pWbs[i] = &protos.Waybill{
			Id:              wb.ID,
			Created:         int32(wb.Created),
			TakeOutUserId:   wb.TakeOutUserID,
			TakeOutUserName: wb.TakeOutUserName,
			DeliveryTime:    int32(wb.DeliveryTime),
		}
	}
	return &protos.ListWaybillRes{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int32(len(wbs)),
		Waybills: pWbs,
	}, nil
}
