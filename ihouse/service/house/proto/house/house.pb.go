// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.0
// source: proto/house/house.proto

package house

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

// 发布房源信息 传入参数
//
//	 {
//	    "title":"上奥世纪中心",
//	    "price":"666",
//	    "area_id":"5",
//	    "address":"西三旗桥东建材城1号",
//	    "room_count":"2",
//	    "acreage":"60",
//	    "unit":"2室1厅",
//	    "capacity":"3",
//	    "beds":"双人床2张",
//	    "deposit":"200",
//	    "min_days":"3",
//	     "max_days":"0",
//	     "facility":["1","2","3","7","12","14","16","17","18","21","22"]
//	}
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acreage   string   `protobuf:"bytes,1,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Address   string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AreaId    string   `protobuf:"bytes,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Beds      string   `protobuf:"bytes,4,opt,name=beds,proto3" json:"beds,omitempty"`
	Capacity  string   `protobuf:"bytes,5,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Deposit   string   `protobuf:"bytes,6,opt,name=deposit,proto3" json:"deposit,omitempty"`
	Facility  []string `protobuf:"bytes,7,rep,name=facility,proto3" json:"facility,omitempty"`
	MaxDays   string   `protobuf:"bytes,8,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	MinDays   string   `protobuf:"bytes,9,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	Price     string   `protobuf:"bytes,10,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount string   `protobuf:"bytes,11,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title     string   `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	Unit      string   `protobuf:"bytes,13,opt,name=unit,proto3" json:"unit,omitempty"`
	UserName  string   `protobuf:"bytes,14,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAcreage() string {
	if x != nil {
		return x.Acreage
	}
	return ""
}

func (x *Request) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Request) GetAreaId() string {
	if x != nil {
		return x.AreaId
	}
	return ""
}

func (x *Request) GetBeds() string {
	if x != nil {
		return x.Beds
	}
	return ""
}

func (x *Request) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

func (x *Request) GetDeposit() string {
	if x != nil {
		return x.Deposit
	}
	return ""
}

func (x *Request) GetFacility() []string {
	if x != nil {
		return x.Facility
	}
	return nil
}

func (x *Request) GetMaxDays() string {
	if x != nil {
		return x.MaxDays
	}
	return ""
}

func (x *Request) GetMinDays() string {
	if x != nil {
		return x.MinDays
	}
	return ""
}

func (x *Request) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Request) GetRoomCount() string {
	if x != nil {
		return x.RoomCount
	}
	return ""
}

func (x *Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Request) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string     `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string     `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *HouseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetData() *HouseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type HouseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseId string `protobuf:"bytes,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
}

func (x *HouseData) Reset() {
	*x = HouseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseData) ProtoMessage() {}

func (x *HouseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseData.ProtoReflect.Descriptor instead.
func (*HouseData) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{2}
}

func (x *HouseData) GetHouseId() string {
	if x != nil {
		return x.HouseId
	}
	return ""
}

// 传入参数
type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{3}
}

func (x *GetReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 传出参数
type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *GetData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{4}
}

