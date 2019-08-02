// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/taxcode/taxcode.proto

package taxcode

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

// 数据库 Taxcodes
type Taxcode struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               int64    `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Cess                 int64    `protobuf:"varint,3,opt,name=cess,proto3" json:"cess,omitempty"`
	Code                 int64    `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Free                 string   `protobuf:"bytes,6,opt,name=free,proto3" json:"free,omitempty"`
	DutyFree             string   `protobuf:"bytes,7,opt,name=duty_free,json=dutyFree,proto3" json:"duty_free,omitempty"`
	DutyFreeDesc         string   `protobuf:"bytes,8,opt,name=duty_free_desc,json=dutyFreeDesc,proto3" json:"duty_free_desc,omitempty"`
	Instruction          string   `protobuf:"bytes,9,opt,name=instruction,proto3" json:"instruction,omitempty"`
	Keyword              string   `protobuf:"bytes,10,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Version              string   `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Taxcode) Reset()         { *m = Taxcode{} }
func (m *Taxcode) String() string { return proto.CompactTextString(m) }
func (*Taxcode) ProtoMessage()    {}
func (*Taxcode) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78da315917493c0, []int{0}
}

func (m *Taxcode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Taxcode.Unmarshal(m, b)
}
func (m *Taxcode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Taxcode.Marshal(b, m, deterministic)
}
func (m *Taxcode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Taxcode.Merge(m, src)
}
func (m *Taxcode) XXX_Size() int {
	return xxx_messageInfo_Taxcode.Size(m)
}
func (m *Taxcode) XXX_DiscardUnknown() {
	xxx_messageInfo_Taxcode.DiscardUnknown(m)
}

var xxx_messageInfo_Taxcode proto.InternalMessageInfo

func (m *Taxcode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Taxcode) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *Taxcode) GetCess() int64 {
	if m != nil {
		return m.Cess
	}
	return 0
}

func (m *Taxcode) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Taxcode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Taxcode) GetFree() string {
	if m != nil {
		return m.Free
	}
	return ""
}

func (m *Taxcode) GetDutyFree() string {
	if m != nil {
		return m.DutyFree
	}
	return ""
}

func (m *Taxcode) GetDutyFreeDesc() string {
	if m != nil {
		return m.DutyFreeDesc
	}
	return ""
}

func (m *Taxcode) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

func (m *Taxcode) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *Taxcode) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Taxcode) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Taxcode) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78da315917493c0, []int{1}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

