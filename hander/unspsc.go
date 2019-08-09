package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/unspsc"
)

// Unspsc 国际分类结构
type Unspsc struct {
	ServiceName string
}

// All 根据条形码查询国际分类
func (srv *Unspsc) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.All", req, res)
}

// List 获取所有国际分类
func (srv *Unspsc) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.List", req, res)
}

// Get 获取国际分类
func (srv *Unspsc) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.Get", req, res)
}

// Create 创建国际分类
func (srv *Unspsc) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.Create", req, res)
}

// Update 更新国际分类
func (srv *Unspsc) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.Update", req, res)
}

// Delete 删除国际分类国际分类
func (srv *Unspsc) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspsc.Delete", req, res)
}