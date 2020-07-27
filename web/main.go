package main

import (
	"xinshoucun/web/dao"
	"xinshoucun/web/handler"
	"xinshoucun/web/router"
	"xinshoucun/web/service"
)

func main() {
	// 初始化数据库,建库，建表
	dao.NewMyConnect()
	// 初始化对象
	obj := dao.NewMyConnect()
	handler.ObjectService = service.NewObject(obj)
	// 初始化router
	router := router.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		panic("engine Run err")
	}
}
