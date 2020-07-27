package router

import (
	"github.com/gin-gonic/gin"
	"xinshoucun/web/handler"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	// 新增
	router.POST("/add", handler.AddOrder)
	// 获取详情
	router.POST("/detail", handler.OrderDetail)
	// 获取列表
	router.POST("/list", handler.OrderList)
	// 更新
	router.POST("/update", handler.UpdateOrder)
	// 上传s
	router.POST("/upload", handler.Upload)
	// 下载
	router.POST("/download", handler.DownloadFile)

	return router
}
