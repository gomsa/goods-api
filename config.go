package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	Service: "goods-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "goods-api", Method: "Goodss.All", Auth: true, Policy: true, Name: "所有部门", Description: "获取所有部门,可以根据条件。"},
		{Service: "goods-api", Method: "Goodss.Create", Auth: false, Policy: true, Name: "创建部门", Description: "创建新部门权限。"},
		{Service: "goods-api", Method: "Goodss.Get", Auth: true, Policy: true, Name: "查询部门", Description: "查询部门信息权限。"},
		{Service: "goods-api", Method: "Goodss.Update", Auth: true, Policy: true, Name: "更新部门", Description: "更新部门信息。"},
		{Service: "goods-api", Method: "Goodss.Delete", Auth: true, Policy: true, Name: "删除部门", Description: "删除部门。"},
		{Service: "goods-api", Method: "Goodss.List", Auth: true, Policy: true, Name: "部门列表", Description: "查询部门列表"},
	},
}
