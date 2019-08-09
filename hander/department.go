package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/department"
)

// Department 部门结构
type Department struct {
	ServiceName string
}

// All 根据条形码查询部门
func (srv *Department) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.All", req, res)
}

// List 获取所有部门
func (srv *Department) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.List", req, res)
}

// Get 获取部门
func (srv *Department) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.Get", req, res)
}

// Create 创建部门
func (srv *Department) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.Create", req, res)
}

// Update 更新部门
func (srv *Department) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.Update", req, res)
}

// Delete 删除部门部门
func (srv *Department) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Departments.Delete", req, res)
}
