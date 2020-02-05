// Code generated by protoc-gen-go. DO NOT EDIT.
// source: slashing.proto

package filecoin_slashing_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Index struct {
	TipSetKey            string              `protobuf:"bytes,1,opt,name=tipSetKey,proto3" json:"tipSetKey,omitempty"`
	Miners               map[string]*Slashes `protobuf:"bytes,2,rep,name=miners,proto3" json:"miners,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetTipSetKey() string {
	if m != nil {
		return m.TipSetKey
	}
	return ""
}

func (m *Index) GetMiners() map[string]*Slashes {
	if m != nil {
		return m.Miners
	}
	return nil
}

type Slashes struct {
	Epochs               []uint64 `protobuf:"varint,1,rep,packed,name=epochs,proto3" json:"epochs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slashes) Reset()         { *m = Slashes{} }
func (m *Slashes) String() string { return proto.CompactTextString(m) }
func (*Slashes) ProtoMessage()    {}
func (*Slashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{1}
}

func (m *Slashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slashes.Unmarshal(m, b)
}
func (m *Slashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slashes.Marshal(b, m, deterministic)
}
func (m *Slashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slashes.Merge(m, src)
}
func (m *Slashes) XXX_Size() int {
	return xxx_messageInfo_Slashes.Size(m)
}
func (m *Slashes) XXX_DiscardUnknown() {
	xxx_messageInfo_Slashes.DiscardUnknown(m)
}

var xxx_messageInfo_Slashes proto.InternalMessageInfo

func (m *Slashes) GetEpochs() []uint64 {
	if m != nil {
		return m.Epochs
	}
	return nil
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetReply struct {
	Index                *Index   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{3}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetIndex() *Index {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterType((*Index)(nil), "filecoin.slashing.pb.Index")
	proto.RegisterMapType((map[string]*Slashes)(nil), "filecoin.slashing.pb.Index.MinersEntry")
	proto.RegisterType((*Slashes)(nil), "filecoin.slashing.pb.Slashes")
	proto.RegisterType((*GetRequest)(nil), "filecoin.slashing.pb.GetRequest")
	proto.RegisterType((*GetReply)(nil), "filecoin.slashing.pb.GetReply")
}

func init() { proto.RegisterFile("slashing.proto", fileDescriptor_31f622956ca78100) }

var fileDescriptor_31f622956ca78100 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xdd, 0xac, 0xad, 0x76, 0x22, 0x52, 0x16, 0x91, 0x50, 0xff, 0xc5, 0x5c, 0xcc, 0x69,
	0xc1, 0xf4, 0x22, 0x8a, 0x88, 0x85, 0x5a, 0x8a, 0x08, 0x61, 0xd3, 0x83, 0xd7, 0xb6, 0x8e, 0x76,
	0x71, 0x4d, 0x62, 0x76, 0x2b, 0xcd, 0xeb, 0xf8, 0x24, 0x3e, 0x9a, 0x24, 0x0d, 0xc6, 0x43, 0xe8,
	0x6d, 0xe6, 0xcb, 0x37, 0xf3, 0x9b, 0xb0, 0xb0, 0xaf, 0xd5, 0x54, 0x2f, 0x64, 0xfc, 0xc6, 0xd3,
	0x2c, 0x31, 0x09, 0x3b, 0x78, 0x95, 0x0a, 0xe7, 0x89, 0x8c, 0x79, 0xfd, 0x61, 0xe6, 0xfd, 0x10,
	0x68, 0x8d, 0xe3, 0x17, 0x5c, 0xb1, 0x63, 0xe8, 0x18, 0x99, 0x46, 0x68, 0x1e, 0x31, 0x77, 0x88,
	0x4b, 0xfc, 0x8e, 0xa8, 0x01, 0xbb, 0x83, 0xf6, 0x87, 0x8c, 0x31, 0xd3, 0x8e, 0xe5, 0x52, 0xdf,
	0x0e, 0x2e, 0x78, 0xd3, 0x3a, 0x5e, 0xae, 0xe2, 0x4f, 0xa5, 0x39, 0x8c, 0x4d, 0x96, 0x8b, 0x6a,
	0xac, 0xf7, 0x0c, 0xf6, 0x3f, 0xcc, 0xba, 0x40, 0xdf, 0xff, 0x72, 0x8a, 0x92, 0xf5, 0xa1, 0xf5,
	0x35, 0x55, 0x4b, 0x74, 0x2c, 0x97, 0xf8, 0x76, 0x70, 0xd2, 0x1c, 0x10, 0x15, 0x35, 0x6a, 0xb1,
	0x76, 0xaf, 0xad, 0x2b, 0xe2, 0x9d, 0xc3, 0x4e, 0x45, 0xd9, 0x21, 0xb4, 0x31, 0x4d, 0xe6, 0x0b,
	0xed, 0x10, 0x97, 0xfa, 0xdb, 0xa2, 0xea, 0xbc, 0x3d, 0x80, 0x11, 0x1a, 0x81, 0x9f, 0x4b, 0xd4,
	0xc6, 0xbb, 0x85, 0xdd, 0xb2, 0x4b, 0x55, 0xce, 0x2e, 0xa1, 0x25, 0x8b, 0x9b, 0xcb, 0x4b, 0xec,
	0xe0, 0x68, 0xc3, 0x6f, 0x89, 0xb5, 0x19, 0x84, 0x40, 0xef, 0xc3, 0x31, 0x1b, 0x03, 0x1d, 0xa1,
	0x61, 0x6e, 0xf3, 0x44, 0x1d, 0xd7, 0x3b, 0xdd, 0x60, 0xa4, 0x2a, 0xf7, 0xb6, 0x06, 0x37, 0x70,
	0x26, 0x13, 0x6e, 0x70, 0x65, 0xa4, 0xc2, 0x46, 0x7b, 0xd0, 0x7d, 0xa8, 0x68, 0x54, 0xc1, 0x90,
	0x7c, 0x5b, 0x74, 0x32, 0x19, 0xce, 0xda, 0xe5, 0xf3, 0xf6, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0x40, 0xaf, 0xff, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/filecoin.slashing.pb.API/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filecoin.slashing.pb.API/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filecoin.slashing.pb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _API_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slashing.proto",
}