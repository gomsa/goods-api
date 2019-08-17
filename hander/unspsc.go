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

// CheckCreate 检查批量创建
func (srv *Unspsc) CheckCreate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.CheckCreate", req, res)
}

// Exist 检查国际商品及服务编码
func (srv *Unspsc) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.Exist", req, res)
}

// All 根据条形码查询国际分类
func (srv *Unspsc) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.All", req, res)
}

// List 获取所有国际分类
func (srv *Unspsc) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.List", req, res)
}

// Get 获取国际分类
func (srv *Unspsc) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.Get", req, res)
}

// Create 创建国际分类
func (srv *Unspsc) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.Create", req, res)
}

// Update 更新国际分类
func (srv *Unspsc) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.Update", req, res)
}

// Delete 删除国际分类国际分类
func (srv *Unspsc) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Unspscs.Delete", req, res)
}
