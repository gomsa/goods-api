package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	Service: "goods-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "goods-api", Method: "Barcodes.Get", Auth: true, Policy: true, Name: "获取商品信息", Description: "通过条码获取商品信息权限"},

		{Service: "goods-api", Method: "Goods.GoodsByBarcode", Auth: true, Policy: true, Name: "条形码查询商品", Description: "根据条形码查询商品权限"},
		{Service: "goods-api", Method: "Goods.Exist", Auth: true, Policy: true, Name: "查询条形码是否存在", Description: "查询条形码是否存在权限"},
		{Service: "goods-api", Method: "Goods.DeleteBarcode", Auth: true, Policy: true, Name: "删除商品条码", Description: "删除商品下面对应条码的子商品"},
		{Service: "goods-api", Method: "Goods.List", Auth: true, Policy: true, Name: "商品列表", Description: "商品查询列表权限"},
		{Service: "goods-api", Method: "Goods.Create", Auth: true, Policy: true, Name: "商品创建", Description: "商品创建权限"},
		{Service: "goods-api", Method: "Goods.Get", Auth: true, Policy: true, Name: "商品查询", Description: "商品查询信息权限"},
		{Service: "goods-api", Method: "Goods.Update", Auth: true, Policy: true, Name: "商品更新", Description: "商品更新信息权限"},
		{Service: "goods-api", Method: "Goods.Delete", Auth: true, Policy: true, Name: "商品删除", Description: "商品删除权限"},

		{Service: "goods-api", Method: "Brand.Exist", Auth: true, Policy: true, Name: "查询品牌是否存在", Description: "查询品牌是否存在权限"},
		{Service: "goods-api", Method: "Brand.All", Auth: true, Policy: true, Name: "品牌全部", Description: "品牌查询全部权限"},
		{Service: "goods-api", Method: "Brand.List", Auth: true, Policy: true, Name: "品牌列表", Description: "品牌查询列表权限"},
		{Service: "goods-api", Method: "Brand.Create", Auth: true, Policy: true, Name: "品牌创建", Description: "品牌创建权限"},
		{Service: "goods-api", Method: "Brand.Get", Auth: true, Policy: true, Name: "品牌查询", Description: "品牌查询信息权限"},
		{Service: "goods-api", Method: "Brand.Update", Auth: true, Policy: true, Name: "品牌更新", Description: "品牌更新信息权限"},
		{Service: "goods-api", Method: "Brand.Delete", Auth: true, Policy: true, Name: "品牌删除", Description: "品牌删除权限"},

		{Service: "goods-api", Method: "Category.All", Auth: true, Policy: true, Name: "分类全部", Description: "分类查询全部权限"},
		{Service: "goods-api", Method: "Category.List", Auth: true, Policy: true, Name: "分类列表", Description: "分类查询列表权限"},
		{Service: "goods-api", Method: "Category.Create", Auth: true, Policy: true, Name: "分类创建", Description: "分类创建权限"},
		{Service: "goods-api", Method: "Category.Get", Auth: true, Policy: true, Name: "分类查询", Description: "分类查询信息权限"},
		{Service: "goods-api", Method: "Category.Update", Auth: true, Policy: true, Name: "分类更新", Description: "分类更新信息权限"},
		{Service: "goods-api", Method: "Category.Delete", Auth: true, Policy: true, Name: "分类删除", Description: "分类删除权限"},

		{Service: "goods-api", Method: "Department.All", Auth: true, Policy: true, Name: "部门全部", Description: "部门查询全部权限"},
		{Service: "goods-api", Method: "Department.List", Auth: true, Policy: true, Name: "部门列表", Description: "部门查询列表权限"},
		{Service: "goods-api", Method: "Department.Create", Auth: true, Policy: true, Name: "部门创建", Description: "部门创建权限"},
		{Service: "goods-api", Method: "Department.Get", Auth: true, Policy: true, Name: "部门查询", Description: "部门查询信息权限"},
		{Service: "goods-api", Method: "Department.Update", Auth: true, Policy: true, Name: "部门更新", Description: "部门更新信息权限"},
		{Service: "goods-api", Method: "Department.Delete", Auth: true, Policy: true, Name: "部门删除", Description: "部门删除权限"},

		{Service: "goods-api", Method: "Firm.Exist", Auth: true, Policy: true, Name: "查询商品公司是否存在", Description: "查询商品公司是否存在权限"},
		{Service: "goods-api", Method: "Firm.All", Auth: true, Policy: true, Name: "商品公司全部", Description: "商品公司查询全部权限"},
		{Service: "goods-api", Method: "Firm.List", Auth: true, Policy: true, Name: "商品公司列表", Description: "商品公司查询列表权限"},
		{Service: "goods-api", Method: "Firm.Create", Auth: true, Policy: true, Name: "商品公司创建", Description: "商品公司创建权限"},
		{Service: "goods-api", Method: "Firm.Get", Auth: true, Policy: true, Name: "商品公司查询", Description: "商品公司查询信息权限"},
		{Service: "goods-api", Method: "Firm.Update", Auth: true, Policy: true, Name: "商品公司更新", Description: "商品公司更新信息权限"},
		{Service: "goods-api", Method: "Firm.Delete", Auth: true, Policy: true, Name: "商品公司删除", Description: "商品公司删除权限"},

		{Service: "goods-api", Method: "Taxcode.All", Auth: true, Policy: true, Name: "国家税收分类编码全部", Description: "国家税收分类编码查询全部权限"},
		{Service: "goods-api", Method: "Taxcode.List", Auth: true, Policy: true, Name: "国家税收分类编码列表", Description: "国家税收分类编码查询列表权限"},
		{Service: "goods-api", Method: "Taxcode.Create", Auth: true, Policy: true, Name: "国家税收分类编码创建", Description: "国家税收分类编码创建权限"},
		{Service: "goods-api", Method: "Taxcode.Get", Auth: true, Policy: true, Name: "国家税收分类编码查询", Description: "国家税收分类编码查询信息权限"},
		{Service: "goods-api", Method: "Taxcode.Update", Auth: true, Policy: true, Name: "国家税收分类编码更新", Description: "国家税收分类编码更新信息权限"},
		{Service: "goods-api", Method: "Taxcode.Delete", Auth: true, Policy: true, Name: "国家税收分类编码删除", Description: "国家税收分类编码删除权限"},

		{Service: "goods-api", Method: "Unspsc.CheckCreate", Auth: true, Policy: true, Name: "国际商品及服务编码批量创建", Description: "国际商品及服务编码批量创建(一组创建)权限"},
		{Service: "goods-api", Method: "Unspsc.Exist", Auth: true, Policy: true, Name: "国际商品及服务编码是否存在", Description: "国际商品及服务编码是否存在权限"},
		{Service: "goods-api", Method: "Unspsc.All", Auth: true, Policy: true, Name: "国际商品及服务编码全部", Description: "国际商品及服务编码查询全部权限"},
		{Service: "goods-api", Method: "Unspsc.List", Auth: true, Policy: true, Name: "国际商品及服务编码列表", Description: "国际商品及服务编码查询列表权限"},
		{Service: "goods-api", Method: "Unspsc.Create", Auth: true, Policy: true, Name: "国际商品及服务编码创建", Description: "国际商品及服务编码创建权限"},
		{Service: "goods-api", Method: "Unspsc.Get", Auth: true, Policy: true, Name: "国际商品及服务编码查询", Description: "国际商品及服务编码查询信息权限"},
		{Service: "goods-api", Method: "Unspsc.Update", Auth: true, Policy: true, Name: "国际商品及服务编码更新", Description: "国际商品及服务编码更新信息权限"},
		{Service: "goods-api", Method: "Unspsc.Delete", Auth: true, Policy: true, Name: "国际商品及服务编码删除", Description: "国际商品及服务编码删除权限"},
	},
}
