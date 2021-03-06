// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/department/department.proto

package department

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

// 数据库 Departments
type Department struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               int64    `protobuf:"varint,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Manages              string   `protobuf:"bytes,5,opt,name=manages,proto3" json:"manages,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{0}
}

func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (m *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(m, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Department) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *Department) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Department) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Department) GetManages() string {
	if m != nil {
		return m.Manages
	}
	return ""
}

func (m *Department) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Department) GetUpdatedAt() string {
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
	return fileDescriptor_a231da9bf7b7101c, []int{1}
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
	ListQuery            *ListQuery  `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Department           *Department `protobuf:"bytes,2,opt,name=Department,proto3" json:"Department,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{2}
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

func (m *Request) GetDepartment() *Department {
	if m != nil {
		return m.Department
	}
	return nil
}

type Response struct {
	Department           *Department   `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	Departments          []*Department `protobuf:"bytes,2,rep,name=departments,proto3" json:"departments,omitempty"`
	Valid                bool          `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a231da9bf7b7101c, []int{3}
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

func (m *Response) GetDepartment() *Department {
	if m != nil {
		return m.Department
	}
	return nil
}

func (m *Response) GetDepartments() []*Department {
	if m != nil {
		return m.Departments
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
	proto.RegisterType((*Department)(nil), "department.Department")
	proto.RegisterType((*ListQuery)(nil), "department.ListQuery")
	proto.RegisterType((*Request)(nil), "department.Request")
	proto.RegisterType((*Response)(nil), "department.Response")
}

func init() { proto.RegisterFile("proto/department/department.proto", fileDescriptor_a231da9bf7b7101c) }

var fileDescriptor_a231da9bf7b7101c = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x7d, 0x49, 0xda, 0xb4, 0xb9, 0x81, 0xb7, 0x98, 0xd7, 0x57, 0x86, 0x07, 0x0f, 0x6a, 0x56,
	0x5d, 0x55, 0x48, 0x55, 0xdc, 0x16, 0x0b, 0x22, 0xb8, 0x31, 0xe0, 0xba, 0x8c, 0xe6, 0x52, 0x03,
	0x93, 0x64, 0x9a, 0x99, 0x2a, 0xfe, 0x86, 0x1f, 0xe3, 0x6f, 0xf8, 0x4b, 0x32, 0x93, 0x69, 0x33,
	0x2e, 0x14, 0xb2, 0xbb, 0xe7, 0x9c, 0x7b, 0xee, 0xbd, 0x39, 0x4c, 0xe0, 0x44, 0x34, 0xb5, 0xaa,
	0x4f, 0x73, 0x14, 0xac, 0x51, 0x25, 0x56, 0xca, 0x29, 0x17, 0x46, 0x23, 0xd0, 0x31, 0xc9, 0xbb,
	0x07, 0xb0, 0x3e, 0x42, 0xf2, 0x1b, 0xfc, 0x22, 0xa7, 0xde, 0xcc, 0x9b, 0x07, 0x99, 0x5f, 0xe4,
	0x64, 0x0a, 0xa1, 0x60, 0x0d, 0x56, 0x8a, 0xfa, 0x86, 0xb3, 0x88, 0x10, 0x18, 0x54, 0xac, 0x44,
	0x1a, 0xcc, 0xbc, 0x79, 0x94, 0x99, 0x9a, 0x4c, 0x60, 0x28, 0x9e, 0xea, 0x0a, 0xe9, 0xc0, 0x90,
	0x2d, 0x20, 0x14, 0x46, 0x25, 0xab, 0xd8, 0x16, 0x25, 0x1d, 0x1a, 0xfe, 0x00, 0xc9, 0x7f, 0x80,
	0xc7, 0x06, 0x99, 0xc2, 0x7c, 0xc3, 0x14, 0x0d, 0x8d, 0x18, 0x59, 0x66, 0xa5, 0xb4, 0xbc, 0x17,
	0xf9, 0x41, 0x1e, 0xb5, 0xb2, 0x65, 0x56, 0x2a, 0xb9, 0x81, 0xe8, 0xb6, 0x90, 0xea, 0x6e, 0x8f,
	0xcd, 0xab, 0x5e, 0xcd, 0x8b, 0xb2, 0x50, 0xf6, 0xf2, 0x16, 0xe8, 0x23, 0x05, 0xdb, 0xa2, 0x3d,
	0xdd, 0xd4, 0x9a, 0x93, 0x75, 0xa3, 0x0e, 0x87, 0xeb, 0x3a, 0x79, 0x81, 0x51, 0x86, 0xbb, 0x3d,
	0x4a, 0x45, 0xce, 0x00, 0x78, 0x21, 0xd5, 0x66, 0xa7, 0xc7, 0x9a, 0x69, 0x71, 0xfa, 0x77, 0xe1,
	0x24, 0x78, 0xdc, 0x99, 0x45, 0xfc, 0xb8, 0xfe, 0xc2, 0xcd, 0xd0, 0xac, 0x8b, 0xd3, 0xa9, 0xeb,
	0xea, 0xd4, 0xcc, 0xe9, 0x4c, 0xde, 0x3c, 0x18, 0x67, 0x28, 0x45, 0x5d, 0x49, 0xd4, 0x43, 0x3a,
	0x87, 0x5d, 0xfd, 0xed, 0x90, 0x8e, 0x26, 0x97, 0x10, 0x77, 0x48, 0x52, 0x7f, 0x16, 0xfc, 0x60,
	0x74, 0x5b, 0x75, 0x6a, 0xcf, 0x8c, 0x17, 0xb9, 0x09, 0x63, 0x9c, 0xb5, 0x20, 0xfd, 0xf0, 0x21,
	0x5e, 0x3b, 0x5d, 0x29, 0x04, 0x2b, 0xce, 0xc9, 0x1f, 0x77, 0xa2, 0x8d, 0xeb, 0xdf, 0xe4, 0x2b,
	0xd9, 0x7e, 0x49, 0xf2, 0x8b, 0x2c, 0x61, 0xa0, 0x83, 0xea, 0x67, 0x4a, 0x21, 0xb8, 0xc6, 0x9e,
	0x9e, 0x73, 0x08, 0xaf, 0xcc, 0x8b, 0xe9, 0x6d, 0xbb, 0x37, 0x2f, 0xa9, 0xb7, 0x6d, 0x8d, 0x1c,
	0x7b, 0xda, 0x1e, 0x42, 0xf3, 0xdb, 0x2d, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xda, 0x6d, 0xeb,
	0xa5, 0x9b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Departments service

type DepartmentsClient interface {
	// 全部部门
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取部门编码列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取部门编码信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建部门编码
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新部门编码
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除部门编码
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type departmentsClient struct {
	c           client.Client
	serviceName string
}

func NewDepartmentsClient(serviceName string, c client.Client) DepartmentsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "department"
	}
	return &departmentsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *departmentsClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Departments.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Departments service

type DepartmentsHandler interface {
	// 全部部门
	All(context.Context, *Request, *Response) error
	// 获取部门编码列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取部门编码信息
	Get(context.Context, *Request, *Response) error
	// 创建部门编码
	Create(context.Context, *Request, *Response) error
	// 更新部门编码
	Update(context.Context, *Request, *Response) error
	// 删除部门编码
	Delete(context.Context, *Request, *Response) error
}

func RegisterDepartmentsHandler(s server.Server, hdlr DepartmentsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Departments{hdlr}, opts...))
}

type Departments struct {
	DepartmentsHandler
}

func (h *Departments) All(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.All(ctx, in, out)
}

func (h *Departments) List(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.List(ctx, in, out)
}

func (h *Departments) Get(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.Get(ctx, in, out)
}

func (h *Departments) Create(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.Create(ctx, in, out)
}

func (h *Departments) Update(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.Update(ctx, in, out)
}

func (h *Departments) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.DepartmentsHandler.Delete(ctx, in, out)
}
