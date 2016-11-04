// Code generated by protoc-gen-go.
// source: slinkpb.proto
// DO NOT EDIT!

/*
Package slinkpb is a generated protocol buffer package.

It is generated from these files:
	slinkpb.proto

It has these top-level messages:
	Timestamp
	Sample
	Packet
	Selection
*/
package slinkpb

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

type Timestamp struct {
	Nanoseconds int64 `protobuf:"varint,1,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Sample struct {
	Epoch *Timestamp `protobuf:"bytes,1,opt,name=epoch" json:"epoch,omitempty"`
	Value int32      `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Sample) GetEpoch() *Timestamp {
	if m != nil {
		return m.Epoch
	}
	return nil
}

type Packet struct {
	Network  string     `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Station  string     `protobuf:"bytes,2,opt,name=station" json:"station,omitempty"`
	Location string     `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	Channel  string     `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	Start    *Timestamp `protobuf:"bytes,5,opt,name=start" json:"start,omitempty"`
	End      *Timestamp `protobuf:"bytes,6,opt,name=end" json:"end,omitempty"`
	Length   int32      `protobuf:"varint,7,opt,name=length" json:"length,omitempty"`
	Sps      float32    `protobuf:"fixed32,8,opt,name=sps" json:"sps,omitempty"`
	Samples  []*Sample  `protobuf:"bytes,9,rep,name=samples" json:"samples,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Packet) GetStart() *Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Packet) GetEnd() *Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Packet) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type Selection struct {
	Streams   string `protobuf:"bytes,1,opt,name=streams" json:"streams,omitempty"`
	Selectors string `protobuf:"bytes,2,opt,name=selectors" json:"selectors,omitempty"`
}

func (m *Selection) Reset()                    { *m = Selection{} }
func (m *Selection) String() string            { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()               {}
func (*Selection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Timestamp)(nil), "slinkpb.Timestamp")
	proto.RegisterType((*Sample)(nil), "slinkpb.Sample")
	proto.RegisterType((*Packet)(nil), "slinkpb.Packet")
	proto.RegisterType((*Selection)(nil), "slinkpb.Selection")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeedLink service

type SeedLinkClient interface {
	Stream(ctx context.Context, in *Selection, opts ...grpc.CallOption) (SeedLink_StreamClient, error)
}

type seedLinkClient struct {
	cc *grpc.ClientConn
}

func NewSeedLinkClient(cc *grpc.ClientConn) SeedLinkClient {
	return &seedLinkClient{cc}
}

func (c *seedLinkClient) Stream(ctx context.Context, in *Selection, opts ...grpc.CallOption) (SeedLink_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeedLink_serviceDesc.Streams[0], c.cc, "/slinkpb.SeedLink/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &seedLinkStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeedLink_StreamClient interface {
	Recv() (*Packet, error)
	grpc.ClientStream
}

type seedLinkStreamClient struct {
	grpc.ClientStream
}

func (x *seedLinkStreamClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SeedLink service

type SeedLinkServer interface {
	Stream(*Selection, SeedLink_StreamServer) error
}

func RegisterSeedLinkServer(s *grpc.Server, srv SeedLinkServer) {
	s.RegisterService(&_SeedLink_serviceDesc, srv)
}

func _SeedLink_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Selection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeedLinkServer).Stream(m, &seedLinkStreamServer{stream})
}

type SeedLink_StreamServer interface {
	Send(*Packet) error
	grpc.ServerStream
}

type seedLinkStreamServer struct {
	grpc.ServerStream
}

