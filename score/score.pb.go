// Code generated by protoc-gen-go. DO NOT EDIT.
// source: score.proto

/*
Package score is a generated protocol buffer package.

It is generated from these files:
	score.proto

It has these top-level messages:
	SetMessage
	GetMessage
	ResetMessage
	SetReply
	GetReply
	ResetReply
*/
package score

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// Message
//
type SetMessage struct {
	Nom   string `protobuf:"bytes,1,opt,name=nom" json:"nom,omitempty"`
	Score int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
}

func (m *SetMessage) Reset()                    { *m = SetMessage{} }
func (m *SetMessage) String() string            { return proto.CompactTextString(m) }
func (*SetMessage) ProtoMessage()               {}
func (*SetMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetMessage) GetNom() string {
	if m != nil {
		return m.Nom
	}
	return ""
}

func (m *SetMessage) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type GetMessage struct {
	Nom string `protobuf:"bytes,1,opt,name=nom" json:"nom,omitempty"`
}

func (m *GetMessage) Reset()                    { *m = GetMessage{} }
func (m *GetMessage) String() string            { return proto.CompactTextString(m) }
func (*GetMessage) ProtoMessage()               {}
func (*GetMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetMessage) GetNom() string {
	if m != nil {
		return m.Nom
	}
	return ""
}

type ResetMessage struct {
}

func (m *ResetMessage) Reset()                    { *m = ResetMessage{} }
func (m *ResetMessage) String() string            { return proto.CompactTextString(m) }
func (*ResetMessage) ProtoMessage()               {}
func (*ResetMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SetReply struct {
}

func (m *SetReply) Reset()                    { *m = SetReply{} }
func (m *SetReply) String() string            { return proto.CompactTextString(m) }
func (*SetReply) ProtoMessage()               {}
func (*SetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetReply struct {
	Score int32 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetReply) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type ResetReply struct {
}

func (m *ResetReply) Reset()                    { *m = ResetReply{} }
func (m *ResetReply) String() string            { return proto.CompactTextString(m) }
func (*ResetReply) ProtoMessage()               {}
func (*ResetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*SetMessage)(nil), "score.SetMessage")
	proto.RegisterType((*GetMessage)(nil), "score.GetMessage")
	proto.RegisterType((*ResetMessage)(nil), "score.ResetMessage")
	proto.RegisterType((*SetReply)(nil), "score.SetReply")
	proto.RegisterType((*GetReply)(nil), "score.GetReply")
	proto.RegisterType((*ResetReply)(nil), "score.ResetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Score service

type ScoreClient interface {
	Set(ctx context.Context, in *SetMessage, opts ...grpc.CallOption) (*SetReply, error)
	Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error)
	Reset(ctx context.Context, in *ResetMessage, opts ...grpc.CallOption) (*ResetReply, error)
}

type scoreClient struct {
	cc *grpc.ClientConn
}

func NewScoreClient(cc *grpc.ClientConn) ScoreClient {
	return &scoreClient{cc}
}

func (c *scoreClient) Set(ctx context.Context, in *SetMessage, opts ...grpc.CallOption) (*SetReply, error) {
	out := new(SetReply)
	err := grpc.Invoke(ctx, "/score.Score/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/score.Score/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) Reset(ctx context.Context, in *ResetMessage, opts ...grpc.CallOption) (*ResetReply, error) {
	out := new(ResetReply)
	err := grpc.Invoke(ctx, "/score.Score/Reset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Score service

type ScoreServer interface {
	Set(context.Context, *SetMessage) (*SetReply, error)
	Get(context.Context, *GetMessage) (*GetReply, error)
	Reset(context.Context, *ResetMessage) (*ResetReply, error)
}

func RegisterScoreServer(s *grpc.Server, srv ScoreServer) {
	s.RegisterService(&_Score_serviceDesc, srv)
}

func _Score_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/score.Score/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).Set(ctx, req.(*SetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/score.Score/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).Get(ctx, req.(*GetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/score.Score/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).Reset(ctx, req.(*ResetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Score_serviceDesc = grpc.ServiceDesc{
	ServiceName: "score.Score",
	HandlerType: (*ScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Score_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Score_Get_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Score_Reset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "score.proto",
}

func init() { proto.RegisterFile("score.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4e, 0xce, 0x2f,
	0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x4c, 0xb8, 0xb8, 0x82,
	0x53, 0x4b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x04, 0xb8, 0x98, 0xf3, 0xf2, 0x73,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0x88, 0x42, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa8, 0x2e, 0x39, 0x2e, 0x2e, 0x77, 0x3c, 0xba, 0x94, 0xf8, 0xb8,
	0x78, 0x82, 0x52, 0x8b, 0xe1, 0x2a, 0x94, 0xb8, 0xb8, 0x38, 0x82, 0x53, 0x4b, 0x82, 0x52, 0x0b,
	0x72, 0x2a, 0x95, 0x14, 0xb8, 0x38, 0xdc, 0xa1, 0x6c, 0x84, 0xe9, 0x8c, 0xc8, 0xa6, 0xf3, 0x70,
	0x71, 0x81, 0x75, 0x83, 0xd5, 0x18, 0x4d, 0x61, 0xe4, 0x62, 0x0d, 0x06, 0x89, 0x0b, 0x69, 0x73,
	0x31, 0x07, 0xa7, 0x96, 0x08, 0x09, 0xea, 0x41, 0xfc, 0x81, 0x70, 0xb7, 0x14, 0x3f, 0x42, 0x08,
	0x62, 0x09, 0x03, 0x48, 0xb1, 0x3b, 0x92, 0x62, 0x77, 0x4c, 0xc5, 0xee, 0x08, 0xc5, 0x86, 0x5c,
	0xac, 0x60, 0x1b, 0x85, 0x84, 0xa1, 0x72, 0xc8, 0xae, 0x97, 0x12, 0x44, 0x16, 0x84, 0x6a, 0x49,
	0x62, 0x03, 0x07, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x08, 0x22, 0x26, 0x55, 0x01,
	0x00, 0x00,
}
