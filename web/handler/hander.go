package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"xinshoucun/web/model"
	"xinshoucun/web/service"
)

// 初始化对象
var ObjectService *service.Object

// 添加数据
func AddOrder(c *gin.Context) {

	rs := model.Order{}
	if err := c.ShouldBind(&rs); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}

	err := ObjectService.AddOrder(rs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// 查询单条数据
func OrderDetail(c *gin.Context) {

	get := model.Order{
		ID:       1,
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}
	if err := c.ShouldBind(&get); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}

	rs, err := ObjectService.OrderDetail(get.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "未找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": rs,
	})
}

// 查询列表数据
func OrderList(c *gin.Context) {
	listRS := model.SearchList{}
	if err := c.ShouldBind(&listRS); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}

	rs, err := ObjectService.OrderList(model.SearchList{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": rs,
	})
}

// 更新数据
func UpdateOrder(c *gin.Context) {

	updaters := model.NewUpdate{}
	if err := c.ShouldBind(&updaters); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}

	err := ObjectService.UpdateOrder(updaters)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// 文件上传
func Upload(c *gin.Context) {

	rs := model.Order{}
	if err := c.ShouldBind(&rs.ID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}
	file, _ := c.FormFile("file")

	times := time.Now()
	// 文件夹路径
	fileDir := fmt.Sprintf("upload/%d/%d/%d", times.Year(), times.Month(), times.Day())
	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err := os.MkdirAll(fileDir, os.ModePerm)
	// 文件路径
	timeStamp := time.Now().Unix()
	FileName := fmt.Sprintf("%d-%s", timeStamp, file.Filename)
	filePathStr := filepath.Join(fileDir, FileName)

	err = c.SaveUploadedFile(file, filePathStr)
	fileUrl := fmt.Sprintf("utils/upload/%d/%d/%d/%s", times.Year(), times.Month(), times.Day(), FileName)
	err2 := ObjectService.UpdateFileUrl(rs.ID, fileUrl)
	if err2 != nil {
		fmt.Println(err2)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "文件上传失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s uploaded!", file.Filename),
	})
}

// 文件下载
func DownloadFile(c *gin.Context) {
	// 通过传入id下载 fileUrl 的文件
	rs := model.Order{}
	if err := c.ShouldBind(&rs.ID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数不正确",
		})
		return
	}
	res, err := ObjectService.DownloadFile(rs.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "下载文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s download!", res),
	})
}
