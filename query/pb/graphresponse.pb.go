// Code generated by protoc-gen-go.
// source: graphresponse.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	graphresponse.proto

It has these top-level messages:
	UidList
	Result
	GraphRequest
	GraphResponse
*/
package pb

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
const _ = proto.ProtoPackageIsVersion1

type UidList struct {
	Uids []uint64 `protobuf:"varint,1,rep,name=uids" json:"uids,omitempty"`
}

func (m *UidList) Reset()                    { *m = UidList{} }
func (m *UidList) String() string            { return proto.CompactTextString(m) }
func (*UidList) ProtoMessage()               {}
func (*UidList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Values    [][]byte   `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Uidmatrix []*UidList `protobuf:"bytes,2,rep,name=uidmatrix" json:"uidmatrix,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetUidmatrix() []*UidList {
	if m != nil {
		return m.Uidmatrix
	}
	return nil
}

type GraphRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GraphRequest) Reset()                    { *m = GraphRequest{} }
func (m *GraphRequest) String() string            { return proto.CompactTextString(m) }
func (*GraphRequest) ProtoMessage()               {}
func (*GraphRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GraphResponse struct {
	Attribute string           `protobuf:"bytes,1,opt,name=attribute" json:"attribute,omitempty"`
	Result    *Result          `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Children  []*GraphResponse `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
}

func (m *GraphResponse) Reset()                    { *m = GraphResponse{} }
func (m *GraphResponse) String() string            { return proto.CompactTextString(m) }
func (*GraphResponse) ProtoMessage()               {}
func (*GraphResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GraphResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GraphResponse) GetChildren() []*GraphResponse {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*UidList)(nil), "pb.UidList")
	proto.RegisterType((*Result)(nil), "pb.Result")
	proto.RegisterType((*GraphRequest)(nil), "pb.GraphRequest")
	proto.RegisterType((*GraphResponse)(nil), "pb.GraphResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for DGraph service

type DGraphClient interface {
	GetResponse(ctx context.Context, in *GraphRequest, opts ...grpc.CallOption) (*GraphResponse, error)
}

type dGraphClient struct {
	cc *grpc.ClientConn
}

func NewDGraphClient(cc *grpc.ClientConn) DGraphClient {
	return &dGraphClient{cc}
}

func (c *dGraphClient) GetResponse(ctx context.Context, in *GraphRequest, opts ...grpc.CallOption) (*GraphResponse, error) {
	out := new(GraphResponse)
	err := grpc.Invoke(ctx, "/pb.DGraph/GetResponse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DGraph service

type DGraphServer interface {
	GetResponse(context.Context, *GraphRequest) (*GraphResponse, error)
}

func RegisterDGraphServer(s *grpc.Server, srv DGraphServer) {
	s.RegisterService(&_DGraph_serviceDesc, srv)
}

func _DGraph_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DGraphServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DGraph/GetResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DGraphServer).GetResponse(ctx, req.(*GraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DGraph_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DGraph",
	HandlerType: (*DGraphServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResponse",
			Handler:    _DGraph_GetResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xed, 0xee, 0x1a, 0xed, 0x74, 0x05, 0x1d, 0x45, 0x8a, 0x28, 0x48, 0xf0, 0xb0, 0x1e,
	0xec, 0x61, 0xf5, 0xec, 0x49, 0xd8, 0x83, 0x9e, 0x02, 0xfe, 0x80, 0xd6, 0x06, 0x37, 0x50, 0xb7,
	0x31, 0x99, 0x8a, 0xde, 0xfc, 0xe9, 0x4e, 0xd3, 0xb0, 0x55, 0xbc, 0xcd, 0xf4, 0xbd, 0x79, 0xef,
	0x6b, 0xe0, 0xf8, 0xd5, 0x95, 0x76, 0xed, 0xb4, 0xb7, 0xed, 0xc6, 0xeb, 0xc2, 0xba, 0x96, 0x5a,
	0x9c, 0xd8, 0x4a, 0x5e, 0xc0, 0xde, 0xb3, 0xa9, 0x9f, 0x8c, 0x27, 0x44, 0x98, 0x75, 0xa6, 0xf6,
	0x79, 0x72, 0x39, 0x5d, 0xcc, 0x54, 0x98, 0xe5, 0x23, 0x08, 0xa5, 0x7d, 0xd7, 0x10, 0x9e, 0x82,
	0xf8, 0x28, 0x9b, 0x4e, 0x0f, 0xfa, 0x5c, 0xc5, 0x0d, 0xaf, 0x21, 0x65, 0xe7, 0x5b, 0x49, 0xce,
	0x7c, 0xe6, 0x13, 0x96, 0xb2, 0x65, 0x56, 0xd8, 0xaa, 0x88, 0xa9, 0x6a, 0x54, 0xe5, 0x15, 0xcc,
	0x57, 0x3d, 0x86, 0xd2, 0xef, 0x7c, 0x49, 0x78, 0x02, 0xbb, 0x3c, 0xb8, 0x2f, 0x4e, 0x4c, 0x16,
	0xa9, 0x1a, 0x16, 0xf9, 0x9d, 0xc0, 0x41, 0xb4, 0x0d, 0xb4, 0x78, 0x0e, 0x69, 0x49, 0x9c, 0x50,
	0x75, 0xa4, 0xa3, 0x77, 0xfc, 0x80, 0x12, 0x84, 0x0b, 0x88, 0xdc, 0x9e, 0x70, 0x3b, 0xf4, 0xed,
	0x03, 0xb4, 0x8a, 0x0a, 0xde, 0xc0, 0xfe, 0xcb, 0xda, 0x34, 0xb5, 0xd3, 0x9b, 0x7c, 0x1a, 0x18,
	0x8f, 0x7a, 0xd7, 0x9f, 0x1a, 0xb5, 0xb5, 0x2c, 0xef, 0x41, 0x3c, 0x04, 0x0d, 0xef, 0x20, 0x5b,
	0x69, 0xda, 0x92, 0x1c, 0xfe, 0xba, 0x0a, 0xff, 0x70, 0xf6, 0x3f, 0x47, 0xee, 0x54, 0x22, 0xbc,
	0xef, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xab, 0xec, 0x88, 0x76, 0x01, 0x00, 0x00,
}
