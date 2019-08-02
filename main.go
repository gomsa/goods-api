package main

import (
	// 公共引入
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/goods-api/hander"
	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"

	// 接口引用
	// brandPB "github.com/gomsa/goods-api/proto/brand"
	// categoryPB "github.com/gomsa/goods-api/proto/category"
	// departmentPB "github.com/gomsa/goods-api/proto/department"
	// firmPB "github.com/gomsa/goods-api/proto/firm"
	goodsPB "github.com/gomsa/goods-api/proto/goods"
	// unspscPB "github.com/gomsa/goods-api/proto/unspsc"

	"github.com/gomsa/tools/config"
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
	serviceName := os.Getenv("USER_NAME")
	// 服务实现
	goodsPB.RegisterGoodsHandler(srv.Server(), &hander.Goods{serviceName})
	// brandPB.RegisterBrandsHandler(srv.Server(), &hander.Brand{serviceName})
	// firmPB.RegisterFirmsHandler(srv.Server(), &hander.Firm{serviceName})
	// categoryPB.RegisterCategorysHandler(srv.Server(), &hander.Category{serviceName})
	// departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{serviceName})
	// unspscPB.RegisterUnspscsHandler(srv.Server(), &hander.Unspsc{serviceName})

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
