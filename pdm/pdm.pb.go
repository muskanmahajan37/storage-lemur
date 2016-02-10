// Code generated by protoc-gen-go.
// source: pdm/pdm.proto
// DO NOT EDIT!

/*
Package pdm is a generated protocol buffer package.

It is generated from these files:
	pdm/pdm.proto

It has these top-level messages:
	Endpoint
	Handle
	ActionItem
	ActionStatus
	Empty
*/
package pdm

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

type Command int32

const (
	Command_NONE    Command = 0
	Command_ARCHIVE Command = 1
	Command_RESTORE Command = 2
	Command_REMOVE  Command = 3
	Command_CANCEL  Command = 4
)

var Command_name = map[int32]string{
	0: "NONE",
	1: "ARCHIVE",
	2: "RESTORE",
	3: "REMOVE",
	4: "CANCEL",
}
var Command_value = map[string]int32{
	"NONE":    0,
	"ARCHIVE": 1,
	"RESTORE": 2,
	"REMOVE":  3,
	"CANCEL":  4,
}

func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}
func (Command) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Endpoint struct {
	FsUrl   string `protobuf:"bytes,2,opt,name=fs_url" json:"fs_url,omitempty"`
	Archive uint32 `protobuf:"varint,1,opt,name=archive" json:"archive,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Handle struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Handle) Reset()                    { *m = Handle{} }
func (m *Handle) String() string            { return proto.CompactTextString(m) }
func (*Handle) ProtoMessage()               {}
func (*Handle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ActionItem struct {
	Id          uint64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Op          Command `protobuf:"varint,2,opt,name=op,enum=pdm.Command" json:"op,omitempty"`
	PrimaryPath string  `protobuf:"bytes,3,opt,name=primary_path" json:"primary_path,omitempty"`
	WritePath   string  `protobuf:"bytes,4,opt,name=write_path" json:"write_path,omitempty"`
	Offset      uint64  `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Length      uint64  `protobuf:"varint,6,opt,name=length" json:"length,omitempty"`
	FileId      []byte  `protobuf:"bytes,7,opt,name=file_id,proto3" json:"file_id,omitempty"`
	Data        []byte  `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ActionItem) Reset()                    { *m = ActionItem{} }
func (m *ActionItem) String() string            { return proto.CompactTextString(m) }
func (*ActionItem) ProtoMessage()               {}
func (*ActionItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ActionStatus struct {
	Id        uint64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Completed bool    `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
	Error     int32   `protobuf:"varint,3,opt,name=error" json:"error,omitempty"`
	Offset    uint64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Length    uint64  `protobuf:"varint,5,opt,name=length" json:"length,omitempty"`
	Handle    *Handle `protobuf:"bytes,6,opt,name=handle" json:"handle,omitempty"`
	FileId    []byte  `protobuf:"bytes,7,opt,name=file_id,proto3" json:"file_id,omitempty"`
	Flags     int32   `protobuf:"varint,8,opt,name=flags" json:"flags,omitempty"`
}

func (m *ActionStatus) Reset()                    { *m = ActionStatus{} }
func (m *ActionStatus) String() string            { return proto.CompactTextString(m) }
func (*ActionStatus) ProtoMessage()               {}
func (*ActionStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ActionStatus) GetHandle() *Handle {
	if m != nil {
		return m.Handle
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Endpoint)(nil), "pdm.Endpoint")
	proto.RegisterType((*Handle)(nil), "pdm.Handle")
	proto.RegisterType((*ActionItem)(nil), "pdm.ActionItem")
	proto.RegisterType((*ActionStatus)(nil), "pdm.ActionStatus")
	proto.RegisterType((*Empty)(nil), "pdm.Empty")
	proto.RegisterEnum("pdm.Command", Command_name, Command_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for DataMover service

type DataMoverClient interface {
	Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*Handle, error)
	GetActions(ctx context.Context, in *Handle, opts ...grpc.CallOption) (DataMover_GetActionsClient, error)
	StatusStream(ctx context.Context, opts ...grpc.CallOption) (DataMover_StatusStreamClient, error)
}

type dataMoverClient struct {
	cc *grpc.ClientConn
}

func NewDataMoverClient(cc *grpc.ClientConn) DataMoverClient {
	return &dataMoverClient{cc}
}

