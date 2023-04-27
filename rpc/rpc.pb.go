// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: rpc/rpc.proto

package rpc

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Request_Call
	//	*Request_Data
	//	*Request_End
	Type   isRequest_Type `protobuf_oneof:"type"`
	Reply  string         `protobuf:"bytes,5,opt,name=reply,proto3" json:"reply,omitempty"`
	Method string         `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[0]
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
	return file_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (m *Request) GetType() isRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Request) GetCall() *Call {
	if x, ok := x.GetType().(*Request_Call); ok {
		return x.Call
	}
	return nil
}

func (x *Request) GetData() *Data {
	if x, ok := x.GetType().(*Request_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Request) GetEnd() *End {
	if x, ok := x.GetType().(*Request_End); ok {
		return x.End
	}
	return nil
}

func (x *Request) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type isRequest_Type interface {
	isRequest_Type()
}

type Request_Call struct {
	Call *Call `protobuf:"bytes,2,opt,name=call,proto3,oneof"`
}

type Request_Data struct {
	Data *Data `protobuf:"bytes,3,opt,name=data,proto3,oneof"`
}

type Request_End struct {
	End *End `protobuf:"bytes,4,opt,name=end,proto3,oneof"`
}

func (*Request_Call) isRequest_Type() {}

func (*Request_Data) isRequest_Type() {}

func (*Request_End) isRequest_Type() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Response_Begin
	//	*Response_Data
	//	*Response_End
	Type isResponse_Type `protobuf_oneof:"type"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[1]
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
	return file_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetType() isResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Response) GetBegin() *Begin {
	if x, ok := x.GetType().(*Response_Begin); ok {
		return x.Begin
	}
	return nil
}

func (x *Response) GetData() *Data {
	if x, ok := x.GetType().(*Response_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Response) GetEnd() *End {
	if x, ok := x.GetType().(*Response_End); ok {
		return x.End
	}
	return nil
}

type isResponse_Type interface {
	isResponse_Type()
}

type Response_Begin struct {
	Begin *Begin `protobuf:"bytes,2,opt,name=begin,proto3,oneof"`
}

type Response_Data struct {
	Data *Data `protobuf:"bytes,3,opt,name=data,proto3,oneof"`
}

type Response_End struct {
	End *End `protobuf:"bytes,4,opt,name=end,proto3,oneof"`
}

func (*Response_Begin) isResponse_Type() {}

func (*Response_Data) isResponse_Type() {}

func (*Response_End) isResponse_Type() {}

type Strings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Strings) Reset() {
	*x = Strings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strings) ProtoMessage() {}

func (x *Strings) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strings.ProtoReflect.Descriptor instead.
func (*Strings) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Strings) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md map[string]*Strings `protobuf:"bytes,1,rep,name=md,proto3" json:"md,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetMd() map[string]*Strings {
	if x != nil {
		return x.Md
	}
	return nil
}

type Call struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method   string    `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Nid      string    `protobuf:"bytes,3,opt,name=nid,proto3" json:"nid,omitempty"`
	Reply    string    `protobuf:"bytes,4,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *Call) Reset() {
	*x = Call{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call) ProtoMessage() {}

func (x *Call) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call.ProtoReflect.Descriptor instead.
func (*Call) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *Call) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Call) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Call) GetNid() string {
	if x != nil {
		return x.Nid
	}
	return ""
}

func (x *Call) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type Begin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Metadata `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Nid    string    `protobuf:"bytes,2,opt,name=nid,proto3" json:"nid,omitempty"`
}

func (x *Begin) Reset() {
	*x = Begin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Begin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Begin) ProtoMessage() {}

func (x *Begin) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Begin.ProtoReflect.Descriptor instead.
func (*Begin) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *Begin) GetHeader() *Metadata {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Begin) GetNid() string {
	if x != nil {
		return x.Nid
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type End struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Trailer *Metadata      `protobuf:"bytes,2,opt,name=trailer,proto3" json:"trailer,omitempty"`
}

func (x *End) Reset() {
	*x = End{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *End) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*End) ProtoMessage() {}

func (x *End) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use End.ProtoReflect.Descriptor instead.
func (*End) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *End) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *End) GetTrailer() *Metadata {
	if x != nil {
		return x.Trailer
	}
	return nil
}

var File_rpc_rpc_proto protoreflect.FileDescriptor

var file_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1f,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x75, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x76, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x02, 0x6d, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6d, 0x64,
	0x1a, 0x43, 0x0a, 0x07, 0x4d, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x71, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x40, 0x0a, 0x05, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x12, 0x25, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_rpc_proto_rawDescOnce sync.Once
	file_rpc_rpc_proto_rawDescData = file_rpc_rpc_proto_rawDesc
)

func file_rpc_rpc_proto_rawDescGZIP() []byte {
	file_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_rpc_proto_rawDescData)
	})
	return file_rpc_rpc_proto_rawDescData
}

var file_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_rpc_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: rpc.Request
	(*Response)(nil),      // 1: rpc.Response
	(*Strings)(nil),       // 2: rpc.Strings
	(*Metadata)(nil),      // 3: rpc.Metadata
	(*Call)(nil),          // 4: rpc.Call
	(*Begin)(nil),         // 5: rpc.Begin
	(*Data)(nil),          // 6: rpc.Data
	(*End)(nil),           // 7: rpc.End
	nil,                   // 8: rpc.Metadata.MdEntry
	(*status.Status)(nil), // 9: google.rpc.Status
}
var file_rpc_rpc_proto_depIdxs = []int32{
	4,  // 0: rpc.Request.call:type_name -> rpc.Call
	6,  // 1: rpc.Request.data:type_name -> rpc.Data
	7,  // 2: rpc.Request.end:type_name -> rpc.End
	5,  // 3: rpc.Response.begin:type_name -> rpc.Begin
	6,  // 4: rpc.Response.data:type_name -> rpc.Data
	7,  // 5: rpc.Response.end:type_name -> rpc.End
	8,  // 6: rpc.Metadata.md:type_name -> rpc.Metadata.MdEntry
	3,  // 7: rpc.Call.metadata:type_name -> rpc.Metadata
	3,  // 8: rpc.Begin.header:type_name -> rpc.Metadata
	9,  // 9: rpc.End.status:type_name -> google.rpc.Status
	3,  // 10: rpc.End.trailer:type_name -> rpc.Metadata
	2,  // 11: rpc.Metadata.MdEntry.value:type_name -> rpc.Strings
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_rpc_rpc_proto_init() }
func file_rpc_rpc_proto_init() {
	if File_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strings); i {
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
		file_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_rpc_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call); i {
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
		file_rpc_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Begin); i {
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
		file_rpc_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_rpc_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*End); i {
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
	file_rpc_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_Call)(nil),
		(*Request_Data)(nil),
		(*Request_End)(nil),
	}
	file_rpc_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_Begin)(nil),
		(*Response_Data)(nil),
		(*Response_End)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_rpc_proto_msgTypes,
	}.Build()
	File_rpc_rpc_proto = out.File
	file_rpc_rpc_proto_rawDesc = nil
	file_rpc_rpc_proto_goTypes = nil
	file_rpc_rpc_proto_depIdxs = nil
}
