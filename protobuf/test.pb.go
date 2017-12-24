// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	SongRequest
	SongReply
	MakeAWishRequest
	MakeAWishReply
*/
package protobuf

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

type SongRequest struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
}

func (m *SongRequest) Reset()                    { *m = SongRequest{} }
func (m *SongRequest) String() string            { return proto.CompactTextString(m) }
func (*SongRequest) ProtoMessage()               {}
func (*SongRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SongRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type SongReply struct {
	Lyric  string `protobuf:"bytes,1,opt,name=lyric" json:"lyric,omitempty"`
	Singer string `protobuf:"bytes,2,opt,name=singer" json:"singer,omitempty"`
	Year   uint32 `protobuf:"varint,3,opt,name=year" json:"year,omitempty"`
}

func (m *SongReply) Reset()                    { *m = SongReply{} }
func (m *SongReply) String() string            { return proto.CompactTextString(m) }
func (*SongReply) ProtoMessage()               {}
func (*SongReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SongReply) GetLyric() string {
	if m != nil {
		return m.Lyric
	}
	return ""
}

func (m *SongReply) GetSinger() string {
	if m != nil {
		return m.Singer
	}
	return ""
}

func (m *SongReply) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

type MakeAWishRequest struct {
	From    string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *MakeAWishRequest) Reset()                    { *m = MakeAWishRequest{} }
func (m *MakeAWishRequest) String() string            { return proto.CompactTextString(m) }
func (*MakeAWishRequest) ProtoMessage()               {}
func (*MakeAWishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MakeAWishRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MakeAWishRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *MakeAWishRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MakeAWishReply struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *MakeAWishReply) Reset()                    { *m = MakeAWishReply{} }
func (m *MakeAWishReply) String() string            { return proto.CompactTextString(m) }
func (*MakeAWishReply) ProtoMessage()               {}
func (*MakeAWishReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MakeAWishReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*SongRequest)(nil), "protobuf.SongRequest")
	proto.RegisterType((*SongReply)(nil), "protobuf.SongReply")
	proto.RegisterType((*MakeAWishRequest)(nil), "protobuf.MakeAWishRequest")
	proto.RegisterType((*MakeAWishReply)(nil), "protobuf.MakeAWishReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Action service

type ActionClient interface {
	Song(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongReply, error)
	MakeAWish(ctx context.Context, in *MakeAWishRequest, opts ...grpc.CallOption) (Action_MakeAWishClient, error)
}

type actionClient struct {
	cc *grpc.ClientConn
}

func NewActionClient(cc *grpc.ClientConn) ActionClient {
	return &actionClient{cc}
}

func (c *actionClient) Song(ctx context.Context, in *SongRequest, opts ...grpc.CallOption) (*SongReply, error) {
	out := new(SongReply)
	err := grpc.Invoke(ctx, "/protobuf.Action/Song", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionClient) MakeAWish(ctx context.Context, in *MakeAWishRequest, opts ...grpc.CallOption) (Action_MakeAWishClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Action_serviceDesc.Streams[0], c.cc, "/protobuf.Action/MakeAWish", opts...)
	if err != nil {
		return nil, err
	}
	x := &actionMakeAWishClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Action_MakeAWishClient interface {
	Recv() (*MakeAWishReply, error)
	grpc.ClientStream
}

type actionMakeAWishClient struct {
	grpc.ClientStream
}

func (x *actionMakeAWishClient) Recv() (*MakeAWishReply, error) {
	m := new(MakeAWishReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Action service

type ActionServer interface {
	Song(context.Context, *SongRequest) (*SongReply, error)
	MakeAWish(*MakeAWishRequest, Action_MakeAWishServer) error
}

func RegisterActionServer(s *grpc.Server, srv ActionServer) {
	s.RegisterService(&_Action_serviceDesc, srv)
}

func _Action_Song_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServer).Song(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Action/Song",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServer).Song(ctx, req.(*SongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Action_MakeAWish_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MakeAWishRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServer).MakeAWish(m, &actionMakeAWishServer{stream})
}

type Action_MakeAWishServer interface {
	Send(*MakeAWishReply) error
	grpc.ServerStream
}

type actionMakeAWishServer struct {
	grpc.ServerStream
}

func (x *actionMakeAWishServer) Send(m *MakeAWishReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Action_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Action",
	HandlerType: (*ActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Song",
			Handler:    _Action_Song_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeAWish",
			Handler:       _Action_MakeAWish_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x71, 0x29, 0x81, 0x1c, 0xa2, 0x42, 0xc7, 0x1f, 0x59, 0x99, 0x90, 0x59, 0x10, 0x43,
	0x84, 0x80, 0x17, 0xe8, 0xc0, 0x58, 0x09, 0x99, 0x81, 0xb9, 0x8d, 0xdc, 0x62, 0x11, 0xec, 0xe2,
	0x5c, 0x87, 0x3c, 0x00, 0xef, 0x8d, 0x73, 0x8d, 0x21, 0x54, 0x4c, 0xbe, 0xef, 0xbe, 0x9f, 0xbe,
	0xf3, 0x1d, 0x00, 0x99, 0x86, 0xca, 0x75, 0xf0, 0xe4, 0xf1, 0x88, 0x9f, 0xc5, 0x66, 0xa9, 0xae,
	0xe1, 0xf8, 0xc5, 0xbb, 0x95, 0x36, 0x9f, 0x9b, 0x68, 0xe3, 0x39, 0x1c, 0x90, 0xa5, 0xda, 0x48,
	0x71, 0x25, 0x6e, 0x72, 0xbd, 0x15, 0x6a, 0x06, 0xf9, 0x16, 0x5a, 0xd7, 0x6d, 0x87, 0xd4, 0x6d,
	0xb0, 0x55, 0x42, 0x58, 0xe0, 0x25, 0x64, 0x8d, 0x75, 0x2b, 0x13, 0xe4, 0x88, 0xdb, 0xbd, 0x42,
	0x84, 0x71, 0x6b, 0xe6, 0x41, 0xee, 0xc7, 0xee, 0x89, 0xe6, 0x5a, 0x3d, 0xc3, 0xe9, 0x6c, 0xfe,
	0x6e, 0xa6, 0xaf, 0xb6, 0x79, 0x4b, 0x83, 0x23, 0xb7, 0x0c, 0xfe, 0xa3, 0x0f, 0xe5, 0x1a, 0x27,
	0x30, 0x22, 0xdf, 0xe7, 0xc5, 0x0a, 0x25, 0x1c, 0x56, 0xde, 0x91, 0x71, 0xc4, 0x71, 0xb9, 0x4e,
	0x52, 0xdd, 0xc2, 0x64, 0x90, 0xd8, 0xfd, 0x72, 0xc0, 0x8a, 0x3f, 0xec, 0xfd, 0x97, 0x80, 0x6c,
	0x5a, 0x91, 0xf5, 0x0e, 0x1f, 0x61, 0xdc, 0xed, 0x85, 0x17, 0x65, 0xba, 0x47, 0x39, 0x38, 0x46,
	0x71, 0xb6, 0xdb, 0x8e, 0xc1, 0x6a, 0x0f, 0x9f, 0x20, 0xff, 0x19, 0x86, 0xc5, 0x2f, 0xb3, 0xbb,
	0x53, 0x21, 0xff, 0xf5, 0x38, 0xe4, 0x4e, 0x2c, 0x32, 0x36, 0x1f, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x01, 0x49, 0x81, 0x98, 0x01, 0x00, 0x00,
}