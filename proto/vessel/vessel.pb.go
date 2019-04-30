// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel/vessel.proto

package shippy_service_vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwerId               string   `protobuf:"bytes,6,opt,name=ower_id,json=owerId,proto3" json:"ower_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd0828b40ed5ef5, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwerId() string {
	if m != nil {
		return m.OwerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd0828b40ed5ef5, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessels              []*Vessel `protobuf:"bytes,1,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd0828b40ed5ef5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "shippy.service.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "shippy.service.vessel.Specification")
	proto.RegisterType((*Response)(nil), "shippy.service.vessel.Response")
}

func init() { proto.RegisterFile("vessel/vessel.proto", fileDescriptor_2bd0828b40ed5ef5) }

var fileDescriptor_2bd0828b40ed5ef5 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0x36, 0x6d, 0xd3, 0x64, 0x3e, 0xe2, 0x61, 0x45, 0x5c, 0x8a, 0xc5, 0x10, 0x3c, 0xe4,
	0x14, 0xa1, 0x1e, 0x3c, 0x8b, 0x20, 0xe8, 0x31, 0x05, 0xf5, 0x56, 0xb6, 0xc9, 0x68, 0x07, 0x92,
	0xec, 0x92, 0x0d, 0x69, 0xfb, 0x67, 0xfc, 0xad, 0xe2, 0x2e, 0xad, 0x54, 0xac, 0xa7, 0x9d, 0x79,
	0xf3, 0x76, 0x78, 0xef, 0x0d, 0x9c, 0xf6, 0x68, 0x0c, 0x56, 0xd7, 0xee, 0xc9, 0x74, 0xab, 0x3a,
	0xc5, 0xcf, 0xcc, 0x8a, 0xb4, 0xde, 0x66, 0x06, 0xdb, 0x9e, 0x0a, 0xcc, 0xdc, 0x30, 0xf9, 0x60,
	0xe0, 0x3f, 0xdb, 0x92, 0x9f, 0x80, 0x47, 0xa5, 0x60, 0x31, 0x4b, 0xc3, 0xdc, 0xa3, 0x92, 0x4f,
	0x20, 0x28, 0xa4, 0x96, 0x05, 0x75, 0x5b, 0xe1, 0xc5, 0x2c, 0x1d, 0xe5, 0xfb, 0x9e, 0x4f, 0x01,
	0x6a, 0xb9, 0x59, 0xac, 0x91, 0xde, 0x57, 0x9d, 0x18, 0xd8, 0x69, 0x58, 0xcb, 0xcd, 0x8b, 0x05,
	0x38, 0x87, 0x61, 0x23, 0x6b, 0x14, 0x43, 0xbb, 0xcc, 0xd6, 0xfc, 0x02, 0x42, 0xd9, 0x4b, 0xaa,
	0xe4, 0xb2, 0x42, 0x31, 0x8a, 0x59, 0x1a, 0xe4, 0xdf, 0x00, 0x3f, 0x87, 0xb1, 0x5a, 0x63, 0xbb,
	0xa0, 0x52, 0xf8, 0xf6, 0x93, 0xff, 0xd5, 0x3e, 0x96, 0xc9, 0x13, 0x44, 0x73, 0x8d, 0x05, 0xbd,
	0x51, 0x21, 0x3b, 0x52, 0xcd, 0x81, 0x2c, 0xf6, 0xa7, 0x2c, 0xef, 0x87, 0xac, 0xe4, 0x1e, 0x82,
	0x1c, 0x8d, 0x56, 0x8d, 0x41, 0x7e, 0x0b, 0x63, 0x17, 0x81, 0x11, 0x2c, 0x1e, 0xa4, 0xff, 0x67,
	0xd3, 0xec, 0xd7, 0x84, 0x32, 0x97, 0x4e, 0xbe, 0x63, 0xcf, 0x08, 0x22, 0x07, 0xcd, 0x1d, 0x8f,
	0xbf, 0x42, 0xf4, 0x40, 0x4d, 0x79, 0xb7, 0xf7, 0x72, 0x75, 0x64, 0xd3, 0x81, 0x8f, 0xc9, 0xe5,
	0x11, 0xd6, 0x4e, 0x61, 0xf2, 0x6f, 0xe9, 0xdb, 0xd3, 0xdd, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x8a, 0xc6, 0x67, 0xd1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "shippy.service.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}