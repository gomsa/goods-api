package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/taxcode"
)

// Taxcode 国家税收分类编码结构
type Taxcode struct {
	ServiceName string
}

// All 根据条形码查询国家税收分类编码
func (srv *Taxcode) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.All", req, res)
}

// List 获取所有国家税收分类编码
func (srv *Taxcode) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.List", req, res)
}

// Get 获取国家税收分类编码
func (srv *Taxcode) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.Get", req, res)
}

// Create 创建国家税收分类编码
func (srv *Taxcode) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.Create", req, res)
}

// Update 更新国家税收分类编码
func (srv *Taxcode) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.Update", req, res)
}

// Delete 删除国家税收分类编码
func (srv *Taxcode) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Taxcode.Delete", req, res)
}