func (x *GetResp) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *GetResp) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *GetResp) GetData() *GetData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses []*Houses `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
}

func (x *GetData) Reset() {
	*x = GetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetData) ProtoMessage() {}

func (x *GetData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetData.ProtoReflect.Descriptor instead.
func (*GetData) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{5}
}

func (x *GetData) GetHouses() []*Houses {
	if x != nil {
		return x.Houses
	}
	return nil
}

type Houses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AreaName   string `protobuf:"bytes,2,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	Ctime      string `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	HouseId    int32  `protobuf:"varint,4,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	ImgUrl     string `protobuf:"bytes,5,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderCount int32  `protobuf:"varint,6,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	Price      int32  `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount  int32  `protobuf:"varint,8,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title      string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	UserAvatar string `protobuf:"bytes,10,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
}

func (x *Houses) Reset() {
	*x = Houses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Houses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Houses) ProtoMessage() {}

func (x *Houses) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Houses.ProtoReflect.Descriptor instead.
func (*Houses) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{6}
}

func (x *Houses) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Houses) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *Houses) GetCtime() string {
	if x != nil {
		return x.Ctime
	}
	return ""
}

func (x *Houses) GetHouseId() int32 {
	if x != nil {
		return x.HouseId
	}
	return 0
}

func (x *Houses) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *Houses) GetOrderCount() int32 {
	if x != nil {
		return x.OrderCount
	}
	return 0
}

func (x *Houses) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Houses) GetRoomCount() int32 {
	if x != nil {
		return x.RoomCount
	}
	return 0
}

func (x *Houses) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Houses) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

// 添加房屋图片 输入参数
// 传入参数  房屋 id、图片二进制、图片后缀
type ImgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseId string `protobuf:"bytes,1,opt,name=houseId,proto3" json:"houseId,omitempty"`
	ImgData []byte `protobuf:"bytes,2,opt,name=imgData,proto3" json:"imgData,omitempty"`
	FileExt string `protobuf:"bytes,3,opt,name=fileExt,proto3" json:"fileExt,omitempty"`
}

func (x *ImgReq) Reset() {
	*x = ImgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgReq) ProtoMessage() {}

func (x *ImgReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgReq.ProtoReflect.Descriptor instead.
func (*ImgReq) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{7}
}

func (x *ImgReq) GetHouseId() string {
	if x != nil {
		return x.HouseId
	}
	return ""
}

func (x *ImgReq) GetImgData() []byte {
	if x != nil {
		return x.ImgData
	}
	return nil
}

func (x *ImgReq) GetFileExt() string {
	if x != nil {
		return x.FileExt
	}
	return ""
}

// {
// "errno": "0",
// "errmsg": "成功",
// "data": {
// "url": "http://101.200.170.171:9998/group1/M00/00/00/Zciqq1oBLmWAHlsrAAaInSze-cQ719.jpg"
// }
// }
type ImgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *ImgData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ImgResp) Reset() {
	*x = ImgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgResp) ProtoMessage() {}

func (x *ImgResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgResp.ProtoReflect.Descriptor instead.
func (*ImgResp) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{8}
}

func (x *ImgResp) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *ImgResp) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *ImgResp) GetData() *ImgData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ImgData) Reset() {
	*x = ImgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgData) ProtoMessage() {}

func (x *ImgData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgData.ProtoReflect.Descriptor instead.
func (*ImgData) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{9}
}

func (x *ImgData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// 轮播图
type IndexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexReq) Reset() {
	*x = IndexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_house_house_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReq) ProtoMessage() {}

func (x *IndexReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_house_house_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReq.ProtoReflect.Descriptor instead.
func (*IndexReq) Descriptor() ([]byte, []int) {
	return file_proto_house_house_proto_rawDescGZIP(), []int{10}
}

var File_proto_house_house_proto protoreflect.FileDescriptor

var file_proto_house_house_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0xed, 0x02,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x72,
	0x65, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x72, 0x65,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x65, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x78, 0x44, 0x61, 0x79, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x44, 0x61,
	0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x09, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x64, 0x22, 0x24, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x3d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x06, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x52, 0x06, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x22, 0x96, 0x02, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x56, 0x0a, 0x06, 0x49, 0x6d,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x6d, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x69, 0x6d, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x45, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45,
	0x78, 0x74, 0x22, 0x68, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x49,
	0x6d, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x07,
	0x49, 0x6d, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0a, 0x0a, 0x08, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x71, 0x32, 0xb6, 0x02, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x08, 0x50, 0x75, 0x62, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x49, 0x6d, 0x67, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_house_house_proto_rawDescOnce sync.Once
	file_proto_house_house_proto_rawDescData = file_proto_house_house_proto_rawDesc
)

func file_proto_house_house_proto_rawDescGZIP() []byte {
	file_proto_house_house_proto_rawDescOnce.Do(func() {
		file_proto_house_house_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_house_house_proto_rawDescData)
	})
	return file_proto_house_house_proto_rawDescData
}

var file_proto_house_house_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_house_house_proto_goTypes = []interface{}{
	(*Request)(nil),   // 0: go.micro.srv.house.Request
	(*Response)(nil),  // 1: go.micro.srv.house.Response
	(*HouseData)(nil), // 2: go.micro.srv.house.HouseData
	(*GetReq)(nil),    // 3: go.micro.srv.house.GetReq
	(*GetResp)(nil),   // 4: go.micro.srv.house.GetResp
	(*GetData)(nil),   // 5: go.micro.srv.house.GetData
	(*Houses)(nil),    // 6: go.micro.srv.house.Houses
	(*ImgReq)(nil),    // 7: go.micro.srv.house.ImgReq
	(*ImgResp)(nil),   // 8: go.micro.srv.house.ImgResp
	(*ImgData)(nil),   // 9: go.micro.srv.house.ImgData
	(*IndexReq)(nil),  // 10: go.micro.srv.house.IndexReq
}
var file_proto_house_house_proto_depIdxs = []int32{
	2,  // 0: go.micro.srv.house.Response.data:type_name -> go.micro.srv.house.HouseData
	5,  // 1: go.micro.srv.house.GetResp.data:type_name -> go.micro.srv.house.GetData
	6,  // 2: go.micro.srv.house.GetData.houses:type_name -> go.micro.srv.house.Houses
	9,  // 3: go.micro.srv.house.ImgResp.data:type_name -> go.micro.srv.house.ImgData
	0,  // 4: go.micro.srv.house.House.PubHouse:input_type -> go.micro.srv.house.Request
	3,  // 5: go.micro.srv.house.House.GetHouseInfo:input_type -> go.micro.srv.house.GetReq
	7,  // 6: go.micro.srv.house.House.UploadHouseImg:input_type -> go.micro.srv.house.ImgReq
	10, // 7: go.micro.srv.house.House.GetIndexHouse:input_type -> go.micro.srv.house.IndexReq
	1,  // 8: go.micro.srv.house.House.PubHouse:output_type -> go.micro.srv.house.Response
	4,  // 9: go.micro.srv.house.House.GetHouseInfo:output_type -> go.micro.srv.house.GetResp
	8,  // 10: go.micro.srv.house.House.UploadHouseImg:output_type -> go.micro.srv.house.ImgResp
	4,  // 11: go.micro.srv.house.House.GetIndexHouse:output_type -> go.micro.srv.house.GetResp
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_house_house_proto_init() }
func file_proto_house_house_proto_init() {
	if File_proto_house_house_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_house_house_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_house_house_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_house_house_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseData); i {
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
		file_proto_house_house_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_proto_house_house_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
		file_proto_house_house_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetData); i {
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
		file_proto_house_house_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Houses); i {
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
		file_proto_house_house_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgReq); i {
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
		file_proto_house_house_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgResp); i {
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
		file_proto_house_house_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgData); i {
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
		file_proto_house_house_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexReq); i {
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
			RawDescriptor: file_proto_house_house_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_house_house_proto_goTypes,
		DependencyIndexes: file_proto_house_house_proto_depIdxs,
		MessageInfos:      file_proto_house_house_proto_msgTypes,
	}.Build()
	File_proto_house_house_proto = out.File
	file_proto_house_house_proto_rawDesc = nil
	file_proto_house_house_proto_goTypes = nil
	file_proto_house_house_proto_depIdxs = nil
}
