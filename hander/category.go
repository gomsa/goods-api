package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/category"
)

// Category 分类结构
type Category struct {
	ServiceName string
}

// All 根据条形码查询分类
func (srv *Category) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.All", req, res)
}

// List 获取所有分类
func (srv *Category) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.List", req, res)
}

// Get 获取分类
func (srv *Category) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.Get", req, res)
}

// Create 创建分类
func (srv *Category) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.Create", req, res)
}

// Update 更新分类
func (srv *Category) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.Update", req, res)
}

// Delete 删除分类分类
func (srv *Category) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Category.Delete", req, res)
}
