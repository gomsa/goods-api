// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Origin               string   `protobuf:"bytes,8,opt,name=origin,proto3" json:"origin,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
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

func (m *ListQuery) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ListQuery) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *ListQuery) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
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

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
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

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*ListQuery)(nil), "user.ListQuery")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x5d, 0x7f, 0xc6, 0x9e, 0x85, 0x2c, 0x88, 0xdd, 0x45, 0x04, 0x76, 0x09, 0x5e, 0xb6, 0xb4,
	0x97, 0x14, 0xd2, 0x5f, 0x10, 0xda, 0xd2, 0x4b, 0x2f, 0x15, 0xe4, 0x5c, 0x94, 0x5a, 0x04, 0x81,
	0x1d, 0xb9, 0x92, 0x9c, 0xa6, 0xbf, 0xa3, 0xff, 0xb5, 0xc7, 0x52, 0x34, 0xb2, 0x43, 0x5a, 0x52,
	0x7c, 0xb1, 0xe7, 0xbd, 0x37, 0x7e, 0x1a, 0xbf, 0x41, 0xf0, 0xab, 0xd1, 0xca, 0xaa, 0xf3, 0xd6,
	0x08, 0x8d, 0x8f, 0x19, 0x62, 0x12, 0xbb, 0xba, 0x78, 0x0b, 0x20, 0x5e, 0x1a, 0xa1, 0xc9, 0x18,
	0x42, 0x59, 0xd2, 0x60, 0x1a, 0x9c, 0xe6, 0x2c, 0x94, 0x25, 0x99, 0x40, 0xe6, 0x1a, 0x36, 0xbc,
	0x16, 0x34, 0x44, 0x76, 0x8f, 0xc9, 0x6f, 0x48, 0x6b, 0xb5, 0x92, 0x95, 0xa0, 0x11, 0x2a, 0x1d,
	0x22, 0x3f, 0x21, 0x11, 0x35, 0x97, 0x15, 0x8d, 0x91, 0xf6, 0xc0, 0x39, 0x35, 0xdc, 0x98, 0x27,
	0xa5, 0x4b, 0x9a, 0x78, 0xa7, 0x1e, 0x13, 0x02, 0x31, 0x9e, 0x90, 0x22, 0x1f, 0xf7, 0xee, 0x7c,
	0xcb, 0x2d, 0xd7, 0x74, 0xe4, 0xdd, 0x3d, 0x72, 0xbc, 0xd2, 0x72, 0x2d, 0x37, 0x34, 0xf3, 0xbc,
	0x47, 0xe4, 0x0f, 0xc0, 0x83, 0x16, 0xdc, 0x8a, 0xf2, 0x9e, 0x5b, 0x9a, 0xa3, 0x96, 0x77, 0xcc,
	0xc2, 0x3a, 0xb9, 0x6d, 0xca, 0x5e, 0x06, 0x2f, 0x77, 0xcc, 0xc2, 0x16, 0x2f, 0x01, 0xe4, 0xb7,
	0xd2, 0xd8, 0xbb, 0x56, 0xe8, 0x67, 0xf7, 0x07, 0x95, 0xac, 0xa5, 0xc5, 0x20, 0x22, 0xe6, 0x81,
	0x9b, 0xb2, 0xe1, 0x6b, 0x9f, 0x43, 0xc4, 0xb0, 0x76, 0x9c, 0x51, 0xda, 0x76, 0x09, 0x60, 0xfd,
	0x21, 0xb3, 0xf8, 0xcb, 0xcc, 0x92, 0xe3, 0x99, 0xa5, 0x07, 0x99, 0x15, 0x39, 0x8c, 0x98, 0x78,
	0x6c, 0x85, 0xb1, 0xc5, 0x0e, 0x32, 0x26, 0x4c, 0xa3, 0x36, 0x46, 0x90, 0xbf, 0x80, 0x5b, 0xc3,
	0xe9, 0xbe, 0xcf, 0x61, 0x86, 0xeb, 0x74, 0xeb, 0x63, 0xc8, 0x93, 0x29, 0x24, 0xee, 0x6d, 0x68,
	0x38, 0x8d, 0x3e, 0x35, 0x78, 0xc1, 0x1d, 0x67, 0x95, 0xe5, 0x15, 0xce, 0x1d, 0x31, 0x0f, 0x1c,
	0xbb, 0xe5, 0x95, 0x2c, 0x71, 0xea, 0x8c, 0x79, 0x30, 0x7f, 0x0d, 0x20, 0x59, 0xe2, 0x57, 0xff,
	0x21, 0xb9, 0xde, 0x49, 0x63, 0xc9, 0x81, 0xe3, 0x64, 0xec, 0xeb, 0x7e, 0xb8, 0xe2, 0x1b, 0x39,
	0x83, 0xd8, 0x45, 0x49, 0x7e, 0x78, 0x65, 0x1f, 0xeb, 0x91, 0xd6, 0x7f, 0x10, 0xdd, 0x88, 0x21,
	0xbf, 0x13, 0x48, 0x2f, 0x71, 0x8f, 0xc3, 0x7d, 0x4b, 0x5c, 0xe8, 0x70, 0xdf, 0x95, 0xa8, 0xc4,
	0x50, 0xdf, 0x2a, 0xc5, 0x1b, 0x72, 0xf1, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x39, 0x4b, 0x73,
	0x3a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersClient interface {
	// 用过 用户名、邮箱、手机 查询用户是否存在
	Exist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 获取用户列表
	List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取用户
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 创建用户
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 更新用户
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 删除用户
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type usersClient struct {
	c           client.Client
	serviceName string
}

func NewUsersClient(serviceName string, c client.Client) UsersClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "user"
	}
	return &usersClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *usersClient) Exist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Exist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	// 用过 用户名、邮箱、手机 查询用户是否存在
	Exist(context.Context, *User, *Response) error
	// 获取用户列表
	List(context.Context, *ListQuery, *Response) error
	// 根据 唯一 获取用户
	Get(context.Context, *User, *Response) error
	// 创建用户
	Create(context.Context, *User, *Response) error
	// 更新用户
	Update(context.Context, *User, *Response) error
	// 删除用户
	Delete(context.Context, *User, *Response) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Users{hdlr}, opts...))
}

type Users struct {
	UsersHandler
}

func (h *Users) Exist(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Exist(ctx, in, out)
}

func (h *Users) List(ctx context.Context, in *ListQuery, out *Response) error {
	return h.UsersHandler.List(ctx, in, out)
}

func (h *Users) Get(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Get(ctx, in, out)
}

func (h *Users) Create(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Create(ctx, in, out)
}

func (h *Users) Update(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Update(ctx, in, out)
}

func (h *Users) Delete(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Delete(ctx, in, out)
}