func (x *seedLinkStreamServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

var _SeedLink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinkpb.SeedLink",
	HandlerType: (*SeedLinkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _SeedLink_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "slinkpb.proto",
}

func init() { proto.RegisterFile("slinkpb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x4f, 0xac, 0x40,
	0x10, 0xc6, 0x1f, 0xc7, 0x03, 0x8e, 0xb9, 0xbc, 0xbc, 0x97, 0xc9, 0x8b, 0xd9, 0x5c, 0x2c, 0x08,
	0xb1, 0xc0, 0xc2, 0x8b, 0x9e, 0xb5, 0x8d, 0x36, 0x16, 0x16, 0x06, 0xac, 0xec, 0xf6, 0xb8, 0x89,
	0x47, 0x80, 0x5d, 0xc2, 0xac, 0xfa, 0xa7, 0xf8, 0xef, 0x1a, 0x16, 0x96, 0xb3, 0xb9, 0x8e, 0xdf,
	0x7e, 0xdf, 0xcc, 0xce, 0x7c, 0x2c, 0xfc, 0xe1, 0xa6, 0x52, 0x75, 0xb7, 0xdb, 0x74, 0xbd, 0x36,
	0x1a, 0xa3, 0x09, 0xd3, 0x2b, 0x88, 0x5f, 0xaa, 0x96, 0xd8, 0xc8, 0xb6, 0xc3, 0x04, 0x56, 0x4a,
	0x2a, 0xcd, 0x54, 0x6a, 0xb5, 0x67, 0xe1, 0x25, 0x5e, 0xe6, 0xe7, 0x3f, 0x8f, 0xd2, 0x47, 0x08,
	0x0b, 0xd9, 0x76, 0x0d, 0x61, 0x06, 0x01, 0x75, 0xba, 0x3c, 0x58, 0xd7, 0x6a, 0x8b, 0x1b, 0x77,
	0xc1, 0xdc, 0x2e, 0x1f, 0x0d, 0xf8, 0x1f, 0x82, 0x0f, 0xd9, 0xbc, 0x93, 0x58, 0x24, 0x5e, 0x16,
	0xe4, 0x23, 0xa4, 0x5f, 0x0b, 0x08, 0x9f, 0x65, 0x59, 0x93, 0x41, 0x01, 0x91, 0x22, 0xf3, 0xa9,
	0xfb, 0xda, 0x36, 0x8b, 0x73, 0x87, 0x83, 0xc2, 0x46, 0x9a, 0x4a, 0x2b, 0x5b, 0x1c, 0xe7, 0x0e,
	0x71, 0x0d, 0xcb, 0x46, 0x97, 0xa3, 0xe4, 0x5b, 0x69, 0xe6, 0xa1, 0xaa, 0x3c, 0x48, 0xa5, 0xa8,
	0x11, 0xbf, 0xc7, 0xaa, 0x09, 0x87, 0xa1, 0xd9, 0xc8, 0xde, 0x88, 0xe0, 0xf4, 0xd0, 0xd6, 0x80,
	0x17, 0xe0, 0x93, 0xda, 0x8b, 0xf0, 0xa4, 0x6f, 0x90, 0xf1, 0x0c, 0xc2, 0x86, 0xd4, 0x9b, 0x39,
	0x88, 0xc8, 0xee, 0x36, 0x11, 0xfe, 0x03, 0x9f, 0x3b, 0x16, 0xcb, 0xc4, 0xcb, 0x16, 0xf9, 0xf0,
	0x89, 0x97, 0x10, 0xb1, 0x0d, 0x8e, 0x45, 0x9c, 0xf8, 0xd9, 0x6a, 0xfb, 0x77, 0xee, 0x39, 0x06,
	0x9a, 0x3b, 0x3d, 0x7d, 0x80, 0xb8, 0xa0, 0x86, 0x4a, 0xb7, 0x0b, 0x9b, 0x9e, 0x64, 0xcb, 0x2e,
	0x9b, 0x09, 0xf1, 0x1c, 0x62, 0xb6, 0x36, 0xdd, 0xf3, 0x94, 0xce, 0xf1, 0x60, 0x7b, 0x07, 0xcb,
	0x82, 0x68, 0xff, 0x54, 0xa9, 0x1a, 0x6f, 0x20, 0x2c, 0x6c, 0x11, 0x1e, 0x17, 0x99, 0x6f, 0x58,
	0x1f, 0x07, 0x19, 0x7f, 0x47, 0xfa, 0xeb, 0xda, 0xbb, 0x8f, 0x5f, 0xdd, 0x0b, 0xd9, 0x85, 0xf6,
	0xc5, 0xdc, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xfd, 0xa4, 0x5c, 0x42, 0x02, 0x00, 0x00,
}
