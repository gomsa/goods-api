package hander

import (
	"context"
	"encoding/json"

	client "github.com/gomsa/tools/k8s/client"

	pb "github.com/gomsa/goods-api/proto/barcode"
	"github.com/gomsa/goods-api/providers/redis"
)

// Barcode 商品结构
type Barcode struct {
	ServiceName string
}

// Get 获取商品信息
func (srv *Barcode) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	label := "barcode:" + req.Goods.Barcode
	redis := redis.NewClient()
	if r, err := redis.Get(label).Result(); err != nil && r == "" {
		json.Unmarshal([]byte(r), res)
	} else {
		err = client.Call(ctx, srv.ServiceName, "Barcode.Get", req, res)
		redis.SetNX(label, res, 0)
	}
	return err
}
