package hander

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"github.com/gomsa/goods/client"
	goodsPB "github.com/gomsa/goods/proto/goods"
	"github.com/gomsa/tools/uitl"

	pb "github.com/gomsa/goods-api/proto/goods"
)

// Goods 部门结构
type Goods struct {
}

// List 部门列表
func (srv *Goods) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &goodsPB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		log.Log(err)
		return err
	}
	goodsRes, err := client.Goods.List(ctx, query)
	if err != nil {
		log.Log(err)
		return err
	}
	err = uitl.Data2Data(goodsRes, res)
	if err != nil {
		log.Log(err)
		return err
	}
	return err
}

// Get 获取部门
func (srv *Goods) Get(ctx context.Context, req *pb.Goods, res *pb.Response) (err error) {
	goods := &goodsPB.Goods{}
	err = uitl.Data2Data(req, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	goodsRes, err := client.Goods.Get(ctx, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	err = uitl.Data2Data(goodsRes, res)
	if err != nil {
		log.Log(err)
		return err
	}
	return err
}

// Create 创建部门
func (srv *Goods) Create(ctx context.Context, req *pb.Goods, res *pb.Response) (err error) {
	goods := &goodsPB.Goods{}
	err = uitl.Data2Data(req, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	goodsRes, err := client.Goods.Create(ctx, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	err = uitl.Data2Data(goodsRes, res)
	if err != nil {
		log.Log(err)
		return err
	}
	return err
}

// Update 更新部门
func (srv *Goods) Update(ctx context.Context, req *pb.Goods, res *pb.Response) (err error) {
	goods := &goodsPB.Goods{}
	err = uitl.Data2Data(req, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	goodsRes, err := client.Goods.Update(ctx, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	err = uitl.Data2Data(goodsRes, res)
	if err != nil {
		log.Log(err)
		return err
	}
	return err
}

// Delete 删除部门
func (srv *Goods) Delete(ctx context.Context, req *pb.Goods, res *pb.Response) (err error) {
	goods := &goodsPB.Goods{
		Id: req.Id,
	}
	if err != nil {
		log.Log(err)
		return err
	}
	goodsRes, err := client.Goods.Delete(ctx, goods)
	if err != nil {
		log.Log(err)
		return err
	}
	err = uitl.Data2Data(goodsRes, res)
	if err != nil {
		log.Log(err)
		return err
	}
	return err
}
