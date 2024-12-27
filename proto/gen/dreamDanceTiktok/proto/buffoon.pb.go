// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: buffoon.proto

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

//获取小丑状态协议
type GetBuffoonStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ModelId       int64                  `protobuf:"varint,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"` //模式id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBuffoonStatusRequest) Reset() {
	*x = GetBuffoonStatusRequest{}
	mi := &file_buffoon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBuffoonStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuffoonStatusRequest) ProtoMessage() {}

func (x *GetBuffoonStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuffoonStatusRequest.ProtoReflect.Descriptor instead.
func (*GetBuffoonStatusRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{0}
}

func (x *GetBuffoonStatusRequest) GetModelId() int64 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

type GetBuffoonStatusResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	CanUseBuffoon      int32                  `protobuf:"varint,1,opt,name=can_use_buffoon,json=canUseBuffoon,proto3" json:"can_use_buffoon,omitempty"`                  //1:小丑活动下架;2:可以召唤小丑;3:召唤过，不能再次召唤
	FightBuffoonId     string                 `protobuf:"bytes,2,opt,name=fight_buffoon_id,json=fightBuffoonId,proto3" json:"fight_buffoon_id,omitempty"`                //打小丑的id
	LastGetBuffoonTime int64                  `protobuf:"varint,3,opt,name=last_get_buffoon_time,json=lastGetBuffoonTime,proto3" json:"last_get_buffoon_time,omitempty"` //上次召唤小丑的时间
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetBuffoonStatusResponse) Reset() {
	*x = GetBuffoonStatusResponse{}
	mi := &file_buffoon_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBuffoonStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuffoonStatusResponse) ProtoMessage() {}

func (x *GetBuffoonStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuffoonStatusResponse.ProtoReflect.Descriptor instead.
func (*GetBuffoonStatusResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{1}
}

func (x *GetBuffoonStatusResponse) GetCanUseBuffoon() int32 {
	if x != nil {
		return x.CanUseBuffoon
	}
	return 0
}

func (x *GetBuffoonStatusResponse) GetFightBuffoonId() string {
	if x != nil {
		return x.FightBuffoonId
	}
	return ""
}

func (x *GetBuffoonStatusResponse) GetLastGetBuffoonTime() int64 {
	if x != nil {
		return x.LastGetBuffoonTime
	}
	return 0
}

//今天的用户排行榜
type TodayBuffoonRankRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` //前多少名
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodayBuffoonRankRequest) Reset() {
	*x = TodayBuffoonRankRequest{}
	mi := &file_buffoon_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodayBuffoonRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayBuffoonRankRequest) ProtoMessage() {}

func (x *TodayBuffoonRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayBuffoonRankRequest.ProtoReflect.Descriptor instead.
func (*TodayBuffoonRankRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{2}
}

func (x *TodayBuffoonRankRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

//获取今天排名前多少的用户
type TodayBuffoonRankResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	List          []*TodayBuffoonRankItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodayBuffoonRankResponse) Reset() {
	*x = TodayBuffoonRankResponse{}
	mi := &file_buffoon_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodayBuffoonRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayBuffoonRankResponse) ProtoMessage() {}

func (x *TodayBuffoonRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayBuffoonRankResponse.ProtoReflect.Descriptor instead.
func (*TodayBuffoonRankResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{3}
}

func (x *TodayBuffoonRankResponse) GetList() []*TodayBuffoonRankItem {
	if x != nil {
		return x.List
	}
	return nil
}

type TodayBuffoonRankItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalScore    int64                  `protobuf:"varint,1,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"` //用户累计积分
	WorldRank     int32                  `protobuf:"varint,2,opt,name=world_rank,json=worldRank,proto3" json:"world_rank,omitempty"`    //用户的排名
	Nickname      string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`                        //用户昵称
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                            //头像
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodayBuffoonRankItem) Reset() {
	*x = TodayBuffoonRankItem{}
	mi := &file_buffoon_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodayBuffoonRankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayBuffoonRankItem) ProtoMessage() {}

func (x *TodayBuffoonRankItem) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayBuffoonRankItem.ProtoReflect.Descriptor instead.
func (*TodayBuffoonRankItem) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{4}
}

func (x *TodayBuffoonRankItem) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *TodayBuffoonRankItem) GetWorldRank() int32 {
	if x != nil {
		return x.WorldRank
	}
	return 0
}

func (x *TodayBuffoonRankItem) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *TodayBuffoonRankItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

//昨天用户瓜分前百分之10的榜单
type SplitBuffoonRankRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"` //前多少名
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SplitBuffoonRankRequest) Reset() {
	*x = SplitBuffoonRankRequest{}
	mi := &file_buffoon_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SplitBuffoonRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitBuffoonRankRequest) ProtoMessage() {}

