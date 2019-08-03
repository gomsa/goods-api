package hander

import (
	"context"

	client "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods-api/proto/goods"
)

// Goods 部门结构
type Goods struct {
	C           client.Client
	ServiceName string
}

// GoodsByBarcode 根据条形码查询商品
func (srv *Goods) GoodsByBarcode(ctx context.Context, request *pb.Request, out *pb.Response) (err error) {
	req := srv.C.NewRequest(srv.ServiceName, "Goods.GoodsByBarcode", request)
	err = srv.C.Call(ctx, req, out)
	log.Log(ctx, req, out)
	return err
}

// Exist 查询条码是否存在
func (srv *Goods) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// DeleteBarcode 删除条码
func (srv *Goods) DeleteBarcode(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// List 获取所有商品
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Get 获取商品
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Create 创建商品
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Update 更新商品
func (srv *Goods) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Delete 删除商品商品
func (srv *Goods) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}
