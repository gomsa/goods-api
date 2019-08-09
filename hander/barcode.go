package hander

import (
	"context"
	"encoding/json"
	"fmt"

	client "github.com/gomsa/tools/k8s/client"
	"github.com/micro/go-micro/util/log"

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
	log.Log(label)
	log.Log(req)
	if r, err := redis.Get(label).Result(); err != nil && r != "" {
		log.Log(r)
		log.Log(err)
		json.Unmarshal([]byte(r), res)
	} else {
		err = client.Call(ctx, srv.ServiceName, "Barcode.Get", req, res)
		log.Log(res, err)
		redis.SetNX(label, res, 0)
	}
	return err
}