func (x *SplitBuffoonRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitBuffoonRankRequest.ProtoReflect.Descriptor instead.
func (*SplitBuffoonRankRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{5}
}

func (x *SplitBuffoonRankRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SplitBuffoonRankResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	List          []*SplitBuffoonRankItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SplitBuffoonRankResponse) Reset() {
	*x = SplitBuffoonRankResponse{}
	mi := &file_buffoon_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SplitBuffoonRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitBuffoonRankResponse) ProtoMessage() {}

func (x *SplitBuffoonRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitBuffoonRankResponse.ProtoReflect.Descriptor instead.
func (*SplitBuffoonRankResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{6}
}

func (x *SplitBuffoonRankResponse) GetList() []*SplitBuffoonRankItem {
	if x != nil {
		return x.List
	}
	return nil
}

type SplitBuffoonRankItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalScore    int64                  `protobuf:"varint,1,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"` //用户累计积分
	WorldRank     int32                  `protobuf:"varint,2,opt,name=world_rank,json=worldRank,proto3" json:"world_rank,omitempty"`    //用户的排名
	Nickname      string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`                        //用户昵称
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                            //头像
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SplitBuffoonRankItem) Reset() {
	*x = SplitBuffoonRankItem{}
	mi := &file_buffoon_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SplitBuffoonRankItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitBuffoonRankItem) ProtoMessage() {}

func (x *SplitBuffoonRankItem) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitBuffoonRankItem.ProtoReflect.Descriptor instead.
func (*SplitBuffoonRankItem) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{7}
}

func (x *SplitBuffoonRankItem) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *SplitBuffoonRankItem) GetWorldRank() int32 {
	if x != nil {
		return x.WorldRank
	}
	return 0
}

func (x *SplitBuffoonRankItem) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SplitBuffoonRankItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

//今天小丑总积分
type TodayBuffoonTotalScoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodayBuffoonTotalScoreRequest) Reset() {
	*x = TodayBuffoonTotalScoreRequest{}
	mi := &file_buffoon_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodayBuffoonTotalScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayBuffoonTotalScoreRequest) ProtoMessage() {}

func (x *TodayBuffoonTotalScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayBuffoonTotalScoreRequest.ProtoReflect.Descriptor instead.
func (*TodayBuffoonTotalScoreRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{8}
}

type TodayBuffoonTotalScoreResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalScore    int64                  `protobuf:"varint,1,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodayBuffoonTotalScoreResponse) Reset() {
	*x = TodayBuffoonTotalScoreResponse{}
	mi := &file_buffoon_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodayBuffoonTotalScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayBuffoonTotalScoreResponse) ProtoMessage() {}

func (x *TodayBuffoonTotalScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayBuffoonTotalScoreResponse.ProtoReflect.Descriptor instead.
func (*TodayBuffoonTotalScoreResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{9}
}

func (x *TodayBuffoonTotalScoreResponse) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

//昨天小丑总积分
type YesterdayBuffoonTotalScoreRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YesterdayBuffoonTotalScoreRequest) Reset() {
	*x = YesterdayBuffoonTotalScoreRequest{}
	mi := &file_buffoon_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YesterdayBuffoonTotalScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YesterdayBuffoonTotalScoreRequest) ProtoMessage() {}

func (x *YesterdayBuffoonTotalScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YesterdayBuffoonTotalScoreRequest.ProtoReflect.Descriptor instead.
func (*YesterdayBuffoonTotalScoreRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{10}
}

type YesterdayBuffoonTotalScoreResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TotalScore    int64                  `protobuf:"varint,1,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YesterdayBuffoonTotalScoreResponse) Reset() {
	*x = YesterdayBuffoonTotalScoreResponse{}
	mi := &file_buffoon_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YesterdayBuffoonTotalScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YesterdayBuffoonTotalScoreResponse) ProtoMessage() {}

func (x *YesterdayBuffoonTotalScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YesterdayBuffoonTotalScoreResponse.ProtoReflect.Descriptor instead.
func (*YesterdayBuffoonTotalScoreResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{11}
}

func (x *YesterdayBuffoonTotalScoreResponse) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

//昨天小丑总积分
type BeginBuffoonRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FightBuffoonId string                 `protobuf:"bytes,1,opt,name=fight_buffoon_id,json=fightBuffoonId,proto3" json:"fight_buffoon_id,omitempty"` //打小丑的id
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BeginBuffoonRequest) Reset() {
	*x = BeginBuffoonRequest{}
	mi := &file_buffoon_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginBuffoonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginBuffoonRequest) ProtoMessage() {}

func (x *BeginBuffoonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginBuffoonRequest.ProtoReflect.Descriptor instead.
func (*BeginBuffoonRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{12}
}

func (x *BeginBuffoonRequest) GetFightBuffoonId() string {
	if x != nil {
		return x.FightBuffoonId
	}
	return ""
}

type BeginBuffoonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BeginBuffoonResponse) Reset() {
	*x = BeginBuffoonResponse{}
	mi := &file_buffoon_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BeginBuffoonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginBuffoonResponse) ProtoMessage() {}

func (x *BeginBuffoonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginBuffoonResponse.ProtoReflect.Descriptor instead.
func (*BeginBuffoonResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{13}
}

//昨天小丑总积分
type EndBuffoonRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FightBuffoonId string                 `protobuf:"bytes,1,opt,name=fight_buffoon_id,json=fightBuffoonId,proto3" json:"fight_buffoon_id,omitempty"` //打小丑的id
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EndBuffoonRequest) Reset() {
	*x = EndBuffoonRequest{}
	mi := &file_buffoon_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndBuffoonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBuffoonRequest) ProtoMessage() {}

func (x *EndBuffoonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBuffoonRequest.ProtoReflect.Descriptor instead.
func (*EndBuffoonRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{14}
}

func (x *EndBuffoonRequest) GetFightBuffoonId() string {
	if x != nil {
		return x.FightBuffoonId
	}
	return ""
}

type EndBuffoonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EndBuffoonResponse) Reset() {
	*x = EndBuffoonResponse{}
	mi := &file_buffoon_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndBuffoonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBuffoonResponse) ProtoMessage() {}

func (x *EndBuffoonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBuffoonResponse.ProtoReflect.Descriptor instead.
func (*EndBuffoonResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{15}
}

//小丑前10分钟的气球
type BalloonsHeadEndRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                    //小丑的id
	TotalScore    int64                  `protobuf:"varint,2,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"` //小丑的总分数
	List          []*BalloonsHeadEndItem `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`                                //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalloonsHeadEndRequest) Reset() {
	*x = BalloonsHeadEndRequest{}
	mi := &file_buffoon_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalloonsHeadEndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalloonsHeadEndRequest) ProtoMessage() {}

func (x *BalloonsHeadEndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalloonsHeadEndRequest.ProtoReflect.Descriptor instead.
func (*BalloonsHeadEndRequest) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{16}
}

func (x *BalloonsHeadEndRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BalloonsHeadEndRequest) GetTotalScore() int64 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *BalloonsHeadEndRequest) GetList() []*BalloonsHeadEndItem {
	if x != nil {
		return x.List
	}
	return nil
}

type BalloonsHeadEndItem struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	UserId           string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                  //玩家id
	TotalScoreBattle int64                  `protobuf:"varint,2,opt,name=total_score_battle,json=totalScoreBattle,proto3" json:"total_score_battle,omitempty"` //礼物分
	TotalScoreEnd    int64                  `protobuf:"varint,3,opt,name=total_score_end,json=totalScoreEnd,proto3" json:"total_score_end,omitempty"`          //礼物算出的结算分
	Rank             int32                  `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`                                                   //排名
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *BalloonsHeadEndItem) Reset() {
	*x = BalloonsHeadEndItem{}
	mi := &file_buffoon_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalloonsHeadEndItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalloonsHeadEndItem) ProtoMessage() {}

func (x *BalloonsHeadEndItem) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalloonsHeadEndItem.ProtoReflect.Descriptor instead.
func (*BalloonsHeadEndItem) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{17}
}

func (x *BalloonsHeadEndItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BalloonsHeadEndItem) GetTotalScoreBattle() int64 {
	if x != nil {
		return x.TotalScoreBattle
	}
	return 0
}

func (x *BalloonsHeadEndItem) GetTotalScoreEnd() int64 {
	if x != nil {
		return x.TotalScoreEnd
	}
	return 0
}

func (x *BalloonsHeadEndItem) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type BalloonsHeadEndResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalloonsHeadEndResponse) Reset() {
	*x = BalloonsHeadEndResponse{}
	mi := &file_buffoon_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalloonsHeadEndResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalloonsHeadEndResponse) ProtoMessage() {}

func (x *BalloonsHeadEndResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buffoon_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalloonsHeadEndResponse.ProtoReflect.Descriptor instead.
func (*BalloonsHeadEndResponse) Descriptor() ([]byte, []int) {
	return file_buffoon_proto_rawDescGZIP(), []int{18}
}

var File_buffoon_proto protoreflect.FileDescriptor

var file_buffoon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x22, 0x9f, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x6f, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x62, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x75, 0x66,
	0x66, 0x6f, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x12, 0x6c, 0x61, 0x73, 0x74, 0x47, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x42, 0x75, 0x66, 0x66,
	0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x18, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x42,
	0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x61, 0x79, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x54, 0x6f, 0x64,
	0x61, 0x79, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x61, 0x6e,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x31, 0x0a, 0x17, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x42, 0x75,
	0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x18, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x53,
	0x70, 0x6c, 0x69, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x61, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x1f, 0x0a, 0x1d, 0x54, 0x6f, 0x64, 0x61, 0x79,
	0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x1e, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x21, 0x59,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x45, 0x0a, 0x22, 0x59, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x42, 0x75, 0x66,
	0x66, 0x6f, 0x6f, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x42,
	0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3d, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x42, 0x75, 0x66, 0x66, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x0a, 0x16, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e,
	0x73, 0x48, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6c, 0x6c,
	0x6f, 0x6f, 0x6e, 0x73, 0x48, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f,
	0x6e, 0x73, 0x48, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x19, 0x0a, 0x17, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x73, 0x48, 0x65, 0x61, 0x64,
	0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_buffoon_proto_rawDescOnce sync.Once
	file_buffoon_proto_rawDescData = file_buffoon_proto_rawDesc
)

func file_buffoon_proto_rawDescGZIP() []byte {
	file_buffoon_proto_rawDescOnce.Do(func() {
		file_buffoon_proto_rawDescData = protoimpl.X.CompressGZIP(file_buffoon_proto_rawDescData)
	})
	return file_buffoon_proto_rawDescData
}

var file_buffoon_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_buffoon_proto_goTypes = []any{
	(*GetBuffoonStatusRequest)(nil),            // 0: game.proto.GetBuffoonStatusRequest
	(*GetBuffoonStatusResponse)(nil),           // 1: game.proto.GetBuffoonStatusResponse
	(*TodayBuffoonRankRequest)(nil),            // 2: game.proto.TodayBuffoonRankRequest
	(*TodayBuffoonRankResponse)(nil),           // 3: game.proto.TodayBuffoonRankResponse
	(*TodayBuffoonRankItem)(nil),               // 4: game.proto.TodayBuffoonRankItem
	(*SplitBuffoonRankRequest)(nil),            // 5: game.proto.SplitBuffoonRankRequest
	(*SplitBuffoonRankResponse)(nil),           // 6: game.proto.SplitBuffoonRankResponse
	(*SplitBuffoonRankItem)(nil),               // 7: game.proto.SplitBuffoonRankItem
	(*TodayBuffoonTotalScoreRequest)(nil),      // 8: game.proto.TodayBuffoonTotalScoreRequest
	(*TodayBuffoonTotalScoreResponse)(nil),     // 9: game.proto.TodayBuffoonTotalScoreResponse
	(*YesterdayBuffoonTotalScoreRequest)(nil),  // 10: game.proto.YesterdayBuffoonTotalScoreRequest
	(*YesterdayBuffoonTotalScoreResponse)(nil), // 11: game.proto.YesterdayBuffoonTotalScoreResponse
	(*BeginBuffoonRequest)(nil),                // 12: game.proto.BeginBuffoonRequest
	(*BeginBuffoonResponse)(nil),               // 13: game.proto.BeginBuffoonResponse
	(*EndBuffoonRequest)(nil),                  // 14: game.proto.EndBuffoonRequest
	(*EndBuffoonResponse)(nil),                 // 15: game.proto.EndBuffoonResponse
	(*BalloonsHeadEndRequest)(nil),             // 16: game.proto.BalloonsHeadEndRequest
	(*BalloonsHeadEndItem)(nil),                // 17: game.proto.BalloonsHeadEndItem
	(*BalloonsHeadEndResponse)(nil),            // 18: game.proto.BalloonsHeadEndResponse
}
var file_buffoon_proto_depIdxs = []int32{
	4,  // 0: game.proto.TodayBuffoonRankResponse.list:type_name -> game.proto.TodayBuffoonRankItem
	7,  // 1: game.proto.SplitBuffoonRankResponse.list:type_name -> game.proto.SplitBuffoonRankItem
	17, // 2: game.proto.BalloonsHeadEndRequest.list:type_name -> game.proto.BalloonsHeadEndItem
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_buffoon_proto_init() }
func file_buffoon_proto_init() {
	if File_buffoon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buffoon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buffoon_proto_goTypes,
		DependencyIndexes: file_buffoon_proto_depIdxs,
		MessageInfos:      file_buffoon_proto_msgTypes,
	}.Build()
	File_buffoon_proto = out.File
	file_buffoon_proto_rawDesc = nil
	file_buffoon_proto_goTypes = nil
	file_buffoon_proto_depIdxs = nil
}