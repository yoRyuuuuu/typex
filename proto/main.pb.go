// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/main.proto

package proto

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

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Player []*Player `protobuf:"bytes,2,rep,name=player,proto3" json:"player,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnectResponse) GetPlayer() []*Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problem string `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
}

func (x *Start) Reset() {
	*x = Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Start) ProtoMessage() {}

func (x *Start) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Start.ProtoReflect.Descriptor instead.
func (*Start) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *Start) GetProblem() string {
	if x != nil {
		return x.Problem
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Request_Answer
	Action isRequest_Action `protobuf_oneof:"action"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
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
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Request) GetAnswer() *Answer {
	if x, ok := x.GetAction().(*Request_Answer); ok {
		return x.Answer
	}
	return nil
}

type isRequest_Action interface {
	isRequest_Action()
}

type Request_Answer struct {
	Answer *Answer `protobuf:"bytes,1,opt,name=answer,proto3,oneof"`
}

func (*Request_Answer) isRequest_Action() {}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Finish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *Finish) Reset() {
	*x = Finish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finish) ProtoMessage() {}

func (x *Finish) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finish.ProtoReflect.Descriptor instead.
func (*Finish) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{7}
}

func (x *Finish) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{8}
}

func (x *Join) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Attack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 攻撃を受けるPlayerのID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 体力の数値
	Health int64 `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *Attack) Reset() {
	*x = Attack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attack) ProtoMessage() {}

func (x *Attack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attack.ProtoReflect.Descriptor instead.
func (*Attack) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{9}
}

func (x *Attack) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attack) GetHealth() int64 {
	if x != nil {
		return x.Health
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*Response_Question
	//	*Response_Start
	//	*Response_Finish
	//	*Response_Join
	//	*Response_Attack
	Action isResponse_Action `protobuf_oneof:"action"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[10]
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
	return file_proto_main_proto_rawDescGZIP(), []int{10}
}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Response) GetQuestion() *Question {
	if x, ok := x.GetAction().(*Response_Question); ok {
		return x.Question
	}
	return nil
}

func (x *Response) GetStart() *Start {
	if x, ok := x.GetAction().(*Response_Start); ok {
		return x.Start
	}
	return nil
}

func (x *Response) GetFinish() *Finish {
	if x, ok := x.GetAction().(*Response_Finish); ok {
		return x.Finish
	}
	return nil
}

func (x *Response) GetJoin() *Join {
	if x, ok := x.GetAction().(*Response_Join); ok {
		return x.Join
	}
	return nil
}

func (x *Response) GetAttack() *Attack {
	if x, ok := x.GetAction().(*Response_Attack); ok {
		return x.Attack
	}
	return nil
}

type isResponse_Action interface {
	isResponse_Action()
}

type Response_Question struct {
	Question *Question `protobuf:"bytes,1,opt,name=question,proto3,oneof"`
}

type Response_Start struct {
	Start *Start `protobuf:"bytes,2,opt,name=start,proto3,oneof"`
}

type Response_Finish struct {
	Finish *Finish `protobuf:"bytes,3,opt,name=finish,proto3,oneof"`
}

type Response_Join struct {
	Join *Join `protobuf:"bytes,4,opt,name=join,proto3,oneof"`
}

type Response_Attack struct {
	Attack *Attack `protobuf:"bytes,5,opt,name=attack,proto3,oneof"`
}

func (*Response_Question) isResponse_Action() {}

func (*Response_Start) isResponse_Action() {}

func (*Response_Finish) isResponse_Action() {}

func (*Response_Join) isResponse_Action() {}

func (*Response_Attack) isResponse_Action() {}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x08, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x08,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x21, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x22,
	0x36, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x30, 0x0a, 0x06, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x04, 0x6a,
	0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x48, 0x00, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5b, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x6f, 0x52, 0x79, 0x75, 0x75, 0x75, 0x75, 0x75, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x78,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_main_proto_goTypes = []interface{}{
	(*ConnectRequest)(nil),  // 0: ConnectRequest
	(*ConnectResponse)(nil), // 1: ConnectResponse
	(*Answer)(nil),          // 2: Answer
	(*Question)(nil),        // 3: Question
	(*Start)(nil),           // 4: Start
	(*Request)(nil),         // 5: Request
	(*Player)(nil),          // 6: Player
	(*Finish)(nil),          // 7: Finish
	(*Join)(nil),            // 8: Join
	(*Attack)(nil),          // 9: Attack
	(*Response)(nil),        // 10: Response
}
var file_proto_main_proto_depIdxs = []int32{
	6,  // 0: ConnectResponse.player:type_name -> Player
	2,  // 1: Request.answer:type_name -> Answer
	6,  // 2: Join.player:type_name -> Player
	3,  // 3: Response.question:type_name -> Question
	4,  // 4: Response.start:type_name -> Start
	7,  // 5: Response.finish:type_name -> Finish
	8,  // 6: Response.join:type_name -> Join
	9,  // 7: Response.attack:type_name -> Attack
	0,  // 8: Game.Connect:input_type -> ConnectRequest
	5,  // 9: Game.Stream:input_type -> Request
	1,  // 10: Game.Connect:output_type -> ConnectResponse
	10, // 11: Game.Stream:output_type -> Response
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_proto_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_proto_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_proto_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_proto_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Start); i {
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
		file_proto_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_main_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finish); i {
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
		file_proto_main_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
		file_proto_main_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attack); i {
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
		file_proto_main_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_proto_main_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Request_Answer)(nil),
	}
	file_proto_main_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*Response_Question)(nil),
		(*Response_Start)(nil),
		(*Response_Finish)(nil),
		(*Response_Join)(nil),
		(*Response_Attack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}
