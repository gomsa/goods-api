package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/firm"
)

// Firm 公司结构
type Firm struct {
	ServiceName string
}

// Exist 查询公司否存在
func (srv *Firm) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.Exist", req, res)
}

// All 根据条形码查询公司
func (srv *Firm) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.All", req, res)
}

// List 获取所有公司
func (srv *Firm) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.List", req, res)
}

// Get 获取公司
func (srv *Firm) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.Get", req, res)
}

// Create 创建公司
func (srv *Firm) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.Create", req, res)
}

// Update 更新公司
func (srv *Firm) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.Update", req, res)
}

// Delete 删除公司公司
func (srv *Firm) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Firms.Delete", req, res)
}
