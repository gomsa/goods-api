// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/unspsc/unspsc.proto

package unspsc

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

// 数据库 Unspscs
type Unspsc struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               int64    `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unspsc) Reset()         { *m = Unspsc{} }
func (m *Unspsc) String() string { return proto.CompactTextString(m) }
func (*Unspsc) ProtoMessage()    {}
func (*Unspsc) Descriptor() ([]byte, []int) {
	return fileDescriptor_169f1574304aeaca, []int{0}
}

func (m *Unspsc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unspsc.Unmarshal(m, b)
}
func (m *Unspsc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unspsc.Marshal(b, m, deterministic)
}
func (m *Unspsc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unspsc.Merge(m, src)
}
func (m *Unspsc) XXX_Size() int {
	return xxx_messageInfo_Unspsc.Size(m)
}
func (m *Unspsc) XXX_DiscardUnknown() {
	xxx_messageInfo_Unspsc.DiscardUnknown(m)
}

var xxx_messageInfo_Unspsc proto.InternalMessageInfo

func (m *Unspsc) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Unspsc) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *Unspsc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Unspsc) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Unspsc) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	Unspsc               *Unspsc  `protobuf:"bytes,2,opt,name=Unspsc,proto3" json:"Unspsc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_169f1574304aeaca, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUnspsc() *Unspsc {
	if m != nil {
		return m.Unspsc
	}
	return nil
}

