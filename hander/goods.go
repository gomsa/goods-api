package hander

import (
	"context"
	"fmt"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/goods"
)

// Goods 商品结构
type Goods struct {
	ServiceName string
}

// GoodsByBarcode 根据条形码查询商品
func (srv *Goods) GoodsByBarcode(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.GoodsByBarcode", req, res)
}

// Exist 查询条码是否存在
func (srv *Goods) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.Exist", req, res)
}

// DeleteBarcode 删除条码
func (srv *Goods) DeleteBarcode(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.DeleteBarcode", req, res)
}

// List 获取所有商品
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.List", req, res)
}

// Get 获取商品
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.Get", req, res)
}

// Create 创建商品
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	if req.Good.Code == "" {
		return fmt.Errorf("自编码不允许为空")
	}
	if req.Good.Name == "" {
		return fmt.Errorf("商品名称不允许为空")
	}
	for _, barcode := range req.Good.Barcodes {
		if barcode.Id == "" {
			return fmt.Errorf("条形码不允许为空")
		}
		if barcode.Price == 0 {
			return fmt.Errorf("%s 商品价格不允许为空", barcode.Id)
		}
	}
	return client.Call(ctx, srv.ServiceName, "Goods.Create", req, res)
}

// Update 更新商品
func (srv *Goods) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.Update", req, res)
}

// Delete 删除商品商品
func (srv *Goods) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Goods.Delete", req, res)
}
