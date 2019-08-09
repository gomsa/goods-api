package hander

import (
	"context"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/barcode"
)

// Barcode 商品结构
type Barcode struct {
	ServiceName string
}

// Get 获取商品信息
func (srv *Barcode) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Barcode.Get", req, res)
}
