// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: config.proto

package api

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

type DanceSuitListItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SuitId        int64                  `protobuf:"varint,1,opt,name=suit_id,json=suitId,proto3" json:"suit_id,omitempty"`                 //套装id
	Sex           int32                  `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`                                     //0是男，1是女
	ElementId     []int64                `protobuf:"varint,3,rep,packed,name=element_id,json=elementId,proto3" json:"element_id,omitempty"` //元素id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceSuitListItem) Reset() {
	*x = DanceSuitListItem{}
	mi := &file_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceSuitListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceSuitListItem) ProtoMessage() {}

func (x *DanceSuitListItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceSuitListItem.ProtoReflect.Descriptor instead.
func (*DanceSuitListItem) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *DanceSuitListItem) GetSuitId() int64 {
	if x != nil {
		return x.SuitId
	}
	return 0
}

func (x *DanceSuitListItem) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *DanceSuitListItem) GetElementId() []int64 {
	if x != nil {
		return x.ElementId
	}
	return nil
}

//等级返回
type DanceLevelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceLevelRequest) Reset() {
	*x = DanceLevelRequest{}
	mi := &file_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceLevelRequest) ProtoMessage() {}

func (x *DanceLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceLevelRequest.ProtoReflect.Descriptor instead.
func (*DanceLevelRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

//套装基础列表
type DanceLevelResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	//base 全部套装，部分，公共，个人等。。
	BaseAllSuit             []*DanceSuitListItem     `protobuf:"bytes,1,rep,name=base_all_suit,json=baseAllSuit,proto3" json:"base_all_suit,omitempty"`                                               //全量套装，库里面的全部套装取出，组装进去
	PublicPoolManSuit       []int64                  `protobuf:"varint,2,rep,packed,name=public_pool_man_suit,json=publicPoolManSuit,proto3" json:"public_pool_man_suit,omitempty"`                   // 公共池的男性套装id
	PublicPoolWomanSuit     []int64                  `protobuf:"varint,3,rep,packed,name=public_pool_woman_suit,json=publicPoolWomanSuit,proto3" json:"public_pool_woman_suit,omitempty"`             //公共池的女性套装id
	ElementList             []*DanceElementLocalItem `protobuf:"bytes,4,rep,name=elementList,proto3" json:"elementList,omitempty"`                                                                    //元素对应的部位信息
	ActivitiesPoolManSuit   []int64                  `protobuf:"varint,5,rep,packed,name=activities_pool_man_suit,json=activitiesPoolManSuit,proto3" json:"activities_pool_man_suit,omitempty"`       // 活动池的男性套装id
	ActivitiesPoolWomanSuit []int64                  `protobuf:"varint,6,rep,packed,name=activities_pool_woman_suit,json=activitiesPoolWomanSuit,proto3" json:"activities_pool_woman_suit,omitempty"` //活动池的女性套装id
	ComposePoolManSuit      []int64                  `protobuf:"varint,7,rep,packed,name=compose_pool_man_suit,json=composePoolManSuit,proto3" json:"compose_pool_man_suit,omitempty"`                // 组合男性套装id
	ComposePoolWomanSuit    []int64                  `protobuf:"varint,8,rep,packed,name=compose_pool_woman_suit,json=composePoolWomanSuit,proto3" json:"compose_pool_woman_suit,omitempty"`          //组合女性套装id
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *DanceLevelResponse) Reset() {
	*x = DanceLevelResponse{}
	mi := &file_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceLevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceLevelResponse) ProtoMessage() {}

func (x *DanceLevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceLevelResponse.ProtoReflect.Descriptor instead.
func (*DanceLevelResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *DanceLevelResponse) GetBaseAllSuit() []*DanceSuitListItem {
	if x != nil {
		return x.BaseAllSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetPublicPoolManSuit() []int64 {
	if x != nil {
		return x.PublicPoolManSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetPublicPoolWomanSuit() []int64 {
	if x != nil {
		return x.PublicPoolWomanSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetElementList() []*DanceElementLocalItem {
	if x != nil {
		return x.ElementList
	}
	return nil
}

func (x *DanceLevelResponse) GetActivitiesPoolManSuit() []int64 {
	if x != nil {
		return x.ActivitiesPoolManSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetActivitiesPoolWomanSuit() []int64 {
	if x != nil {
		return x.ActivitiesPoolWomanSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetComposePoolManSuit() []int64 {
	if x != nil {
		return x.ComposePoolManSuit
	}
	return nil
}

func (x *DanceLevelResponse) GetComposePoolWomanSuit() []int64 {
	if x != nil {
		return x.ComposePoolWomanSuit
	}
	return nil
}

type DanceElementLocalItem struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	DanceElementNo int64                  `protobuf:"varint,1,opt,name=dance_element_no,json=danceElementNo,proto3" json:"dance_element_no,omitempty"`     //散装编码
	ElementLocalS  []int64                `protobuf:"varint,2,rep,packed,name=element_local_s,json=elementLocalS,proto3" json:"element_local_s,omitempty"` //散装部位数据
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DanceElementLocalItem) Reset() {
	*x = DanceElementLocalItem{}
	mi := &file_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceElementLocalItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceElementLocalItem) ProtoMessage() {}

func (x *DanceElementLocalItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceElementLocalItem.ProtoReflect.Descriptor instead.
func (*DanceElementLocalItem) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *DanceElementLocalItem) GetDanceElementNo() int64 {
	if x != nil {
		return x.DanceElementNo
	}
	return 0
}

func (x *DanceElementLocalItem) GetElementLocalS() []int64 {
	if x != nil {
		return x.ElementLocalS
	}
	return nil
}

//光环按钮
type DanceHaloButtonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloButtonRequest) Reset() {
	*x = DanceHaloButtonRequest{}
	mi := &file_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloButtonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloButtonRequest) ProtoMessage() {}

func (x *DanceHaloButtonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloButtonRequest.ProtoReflect.Descriptor instead.
func (*DanceHaloButtonRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

//套装基础列表
type DanceHaloButtonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsOpen        bool                   `protobuf:"varint,1,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"` //false:关闭;true:打开
	Icon          string                 `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`                    //url地址
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloButtonResponse) Reset() {
	*x = DanceHaloButtonResponse{}
	mi := &file_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloButtonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloButtonResponse) ProtoMessage() {}

func (x *DanceHaloButtonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloButtonResponse.ProtoReflect.Descriptor instead.
func (*DanceHaloButtonResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *DanceHaloButtonResponse) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *DanceHaloButtonResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

//光环按钮
type DanceHaloRulesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRulesRequest) Reset() {
	*x = DanceHaloRulesRequest{}
	mi := &file_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRulesRequest) ProtoMessage() {}

func (x *DanceHaloRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRulesRequest.ProtoReflect.Descriptor instead.
func (*DanceHaloRulesRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

//套装基础列表
type DanceHaloRulesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BeginTime     int64                  `protobuf:"varint,1,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"` //开始时间
	EndTime       int64                  `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       //结束时间
	Desc          string                 `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`                             //活动说明
	List          []*DanceHaloRulesItem  `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`                             //散装部位数据
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRulesResponse) Reset() {
	*x = DanceHaloRulesResponse{}
	mi := &file_config_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRulesResponse) ProtoMessage() {}

func (x *DanceHaloRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRulesResponse.ProtoReflect.Descriptor instead.
func (*DanceHaloRulesResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{7}
}

func (x *DanceHaloRulesResponse) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *DanceHaloRulesResponse) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *DanceHaloRulesResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DanceHaloRulesResponse) GetList() []*DanceHaloRulesItem {
	if x != nil {
		return x.List
	}
	return nil
}

type DanceHaloRulesItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Score         int64                  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`                          //星光值
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                             //光环名字
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`                               //光环图片地址
	DayNumber     int32                  `protobuf:"varint,4,opt,name=day_number,json=dayNumber,proto3" json:"day_number,omitempty"` //有效天数
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRulesItem) Reset() {
	*x = DanceHaloRulesItem{}
	mi := &file_config_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRulesItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRulesItem) ProtoMessage() {}

func (x *DanceHaloRulesItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRulesItem.ProtoReflect.Descriptor instead.
func (*DanceHaloRulesItem) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{8}
}

func (x *DanceHaloRulesItem) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *DanceHaloRulesItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DanceHaloRulesItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DanceHaloRulesItem) GetDayNumber() int32 {
	if x != nil {
		return x.DayNumber
	}
	return 0
}

type DanceHaloRankRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRankRequest) Reset() {
	*x = DanceHaloRankRequest{}
	mi := &file_config_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRankRequest) ProtoMessage() {}

func (x *DanceHaloRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRankRequest.ProtoReflect.Descriptor instead.
func (*DanceHaloRankRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{9}
}

func (x *DanceHaloRankRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

//获取排名前多少的用户
type DanceHaloRankResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*DanceHaloRankItem   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRankResponse) Reset() {
	*x = DanceHaloRankResponse{}
	mi := &file_config_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRankResponse) ProtoMessage() {}

func (x *DanceHaloRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRankResponse.ProtoReflect.Descriptor instead.
func (*DanceHaloRankResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{10}
}

func (x *DanceHaloRankResponse) GetList() []*DanceHaloRankItem {
	if x != nil {
		return x.List
	}
	return nil
}

type DanceHaloRankItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalScore    int64                  `protobuf:"varint,1,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"` //累计积分
	WorldRank     int32                  `protobuf:"varint,2,opt,name=world_rank,json=worldRank,proto3" json:"world_rank,omitempty"`
	Nickname      string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"` //昵称
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`     //头像
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DanceHaloRankItem) Reset() {
	*x = DanceHaloRankItem{}
	mi := &file_config_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DanceHaloRankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanceHaloRankItem) ProtoMessage() {}

func (x *DanceHaloRankItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanceHaloRankItem.ProtoReflect.Descriptor instead.
func (*DanceHaloRankItem) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{11}
}

func (x *DanceHaloRankItem) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *DanceHaloRankItem) GetWorldRank() int32 {
	if x != nil {
		return x.WorldRank
	}
	return 0
}

func (x *DanceHaloRankItem) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *DanceHaloRankItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x11, 0x44, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x75, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x75, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x75, 0x69, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x61, 0x6e,
	0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe2,
	0x03, 0x0a, 0x12, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x53,
	0x75, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x69, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f,
	0x6f, 0x6c, 0x4d, 0x61, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x6f, 0x6d, 0x61, 0x6e, 0x5f, 0x73,
	0x75, 0x69, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x50, 0x6f, 0x6f, 0x6c, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x12, 0x43,
	0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x1a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x77, 0x6f, 0x6d, 0x61, 0x6e, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x6f, 0x6f, 0x6c,
	0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x5f, 0x73, 0x75,
	0x69, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x6f, 0x6d,
	0x61, 0x6e, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x14, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x53,
	0x75, 0x69, 0x74, 0x22, 0x69, 0x0a, 0x15, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x10,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x22, 0x18,
	0x0a, 0x16, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x6c, 0x6f, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x17, 0x44, 0x61, 0x6e, 0x63,
	0x65, 0x48, 0x61, 0x6c, 0x6f, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x22, 0x17, 0x0a, 0x15, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x6c, 0x6f, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x44, 0x61,
	0x6e, 0x63, 0x65, 0x48, 0x61, 0x6c, 0x6f, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61,
	0x6e, 0x63, 0x65, 0x48, 0x61, 0x6c, 0x6f, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x12, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x48,
	0x61, 0x6c, 0x6f, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x61,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x14, 0x44, 0x61, 0x6e, 0x63, 0x65,
	0x48, 0x61, 0x6c, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x15, 0x44, 0x61, 0x6e, 0x63, 0x65,
	0x48, 0x61, 0x6c, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x6e, 0x63,
	0x65, 0x48, 0x61, 0x6c, 0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x6c,
	0x6f, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x1c, 0x5a,
	0x1a, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6b, 0x74, 0x6f,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_config_proto_goTypes = []any{
	(*DanceSuitListItem)(nil),       // 0: game.proto.DanceSuitListItem
	(*DanceLevelRequest)(nil),       // 1: game.proto.DanceLevelRequest
	(*DanceLevelResponse)(nil),      // 2: game.proto.DanceLevelResponse
	(*DanceElementLocalItem)(nil),   // 3: game.proto.DanceElementLocalItem
	(*DanceHaloButtonRequest)(nil),  // 4: game.proto.DanceHaloButtonRequest
	(*DanceHaloButtonResponse)(nil), // 5: game.proto.DanceHaloButtonResponse
	(*DanceHaloRulesRequest)(nil),   // 6: game.proto.DanceHaloRulesRequest
	(*DanceHaloRulesResponse)(nil),  // 7: game.proto.DanceHaloRulesResponse
	(*DanceHaloRulesItem)(nil),      // 8: game.proto.DanceHaloRulesItem
	(*DanceHaloRankRequest)(nil),    // 9: game.proto.DanceHaloRankRequest
	(*DanceHaloRankResponse)(nil),   // 10: game.proto.DanceHaloRankResponse
	(*DanceHaloRankItem)(nil),       // 11: game.proto.DanceHaloRankItem
}
var file_config_proto_depIdxs = []int32{
	0,  // 0: game.proto.DanceLevelResponse.base_all_suit:type_name -> game.proto.DanceSuitListItem
	3,  // 1: game.proto.DanceLevelResponse.elementList:type_name -> game.proto.DanceElementLocalItem
	8,  // 2: game.proto.DanceHaloRulesResponse.list:type_name -> game.proto.DanceHaloRulesItem
	11, // 3: game.proto.DanceHaloRankResponse.list:type_name -> game.proto.DanceHaloRankItem
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}