type Response struct {
	Unspsc               *Unspsc   `protobuf:"bytes,1,opt,name=unspsc,proto3" json:"unspsc,omitempty"`
	Unspscs              []*Unspsc `protobuf:"bytes,2,rep,name=unspscs,proto3" json:"unspscs,omitempty"`
	Valid                bool      `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_169f1574304aeaca, []int{2}
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

func (m *Response) GetUnspsc() *Unspsc {
	if m != nil {
		return m.Unspsc
	}
	return nil
}

func (m *Response) GetUnspscs() []*Unspsc {
	if m != nil {
		return m.Unspscs
	}
	return nil
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Unspsc)(nil), "unspsc.Unspsc")
	proto.RegisterType((*Request)(nil), "unspsc.Request")
	proto.RegisterType((*Response)(nil), "unspsc.Response")
}

func init() { proto.RegisterFile("proto/unspsc/unspsc.proto", fileDescriptor_169f1574304aeaca) }

var fileDescriptor_169f1574304aeaca = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x4b, 0x84, 0x40,
	0x14, 0xc5, 0xf3, 0xdf, 0xb8, 0x7b, 0x17, 0xb6, 0x18, 0x22, 0x2c, 0x08, 0xc4, 0x87, 0x90, 0x6a,
	0x37, 0xb2, 0x4f, 0x20, 0x5b, 0xf4, 0xd2, 0xd3, 0x80, 0xcf, 0x61, 0x7a, 0x21, 0xc9, 0xd4, 0x9c,
	0x31, 0x7a, 0xee, 0x7b, 0x07, 0xb1, 0x77, 0x46, 0xf6, 0x65, 0x1f, 0xdc, 0x27, 0xef, 0x3d, 0xe7,
	0x37, 0xe3, 0x39, 0x28, 0x9c, 0x77, 0x7d, 0xab, 0xda, 0xbb, 0xa1, 0x91, 0x9d, 0x2c, 0xcc, 0x63,
	0x4d, 0x1a, 0x67, 0x7a, 0x8b, 0x7e, 0x2d, 0x60, 0x19, 0x8d, 0x7c, 0x09, 0x76, 0x55, 0x06, 0x56,
	0x68, 0xc5, 0x8e, 0xb0, 0xab, 0x92, 0x9f, 0x01, 0xeb, 0xf2, 0x1e, 0x1b, 0x15, 0xd8, 0xa4, 0x99,
	0x8d, 0x73, 0x70, 0x9b, 0xfc, 0x13, 0x03, 0x27, 0xb4, 0xe2, 0xb9, 0xa0, 0x99, 0x5f, 0x02, 0x14,
	0x3d, 0xe6, 0x0a, 0xcb, 0xd7, 0x5c, 0x05, 0x2e, 0x39, 0x73, 0xa3, 0xa4, 0x6a, 0x6b, 0x0f, 0x5d,
	0x39, 0xda, 0x9e, 0xb6, 0x8d, 0x92, 0xaa, 0xe8, 0x1e, 0x7c, 0x81, 0x5f, 0x03, 0x4a, 0xc5, 0xaf,
	0xc6, 0x38, 0xf4, 0xd2, 0x45, 0xb2, 0x5c, 0x9b, 0xd8, 0x5a, 0x15, 0xc6, 0x8d, 0x7a, 0x98, 0x09,
	0x94, 0x5d, 0xdb, 0x48, 0xdc, 0x9e, 0xd1, 0x10, 0x85, 0xdf, 0x73, 0x46, 0xaf, 0x3c, 0x06, 0x5f,
	0x4f, 0x32, 0xb0, 0x43, 0x67, 0x0f, 0x38, 0xda, 0xfc, 0x14, 0xbc, 0xef, 0xbc, 0xae, 0x4a, 0xea,
	0x38, 0x13, 0x7a, 0x49, 0xfe, 0x6c, 0xf0, 0x33, 0x43, 0x24, 0xb0, 0xd8, 0xbc, 0x63, 0xf1, 0xb1,
	0xa1, 0x8e, 0xfc, 0x78, 0xbc, 0xc9, 0xf4, 0xb8, 0x38, 0xd9, 0x09, 0x3a, 0x65, 0x74, 0xc4, 0x6f,
	0xc1, 0x7b, 0xfa, 0xa9, 0xa4, 0x9a, 0x46, 0x5f, 0x83, 0x93, 0xd6, 0xf5, 0x34, 0xf6, 0x06, 0xdc,
	0x97, 0x43, 0x2e, 0x7e, 0xc6, 0x89, 0xec, 0x0a, 0xd8, 0x21, 0x0d, 0x57, 0xc0, 0x32, 0xfa, 0xaa,
	0x93, 0xf1, 0x47, 0xac, 0x71, 0x22, 0xfe, 0xc6, 0xe8, 0xd7, 0x7d, 0xf8, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xf2, 0xc8, 0xff, 0xb6, 0xd7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Unspscs service

type UnspscsClient interface {
	// 检查批量创建
	CheckCreate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 检查国际商品及服务编码
	Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 全部国际商品及服务编码
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取国际商品及服务编码编码列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取国际商品及服务编码编码信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建国际商品及服务编码编码
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新国际商品及服务编码编码
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除国际商品及服务编码编码
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type unspscsClient struct {
	c           client.Client
	serviceName string
}

func NewUnspscsClient(serviceName string, c client.Client) UnspscsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "unspsc"
	}
	return &unspscsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *unspscsClient) CheckCreate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.CheckCreate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.Exist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unspscsClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Unspscs.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Unspscs service

type UnspscsHandler interface {
	// 检查批量创建
	CheckCreate(context.Context, *Request, *Response) error
	// 检查国际商品及服务编码
	Exist(context.Context, *Request, *Response) error
	// 全部国际商品及服务编码
	All(context.Context, *Request, *Response) error
	// 获取国际商品及服务编码编码列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取国际商品及服务编码编码信息
	Get(context.Context, *Request, *Response) error
	// 创建国际商品及服务编码编码
	Create(context.Context, *Request, *Response) error
	// 更新国际商品及服务编码编码
	Update(context.Context, *Request, *Response) error
	// 删除国际商品及服务编码编码
	Delete(context.Context, *Request, *Response) error
}

func RegisterUnspscsHandler(s server.Server, hdlr UnspscsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Unspscs{hdlr}, opts...))
}

type Unspscs struct {
	UnspscsHandler
}

func (h *Unspscs) CheckCreate(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.CheckCreate(ctx, in, out)
}

func (h *Unspscs) Exist(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.Exist(ctx, in, out)
}

func (h *Unspscs) All(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.All(ctx, in, out)
}

func (h *Unspscs) List(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.List(ctx, in, out)
}

func (h *Unspscs) Get(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.Get(ctx, in, out)
}

func (h *Unspscs) Create(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.Create(ctx, in, out)
}

func (h *Unspscs) Update(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.Update(ctx, in, out)
}

func (h *Unspscs) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.UnspscsHandler.Delete(ctx, in, out)
}
