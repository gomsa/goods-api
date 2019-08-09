package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/brand"
)

// Brand 品牌结构
type Brand struct {
	ServiceName string
}

// Exist 查询品牌否存在
func (srv *Brand) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.Exist", req, res)
}

// All 根据条形码查询品牌
func (srv *Brand) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.All", req, res)
}

// List 获取所有品牌
func (srv *Brand) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.List", req, res)
}

// Get 获取品牌
func (srv *Brand) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.Get", req, res)
}

// Create 创建品牌
func (srv *Brand) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.Create", req, res)
}

// Update 更新品牌
func (srv *Brand) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brands.Update", req, res)
}

// Delete 删除品牌品牌
func (srv *Brand) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Brand.Delete", req, res)
}
