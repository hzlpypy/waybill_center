// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: waybill_center.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListWaybillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	OrderIds []string `protobuf:"bytes,3,rep,name=order_ids,json=orderIds,proto3" json:"order_ids,omitempty"`
	FindAll  bool     `protobuf:"varint,4,opt,name=find_all,json=findAll,proto3" json:"find_all,omitempty"`
}

func (x *ListWaybillReq) Reset() {
	*x = ListWaybillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waybill_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWaybillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWaybillReq) ProtoMessage() {}

func (x *ListWaybillReq) ProtoReflect() protoreflect.Message {
	mi := &file_waybill_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWaybillReq.ProtoReflect.Descriptor instead.
func (*ListWaybillReq) Descriptor() ([]byte, []int) {
	return file_waybill_center_proto_rawDescGZIP(), []int{0}
}

func (x *ListWaybillReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListWaybillReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWaybillReq) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

func (x *ListWaybillReq) GetFindAll() bool {
	if x != nil {
		return x.FindAll
	}
	return false
}

type Waybill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订单编号
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间
	Created int32 `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	// 配送外卖员ID
	TakeOutUserId string `protobuf:"bytes,3,opt,name=take_out_user_id,json=takeOutUserId,proto3" json:"take_out_user_id,omitempty"`
	// 配送外卖员名称
	TakeOutUserName string `protobuf:"bytes,4,opt,name=take_out_user_name,json=takeOutUserName,proto3" json:"take_out_user_name,omitempty"`
	// 送达时间 若为0, 表示未送达
	DeliveryTime int32 `protobuf:"varint,7,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"`
}

func (x *Waybill) Reset() {
	*x = Waybill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waybill_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Waybill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Waybill) ProtoMessage() {}

func (x *Waybill) ProtoReflect() protoreflect.Message {
	mi := &file_waybill_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Waybill.ProtoReflect.Descriptor instead.
func (*Waybill) Descriptor() ([]byte, []int) {
	return file_waybill_center_proto_rawDescGZIP(), []int{1}
}

func (x *Waybill) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Waybill) GetCreated() int32 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Waybill) GetTakeOutUserId() string {
	if x != nil {
		return x.TakeOutUserId
	}
	return ""
}

func (x *Waybill) GetTakeOutUserName() string {
	if x != nil {
		return x.TakeOutUserName
	}
	return ""
}

func (x *Waybill) GetDeliveryTime() int32 {
	if x != nil {
		return x.DeliveryTime
	}
	return 0
}

type ListWaybillRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int32      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Waybills []*Waybill `protobuf:"bytes,4,rep,name=waybills,proto3" json:"waybills,omitempty"`
}

func (x *ListWaybillRes) Reset() {
	*x = ListWaybillRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waybill_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWaybillRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWaybillRes) ProtoMessage() {}

func (x *ListWaybillRes) ProtoReflect() protoreflect.Message {
	mi := &file_waybill_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWaybillRes.ProtoReflect.Descriptor instead.
func (*ListWaybillRes) Descriptor() ([]byte, []int) {
	return file_waybill_center_proto_rawDescGZIP(), []int{2}
}

func (x *ListWaybillRes) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListWaybillRes) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWaybillRes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListWaybillRes) GetWaybills() []*Waybill {
	if x != nil {
		return x.Waybills
	}
	return nil
}

type OrderReceiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 外卖员ID
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 订单IDs
	OrderIds []string `protobuf:"bytes,2,rep,name=order_ids,json=orderIds,proto3" json:"order_ids,omitempty"`
}

func (x *OrderReceiveReq) Reset() {
	*x = OrderReceiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waybill_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReceiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReceiveReq) ProtoMessage() {}

func (x *OrderReceiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_waybill_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReceiveReq.ProtoReflect.Descriptor instead.
func (*OrderReceiveReq) Descriptor() ([]byte, []int) {
	return file_waybill_center_proto_rawDescGZIP(), []int{3}
}

func (x *OrderReceiveReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderReceiveReq) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

var File_waybill_center_proto protoreflect.FileDescriptor

var file_waybill_center_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x22, 0x79, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x22, 0xae, 0x01, 0x0a, 0x07,
	0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x27, 0x0a, 0x10, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x6b,
	0x65, 0x4f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x61,
	0x6b, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x6b, 0x65, 0x4f, 0x75, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x62, 0x69, 0x6c,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x52, 0x08, 0x77, 0x61, 0x79,
	0x62, 0x69, 0x6c, 0x6c, 0x73, 0x22, 0x47, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x32, 0x52,
	0x0a, 0x0d, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x41, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61,
	0x79, 0x62, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x79, 0x62, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waybill_center_proto_rawDescOnce sync.Once
	file_waybill_center_proto_rawDescData = file_waybill_center_proto_rawDesc
)

func file_waybill_center_proto_rawDescGZIP() []byte {
	file_waybill_center_proto_rawDescOnce.Do(func() {
		file_waybill_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_waybill_center_proto_rawDescData)
	})
	return file_waybill_center_proto_rawDescData
}

var file_waybill_center_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waybill_center_proto_goTypes = []interface{}{
	(*ListWaybillReq)(nil),  // 0: protobuf.ListWaybillReq
	(*Waybill)(nil),         // 1: protobuf.Waybill
	(*ListWaybillRes)(nil),  // 2: protobuf.ListWaybillRes
	(*OrderReceiveReq)(nil), // 3: protobuf.OrderReceiveReq
}
var file_waybill_center_proto_depIdxs = []int32{
	1, // 0: protobuf.ListWaybillRes.waybills:type_name -> protobuf.Waybill
	0, // 1: protobuf.WaybillCenter.ListWaybill:input_type -> protobuf.ListWaybillReq
	2, // 2: protobuf.WaybillCenter.ListWaybill:output_type -> protobuf.ListWaybillRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_waybill_center_proto_init() }
func file_waybill_center_proto_init() {
	if File_waybill_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waybill_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWaybillReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waybill_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Waybill); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waybill_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWaybillRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waybill_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReceiveReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waybill_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_waybill_center_proto_goTypes,
		DependencyIndexes: file_waybill_center_proto_depIdxs,
		MessageInfos:      file_waybill_center_proto_msgTypes,
	}.Build()
	File_waybill_center_proto = out.File
	file_waybill_center_proto_rawDesc = nil
	file_waybill_center_proto_goTypes = nil
	file_waybill_center_proto_depIdxs = nil
}
