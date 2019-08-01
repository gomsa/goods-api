package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/goods-api/hander"
	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"

	// 接口引用
	brandPB "github.com/gomsa/goods-api/proto/brand"
	categoryPB "github.com/gomsa/goods-api/proto/category"
	departmentPB "github.com/gomsa/goods-api/proto/department"
	firmPB "github.com/gomsa/goods-api/proto/firm"
	goodsPB "github.com/gomsa/goods-api/proto/goods"
	unspscPB "github.com/gomsa/goods-api/proto/unspsc"
)

func main() {
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
	}
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()

	// 服务实现
	goodsPB.RegisterGoodsHandler(srv.Server(), &hander.Goods{})
	brandPB.RegisterBrandsHandler(srv.Server(), &hander.Brand{})
	firmPB.RegisterFirmsHandler(srv.Server(), &hander.Firm{})
	categoryPB.RegisterCategorysHandler(srv.Server(), &hander.Category{})
	departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{})
	unspscPB.RegisterUnspscsHandler(srv.Server(), &hander.Unspsc{})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	if err := client.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
