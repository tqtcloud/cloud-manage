// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: apps/order/pb/order.proto

package order

import (
	resource "github.com/tqtcloud/cloud-manage/apps/resource"
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

type ORDER_TYPE int32

const (
	// 资源新增订单
	ORDER_TYPE_NEW ORDER_TYPE = 0
	// 续费。
	ORDER_TYPE_RENEW ORDER_TYPE = 1
	// 升级。
	ORDER_TYPE_UPGRADE ORDER_TYPE = 2
	// 降配
	ORDER_TYPE_DOWNGRADE ORDER_TYPE = 3
	// 退款
	ORDER_TYPE_REFUND ORDER_TYPE = 4
	// 付费方式转换
	ORDER_TYPE_CONVERT ORDER_TYPE = 5
	// 配置调整
	ORDER_TYPE_MODIFY ORDER_TYPE = 6
)

// Enum value maps for ORDER_TYPE.
var (
	ORDER_TYPE_name = map[int32]string{
		0: "NEW",
		1: "RENEW",
		2: "UPGRADE",
		3: "DOWNGRADE",
		4: "REFUND",
		5: "CONVERT",
		6: "MODIFY",
	}
	ORDER_TYPE_value = map[string]int32{
		"NEW":       0,
		"RENEW":     1,
		"UPGRADE":   2,
		"DOWNGRADE": 3,
		"REFUND":    4,
		"CONVERT":   5,
		"MODIFY":    6,
	}
)

func (x ORDER_TYPE) Enum() *ORDER_TYPE {
	p := new(ORDER_TYPE)
	*p = x
	return p
}

func (x ORDER_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ORDER_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_order_pb_order_proto_enumTypes[0].Descriptor()
}

func (ORDER_TYPE) Type() protoreflect.EnumType {
	return &file_apps_order_pb_order_proto_enumTypes[0]
}

func (x ORDER_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ORDER_TYPE.Descriptor instead.
func (ORDER_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_order_pb_order_proto_rawDescGZIP(), []int{0}
}

type ORDER_STATUS int32

const (
	// 未支付
	ORDER_STATUS_UNPAID ORDER_STATUS = 0
	// 订单过期
	ORDER_STATUS_EXPIRED ORDER_STATUS = 1
	// 支付中
	ORDER_STATUS_PAYING ORDER_STATUS = 2
	// 已取消
	ORDER_STATUS_CANCELLED ORDER_STATUS = 3
	// 已支付
	ORDER_STATUS_PAID ORDER_STATUS = 4
)

// Enum value maps for ORDER_STATUS.
var (
	ORDER_STATUS_name = map[int32]string{
		0: "UNPAID",
		1: "EXPIRED",
		2: "PAYING",
		3: "CANCELLED",
		4: "PAID",
	}
	ORDER_STATUS_value = map[string]int32{
		"UNPAID":    0,
		"EXPIRED":   1,
		"PAYING":    2,
		"CANCELLED": 3,
		"PAID":      4,
	}
)

func (x ORDER_STATUS) Enum() *ORDER_STATUS {
	p := new(ORDER_STATUS)
	*p = x
	return p
}

func (x ORDER_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ORDER_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_order_pb_order_proto_enumTypes[1].Descriptor()
}

func (ORDER_STATUS) Type() protoreflect.EnumType {
	return &file_apps_order_pb_order_proto_enumTypes[1]
}

func (x ORDER_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ORDER_STATUS.Descriptor instead.
func (ORDER_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_apps_order_pb_order_proto_rawDescGZIP(), []int{1}
}

type OrderSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表项
	// @gotags: json:"items"
	Items []*Order `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *OrderSet) Reset() {
	*x = OrderSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_order_pb_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSet) ProtoMessage() {}

func (x *OrderSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_order_pb_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSet.ProtoReflect.Descriptor instead.
func (*OrderSet) Descriptor() ([]byte, []int) {
	return file_apps_order_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderSet) GetItems() []*Order {
	if x != nil {
		return x.Items
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 本次账单同步关联的任务Id
	// @gotags: json:"task_id"
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor resource.VENDOR `protobuf:"varint,2,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.VENDOR" json:"vendor"`
	// 大订单号
	// @gotags: json:"big_order_id"
	BigOrderId string `protobuf:"bytes,3,opt,name=big_order_id,json=bigOrderId,proto3" json:"big_order_id"`
	// 订单Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id"`
	// 订单类型
	// @gotags: json:"order_type"
	OrderType string `protobuf:"bytes,5,opt,name=order_type,json=orderType,proto3" json:"order_type"`
	// 订单状态
	// @gotags: json:"status"
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	// 支付者
	// @gotags: json:"payer"
	Payer string `protobuf:"bytes,7,opt,name=payer,proto3" json:"payer"`
	// 创建时间
	// @gotags: json:"create_at"
	CreateAt int64 `protobuf:"varint,8,opt,name=create_at,json=createAt,proto3" json:"create_at"`
	// 支付时间
	// @gotags: json:"pay_at"
	PayAt int64 `protobuf:"varint,9,opt,name=pay_at,json=payAt,proto3" json:"pay_at"`
	// 创建人
	// @gotags: json:"create_by"
	CreateBy string `protobuf:"bytes,10,opt,name=create_by,json=createBy,proto3" json:"create_by"`
	// 计费方式，比如：月95，日均峰值
	// @gotags: json:"pay_mode"
	PayMode resource.PayMode `protobuf:"varint,11,opt,name=pay_mode,json=payMode,proto3,enum=infraboard.cmdb.resource.PayMode" json:"pay_mode"`
	// 定义费用
	// @gotags: json:"cost"
	Cost *Cost `protobuf:"bytes,12,opt,name=cost,proto3" json:"cost"`
	// 产品编码
	// @gotags: json:"product_code"
	ProductCode string `protobuf:"bytes,13,opt,name=product_code,json=productCode,proto3" json:"product_code"`
	// 产品编码中文名称
	// @gotags: json:"product_name"
	ProductName string `protobuf:"bytes,14,opt,name=product_name,json=productName,proto3" json:"product_name"`
	// 子产品编码
	// @gotags: json:"sub_product_code"
	SubProductCode string `protobuf:"bytes,15,opt,name=sub_product_code,json=subProductCode,proto3" json:"sub_product_code"`
	// 子产品编码中文名称
	// @gotags: json:"sub_product_name"
	SubProductName string `protobuf:"bytes,16,opt,name=sub_product_name,json=subProductName,proto3" json:"sub_product_name"`
	// 购买产品详情
	// @gotags: json:"product_info"
	ProductInfo string `protobuf:"bytes,17,opt,name=product_info,json=productInfo,proto3" json:"product_info"`
	// 订单对应的资源id
	// @gotags: json:"resource_id"
	ResourceId []string `protobuf:"bytes,18,rep,name=resource_id,json=resourceId,proto3" json:"resource_id"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_order_pb_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_apps_order_pb_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_apps_order_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Order) GetVendor() resource.VENDOR {
	if x != nil {
		return x.Vendor
	}
	return resource.VENDOR(0)
}

func (x *Order) GetBigOrderId() string {
	if x != nil {
		return x.BigOrderId
	}
	return ""
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetPayer() string {
	if x != nil {
		return x.Payer
	}
	return ""
}

func (x *Order) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Order) GetPayAt() int64 {
	if x != nil {
		return x.PayAt
	}
	return 0
}

func (x *Order) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *Order) GetPayMode() resource.PayMode {
	if x != nil {
		return x.PayMode
	}
	return resource.PayMode(0)
}

func (x *Order) GetCost() *Cost {
	if x != nil {
		return x.Cost
	}
	return nil
}

func (x *Order) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Order) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Order) GetSubProductCode() string {
	if x != nil {
		return x.SubProductCode
	}
	return ""
}

func (x *Order) GetSubProductName() string {
	if x != nil {
		return x.SubProductName
	}
	return ""
}

func (x *Order) GetProductInfo() string {
	if x != nil {
		return x.ProductInfo
	}
	return ""
}

func (x *Order) GetResourceId() []string {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

type Cost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 官网价,原价（分）
	// @gotags: json:"sale_price"
	SalePrice float64 `protobuf:"fixed64,1,opt,name=sale_price,json=salePrice,proto3" json:"sale_price"`
	// 折扣率
	// @gotags: json:"policy"
	Policy float64 `protobuf:"fixed64,2,opt,name=policy,proto3" json:"policy"`
	// 单价（分）
	// @gotags: json:"unit_price"
	UnitPrice float64 `protobuf:"fixed64,3,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price"`
	// 购买时长
	// @gotags: json:"time_span"
	TimeSpan float64 `protobuf:"fixed64,4,opt,name=time_span,json=timeSpan,proto3" json:"time_span"`
	// 购买时长
	// @gotags: json:"time_unit"
	TimeUnit string `protobuf:"bytes,5,opt,name=time_unit,json=timeUnit,proto3" json:"time_unit"`
	// 实际支付金额（分）
	// @gotags: json:"real_cost"
	RealCost float64 `protobuf:"fixed64,6,opt,name=real_cost,json=realCost,proto3" json:"real_cost"`
	// 代金券抵扣金额（分）
	// @gotags: json:"voucher_pay"
	VoucherPay float64 `protobuf:"fixed64,7,opt,name=voucher_pay,json=voucherPay,proto3" json:"voucher_pay"`
}

func (x *Cost) Reset() {
	*x = Cost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_order_pb_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cost) ProtoMessage() {}

func (x *Cost) ProtoReflect() protoreflect.Message {
	mi := &file_apps_order_pb_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cost.ProtoReflect.Descriptor instead.
func (*Cost) Descriptor() ([]byte, []int) {
	return file_apps_order_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *Cost) GetSalePrice() float64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *Cost) GetPolicy() float64 {
	if x != nil {
		return x.Policy
	}
	return 0
}

func (x *Cost) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *Cost) GetTimeSpan() float64 {
	if x != nil {
		return x.TimeSpan
	}
	return 0
}

func (x *Cost) GetTimeUnit() string {
	if x != nil {
		return x.TimeUnit
	}
	return ""
}

func (x *Cost) GetRealCost() float64 {
	if x != nil {
		return x.RealCost
	}
	return 0
}

func (x *Cost) GetVoucherPay() float64 {
	if x != nil {
		return x.VoucherPay
	}
	return 0
}

var File_apps_order_pb_order_proto protoreflect.FileDescriptor

var file_apps_order_pb_order_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xf7, 0x04, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x52, 0x06,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x69, 0x67, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69,
	0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x79, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x50, 0x61, 0x79, 0x2a, 0x61, 0x0a, 0x0a, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x4f, 0x57,
	0x4e, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x46, 0x55,
	0x4e, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x10,
	0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x10, 0x06, 0x2a, 0x4c, 0x0a,
	0x0c, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x4e, 0x50, 0x41, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x59, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x04, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x71, 0x74, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_order_pb_order_proto_rawDescOnce sync.Once
	file_apps_order_pb_order_proto_rawDescData = file_apps_order_pb_order_proto_rawDesc
)

func file_apps_order_pb_order_proto_rawDescGZIP() []byte {
	file_apps_order_pb_order_proto_rawDescOnce.Do(func() {
		file_apps_order_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_order_pb_order_proto_rawDescData)
	})
	return file_apps_order_pb_order_proto_rawDescData
}

var file_apps_order_pb_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_order_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_order_pb_order_proto_goTypes = []interface{}{
	(ORDER_TYPE)(0),       // 0: infraboard.cmdb.order.ORDER_TYPE
	(ORDER_STATUS)(0),     // 1: infraboard.cmdb.order.ORDER_STATUS
	(*OrderSet)(nil),      // 2: infraboard.cmdb.order.OrderSet
	(*Order)(nil),         // 3: infraboard.cmdb.order.Order
	(*Cost)(nil),          // 4: infraboard.cmdb.order.Cost
	(resource.VENDOR)(0),  // 5: infraboard.cmdb.resource.VENDOR
	(resource.PayMode)(0), // 6: infraboard.cmdb.resource.PayMode
}
var file_apps_order_pb_order_proto_depIdxs = []int32{
	3, // 0: infraboard.cmdb.order.OrderSet.items:type_name -> infraboard.cmdb.order.Order
	5, // 1: infraboard.cmdb.order.Order.vendor:type_name -> infraboard.cmdb.resource.VENDOR
	6, // 2: infraboard.cmdb.order.Order.pay_mode:type_name -> infraboard.cmdb.resource.PayMode
	4, // 3: infraboard.cmdb.order.Order.cost:type_name -> infraboard.cmdb.order.Cost
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_order_pb_order_proto_init() }
func file_apps_order_pb_order_proto_init() {
	if File_apps_order_pb_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_order_pb_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSet); i {
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
		file_apps_order_pb_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_apps_order_pb_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cost); i {
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
			RawDescriptor: file_apps_order_pb_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_order_pb_order_proto_goTypes,
		DependencyIndexes: file_apps_order_pb_order_proto_depIdxs,
		EnumInfos:         file_apps_order_pb_order_proto_enumTypes,
		MessageInfos:      file_apps_order_pb_order_proto_msgTypes,
	}.Build()
	File_apps_order_pb_order_proto = out.File
	file_apps_order_pb_order_proto_rawDesc = nil
	file_apps_order_pb_order_proto_goTypes = nil
	file_apps_order_pb_order_proto_depIdxs = nil
}