type Request struct {
	ListQuery            *ListQuery `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Taxcode              *Taxcode   `protobuf:"bytes,2,opt,name=Taxcode,proto3" json:"Taxcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78da315917493c0, []int{2}
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

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetTaxcode() *Taxcode {
	if m != nil {
		return m.Taxcode
	}
	return nil
}

type Response struct {
	Taxcode              *Taxcode   `protobuf:"bytes,1,opt,name=taxcode,proto3" json:"taxcode,omitempty"`
	Taxcodes             []*Taxcode `protobuf:"bytes,2,rep,name=taxcodes,proto3" json:"taxcodes,omitempty"`
	Valid                bool       `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d78da315917493c0, []int{3}
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

func (m *Response) GetTaxcode() *Taxcode {
	if m != nil {
		return m.Taxcode
	}
	return nil
}

func (m *Response) GetTaxcodes() []*Taxcode {
	if m != nil {
		return m.Taxcodes
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
	proto.RegisterType((*Taxcode)(nil), "taxcode.Taxcode")
	proto.RegisterType((*ListQuery)(nil), "taxcode.ListQuery")
	proto.RegisterType((*Request)(nil), "taxcode.Request")
	proto.RegisterType((*Response)(nil), "taxcode.Response")
}

func init() { proto.RegisterFile("proto/taxcode/taxcode.proto", fileDescriptor_d78da315917493c0) }

var fileDescriptor_d78da315917493c0 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x89, 0x9d, 0xc6, 0xf6, 0xa4, 0x54, 0x30, 0x42, 0x68, 0x45, 0x85, 0x14, 0x59, 0x1c,
	0x2a, 0x14, 0x5a, 0x11, 0x9e, 0x20, 0xa2, 0x02, 0x21, 0x71, 0xc1, 0x2a, 0x67, 0xcb, 0x78, 0x07,
	0x58, 0xe1, 0x7a, 0xdd, 0xdd, 0x75, 0x21, 0x3c, 0x11, 0x8f, 0xc4, 0xe3, 0xa0, 0x1d, 0xaf, 0x4d,
	0x0f, 0x3d, 0x24, 0x27, 0xcf, 0x7c, 0xff, 0x3f, 0x33, 0xde, 0x59, 0x1b, 0x4e, 0x3b, 0xa3, 0x9d,
	0xbe, 0x70, 0xd5, 0xaf, 0x5a, 0x4b, 0x1a, 0x9f, 0xe7, 0x4c, 0x31, 0x09, 0x69, 0xfe, 0x37, 0x82,
	0xe4, 0x6a, 0x88, 0xf1, 0x04, 0x22, 0x25, 0xc5, 0x6c, 0x35, 0x3b, 0x8b, 0x8b, 0x48, 0x49, 0x7c,
	0x0a, 0x8b, 0xae, 0x32, 0xd4, 0x3a, 0x11, 0x31, 0x0b, 0x19, 0x22, 0xcc, 0x6b, 0xb2, 0x56, 0xc4,
	0x4c, 0x39, 0x66, 0xa6, 0x25, 0x89, 0x79, 0x60, 0xbe, 0x1f, 0xc2, 0xbc, 0xad, 0xae, 0x49, 0x1c,
	0xad, 0x66, 0x67, 0x59, 0xc1, 0xb1, 0x67, 0x5f, 0x0d, 0x91, 0x58, 0x0c, 0xcc, 0xc7, 0x78, 0x0a,
	0x99, 0xec, 0xdd, 0xae, 0x64, 0x21, 0x61, 0x21, 0xf5, 0xe0, 0x9d, 0x17, 0x5f, 0xc0, 0xc9, 0x24,
	0x96, 0x92, 0x6c, 0x2d, 0x52, 0x76, 0x1c, 0x8f, 0x8e, 0x4b, 0xb2, 0x35, 0xae, 0x60, 0xa9, 0x5a,
	0xeb, 0x4c, 0x5f, 0x3b, 0xa5, 0x5b, 0x91, 0xb1, 0xe5, 0x2e, 0x42, 0x01, 0xc9, 0x0f, 0xda, 0xfd,
	0xd4, 0x46, 0x0a, 0x60, 0x75, 0x4c, 0xbd, 0x72, 0x4b, 0xc6, 0xfa, 0xba, 0xe5, 0xa0, 0x84, 0x14,
	0x9f, 0x03, 0xd4, 0x86, 0x2a, 0x47, 0xb2, 0xac, 0x9c, 0x38, 0x66, 0x31, 0x0b, 0x64, 0xeb, 0xbc,
	0xdc, 0x77, 0x72, 0x94, 0x1f, 0x0e, 0x72, 0x20, 0x5b, 0x97, 0x7f, 0x80, 0xec, 0xa3, 0xb2, 0xee,
	0x53, 0x4f, 0x66, 0x87, 0x4f, 0xe0, 0xa8, 0x51, 0xd7, 0xca, 0x85, 0xf5, 0x0e, 0x89, 0xdf, 0x46,
	0x57, 0x7d, 0xa3, 0xb0, 0x5f, 0x8e, 0x3d, 0xb3, 0xda, 0x38, 0xde, 0x6e, 0x56, 0x70, 0x9c, 0x7f,
	0x87, 0xa4, 0xa0, 0x9b, 0x9e, 0xac, 0xc3, 0xd7, 0x00, 0x8d, 0xb2, 0xae, 0xbc, 0xf1, 0x6d, 0xb9,
	0xdb, 0x72, 0x83, 0xe7, 0xe3, 0xed, 0x4e, 0x03, 0x8b, 0xac, 0x99, 0x66, 0xbf, 0x9c, 0xae, 0x98,
	0x07, 0x2d, 0x37, 0x8f, 0x26, 0x7f, 0xe0, 0xc5, 0x68, 0xc8, 0x7f, 0x43, 0x5a, 0x90, 0xed, 0x74,
	0x6b, 0xc9, 0xd7, 0x05, 0x5f, 0x98, 0x73, 0x4f, 0x5d, 0x00, 0xb8, 0x86, 0x34, 0x84, 0x56, 0x44,
	0xab, 0xf8, 0x5e, 0xf3, 0xe4, 0xf0, 0xdb, 0xb8, 0xad, 0x1a, 0x25, 0xf9, 0x90, 0x69, 0x31, 0x24,
	0x9b, 0x3f, 0x11, 0xa4, 0x57, 0xa3, 0x65, 0x0d, 0xf1, 0xb6, 0x69, 0xf0, 0x7f, 0x97, 0xb0, 0x80,
	0x67, 0x8f, 0xef, 0x90, 0xe1, 0x45, 0xf3, 0x07, 0xf8, 0x0a, 0xe6, 0xfe, 0xe8, 0xfb, 0xda, 0xd7,
	0x10, 0xbf, 0xa7, 0xbd, 0xdd, 0x17, 0xb0, 0x78, 0xcb, 0x97, 0x7e, 0x40, 0xc1, 0x67, 0xfe, 0x0c,
	0x0e, 0x28, 0xb8, 0xa4, 0x86, 0xf6, 0x2e, 0xf8, 0xb2, 0xe0, 0xdf, 0xf8, 0xcd, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x3c, 0xb3, 0xb8, 0xe5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Taxcodes service

type TaxcodesClient interface {
	// 全部国家税收分类编码
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取国家税收分类编码列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取国家税收分类编码信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建国家税收分类编码
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新国家税收分类编码
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除国家税收分类编码
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type taxcodesClient struct {
	c           client.Client
	serviceName string
}

func NewTaxcodesClient(serviceName string, c client.Client) TaxcodesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "taxcode"
	}
	return &taxcodesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *taxcodesClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxcodesClient) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxcodesClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxcodesClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxcodesClient) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxcodesClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Taxcodes.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Taxcodes service

type TaxcodesHandler interface {
	// 全部国家税收分类编码
	All(context.Context, *Request, *Response) error
	// 获取国家税收分类编码列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取国家税收分类编码信息
	Get(context.Context, *Request, *Response) error
	// 创建国家税收分类编码
	Create(context.Context, *Request, *Response) error
	// 更新国家税收分类编码
	Update(context.Context, *Request, *Response) error
	// 删除国家税收分类编码
	Delete(context.Context, *Request, *Response) error
}

func RegisterTaxcodesHandler(s server.Server, hdlr TaxcodesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Taxcodes{hdlr}, opts...))
}

type Taxcodes struct {
	TaxcodesHandler
}

func (h *Taxcodes) All(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.All(ctx, in, out)
}

func (h *Taxcodes) List(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.List(ctx, in, out)
}

func (h *Taxcodes) Get(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.Get(ctx, in, out)
}

func (h *Taxcodes) Create(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.Create(ctx, in, out)
}

func (h *Taxcodes) Update(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.Update(ctx, in, out)
}

func (h *Taxcodes) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.TaxcodesHandler.Delete(ctx, in, out)
}