func (c *dataMoverClient) Register(ctx context.Context, in *Endpoint, opts ...grpc.CallOption) (*Handle, error) {
	out := new(Handle)
	err := grpc.Invoke(ctx, "/pdm.DataMover/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMoverClient) GetActions(ctx context.Context, in *Handle, opts ...grpc.CallOption) (DataMover_GetActionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMover_serviceDesc.Streams[0], c.cc, "/pdm.DataMover/GetActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMoverGetActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataMover_GetActionsClient interface {
	Recv() (*ActionItem, error)
	grpc.ClientStream
}

type dataMoverGetActionsClient struct {
	grpc.ClientStream
}

func (x *dataMoverGetActionsClient) Recv() (*ActionItem, error) {
	m := new(ActionItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataMoverClient) StatusStream(ctx context.Context, opts ...grpc.CallOption) (DataMover_StatusStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMover_serviceDesc.Streams[1], c.cc, "/pdm.DataMover/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMoverStatusStreamClient{stream}
	return x, nil
}

type DataMover_StatusStreamClient interface {
	Send(*ActionStatus) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type dataMoverStatusStreamClient struct {
	grpc.ClientStream
}

func (x *dataMoverStatusStreamClient) Send(m *ActionStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataMoverStatusStreamClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DataMover service

type DataMoverServer interface {
	Register(context.Context, *Endpoint) (*Handle, error)
	GetActions(*Handle, DataMover_GetActionsServer) error
	StatusStream(DataMover_StatusStreamServer) error
}

func RegisterDataMoverServer(s *grpc.Server, srv DataMoverServer) {
	s.RegisterService(&_DataMover_serviceDesc, srv)
}

func _DataMover_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Endpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DataMoverServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _DataMover_GetActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Handle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataMoverServer).GetActions(m, &dataMoverGetActionsServer{stream})
}

type DataMover_GetActionsServer interface {
	Send(*ActionItem) error
	grpc.ServerStream
}

type dataMoverGetActionsServer struct {
	grpc.ServerStream
}

func (x *dataMoverGetActionsServer) Send(m *ActionItem) error {
	return x.ServerStream.SendMsg(m)
}

func _DataMover_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataMoverServer).StatusStream(&dataMoverStatusStreamServer{stream})
}

type DataMover_StatusStreamServer interface {
	SendAndClose(*Empty) error
	Recv() (*ActionStatus, error)
	grpc.ServerStream
}

type dataMoverStatusStreamServer struct {
	grpc.ServerStream
}

func (x *dataMoverStatusStreamServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataMoverStatusStreamServer) Recv() (*ActionStatus, error) {
	m := new(ActionStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataMover_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pdm.DataMover",
	HandlerType: (*DataMoverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DataMover_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActions",
			Handler:       _DataMover_GetActions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StatusStream",
			Handler:       _DataMover_StatusStream_Handler,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x71, 0x9b, 0x7f, 0x3d, 0x4d, 0xb7, 0xce, 0xda, 0x45, 0x54, 0x6e, 0x50, 0x2e, 0xd0,
	0x04, 0x68, 0x43, 0xe5, 0x09, 0xaa, 0x62, 0xb1, 0x49, 0xac, 0x95, 0x52, 0xc4, 0x6d, 0x65, 0x1a,
	0xa7, 0xb5, 0x14, 0xc7, 0x96, 0xe3, 0x0d, 0xed, 0x35, 0xb8, 0x43, 0xbc, 0x2c, 0x27, 0x0e, 0x13,
	0x2d, 0x5c, 0x44, 0xf2, 0x39, 0x5f, 0x72, 0xbe, 0xdf, 0xf9, 0x1c, 0x98, 0x98, 0x52, 0xdd, 0xe0,
	0x73, 0x6d, 0xac, 0x76, 0x9a, 0x0e, 0xf1, 0x98, 0xbf, 0x85, 0x84, 0x35, 0xa5, 0xd1, 0xb2, 0x71,
	0xf4, 0x0c, 0xa2, 0xaa, 0xdd, 0x3e, 0xd8, 0x3a, 0x1b, 0xbc, 0x22, 0x57, 0x23, 0x7a, 0x0e, 0x31,
	0xb7, 0xbb, 0x83, 0x7c, 0x14, 0x19, 0xc1, 0xc6, 0x24, 0xbf, 0x84, 0xe8, 0x96, 0x37, 0x65, 0x2d,
	0x28, 0xc0, 0x40, 0x96, 0xbe, 0x1b, 0xe4, 0xbf, 0x08, 0xc0, 0x62, 0xe7, 0xa4, 0x6e, 0xee, 0x9c,
	0x50, 0xc7, 0x12, 0xcd, 0x60, 0xa0, 0x8d, 0x9f, 0x76, 0x36, 0x4f, 0xaf, 0x3b, 0xeb, 0xa5, 0x56,
	0x0a, 0x47, 0xd0, 0x4b, 0x48, 0x8d, 0x95, 0x8a, 0xdb, 0xa7, 0xad, 0xe1, 0xee, 0x90, 0x0d, 0xbd,
	0x23, 0x7e, 0xfc, 0xdd, 0x4a, 0x27, 0xfa, 0x5e, 0xe0, 0x7b, 0x48, 0xa5, 0xab, 0xaa, 0x15, 0x2e,
	0x0b, 0xfd, 0x4c, 0xac, 0x6b, 0xd1, 0xec, 0x51, 0x8f, 0x7c, 0x8d, 0x94, 0x95, 0xac, 0xc5, 0x16,
	0x4d, 0x63, 0x6c, 0xa4, 0x34, 0x85, 0xa0, 0xe4, 0x8e, 0x67, 0x49, 0x57, 0xe5, 0x3f, 0x09, 0xa4,
	0x3d, 0xdd, 0xc6, 0x71, 0xf7, 0xd0, 0x9e, 0xf0, 0x5d, 0xc0, 0x68, 0xa7, 0x95, 0xa9, 0x85, 0x13,
	0xa5, 0xc7, 0x4c, 0xe8, 0x04, 0x42, 0x61, 0xad, 0xb6, 0x9e, 0x28, 0x3c, 0x72, 0x0f, 0xfe, 0x71,
	0xef, 0x69, 0x5e, 0x42, 0x74, 0xf0, 0x91, 0x78, 0x9a, 0xf1, 0x7c, 0xec, 0xb7, 0xfc, 0x93, 0xd2,
	0x7f, 0x68, 0x38, 0xbc, 0xaa, 0xf9, 0xbe, 0xf5, 0x6c, 0x61, 0x1e, 0x43, 0xc8, 0x94, 0x71, 0x4f,
	0x6f, 0x18, 0xc4, 0xcf, 0xc1, 0x24, 0x10, 0xac, 0xd6, 0x2b, 0x36, 0x7d, 0x41, 0xc7, 0x10, 0x2f,
	0x8a, 0xe5, 0xed, 0xdd, 0x57, 0x36, 0x25, 0x5d, 0x51, 0xb0, 0xcd, 0x97, 0x75, 0xc1, 0xa6, 0x03,
	0x5c, 0x21, 0x2a, 0xd8, 0xfd, 0x1a, 0x85, 0x61, 0x77, 0x5e, 0x2e, 0x56, 0x4b, 0xf6, 0x79, 0x1a,
	0xcc, 0x7f, 0x10, 0x18, 0x7d, 0xc4, 0xd5, 0xef, 0xf5, 0xa3, 0xb0, 0xf4, 0x35, 0x24, 0x85, 0xd8,
	0xcb, 0xd6, 0xe1, 0x79, 0xe2, 0xb1, 0x9e, 0x6f, 0x7a, 0x76, 0x42, 0xf9, 0x0e, 0xe0, 0x93, 0x70,
	0x7d, 0x46, 0x2d, 0x3d, 0x96, 0x66, 0xe7, 0xbe, 0xf8, 0x7b, 0xb9, 0xef, 0x09, 0xbd, 0x81, 0xb4,
	0x0f, 0x72, 0xe3, 0xac, 0xe0, 0x8a, 0x5e, 0x1c, 0xbd, 0xd2, 0x0b, 0x33, 0xe8, 0xcd, 0xba, 0xcd,
	0xae, 0xc8, 0xb7, 0xc8, 0xff, 0x6d, 0x1f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x16, 0xf8, 0x53,
	0x64, 0x7e, 0x02, 0x00, 0x00,
}