package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	Service: "goods-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "goods-api", Method: "Goods.Create", Auth: false, Policy: true, Name: "创建商品", Description: "创建新商品权限。"},
		{Service: "goods-api", Method: "Goods.Get", Auth: true, Policy: true, Name: "查询商品", Description: "查询商品信息权限。"},
		{Service: "goods-api", Method: "Goods.Update", Auth: true, Policy: true, Name: "更新商品", Description: "更新商品信息。"},
		{Service: "goods-api", Method: "Goods.Delete", Auth: true, Policy: true, Name: "删除商品", Description: "删除商品。"},
		{Service: "goods-api", Method: "Goods.List", Auth: true, Policy: true, Name: "商品列表", Description: "查询商品列表"},
	},
}
