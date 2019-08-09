package main

import (
	// 公共引入
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	userClient "github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"

	"github.com/gomsa/goods-api/hander"
	// 接口引用
	barcodePB "github.com/gomsa/goods-api/proto/barcode"
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

	ServiceName := os.Getenv("GOODS_NAME")
	// 服务实现
	barcodePB.RegisterBarcodesHandler(srv.Server(), &hander.Barcode{ServiceName})
	goodsPB.RegisterGoodsHandler(srv.Server(), &hander.Goods{ServiceName})
	brandPB.RegisterBrandsHandler(srv.Server(), &hander.Brand{ServiceName})
	firmPB.RegisterFirmsHandler(srv.Server(), &hander.Firm{ServiceName})
	categoryPB.RegisterCategorysHandler(srv.Server(), &hander.Category{ServiceName})
	departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{ServiceName})
	unspscPB.RegisterUnspscsHandler(srv.Server(), &hander.Unspsc{ServiceName})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	if err := userClient.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
