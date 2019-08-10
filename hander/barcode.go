package hander

import (
	"context"
	"encoding/json"
	"fmt"

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
	// 防止空指针报错
	if req.Goods == nil {
		return fmt.Errorf("请求参数错误")
	}
	label := "barcode:" + req.Goods.Barcode
	redis := redis.NewClient()
	if r, err := redis.Get(label).Result(); err == nil && r != "" {
		json.Unmarshal([]byte(r), res)
	} else {
		err = client.Call(ctx, srv.ServiceName, "Barcodes.Get", req, res)
		if err != nil {
			return err
		}
		j, _ := json.Marshal(res)
		err = redis.SetNX(label, j, 0).Err()
		if err != nil {
			return err
		}
	}
	return err
}
