package service

import (
	"fmt"
	"testing"
	"xinshoucun/web/dao"
	"xinshoucun/web/model"
)

// 初始化数据库连接
func InitService() *Object {

	// 初始化数据库,建库，建表
	// 初始化对象
	obj := dao.NewMyConnect()
	rs := NewObject(obj)
	return rs
}

// 添加
func TestObject_AddOrder(t *testing.T) {

	Object := InitService()

	rs := model.Order{
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}

	err := Object.AddOrder(rs)
	fmt.Println("err", err)
}

// 数据
func TestObject_OrderDetail(t *testing.T) {

	Object := InitService()

	var id uint = 1
	rs, err := Object.OrderDetail(id)
	if err != nil {
		fmt.Println("err", err)
	}
	if id != rs.ID {
		fmt.Println("id != rs.ID err")
	}
}

// 列表
func TestObject_OrderList(t *testing.T) {
	Object := InitService()
	tmp := model.SearchList{
		UserName: "fengshi",
		Page:     1,
		Limit:    1,
	}
	_, err := Object.OrderList(tmp)
	if err != nil {
		fmt.Println("err", err)
	}
}

// 更新
func TestObject_UpdateOrder(t *testing.T) {

	Object := InitService()

	rs := map[string]interface{}{"Amount": 1, "Status": "1", "FileUrl": "1"}
	err := Object.UpdateOrder(rs)
	fmt.Println("err", err)

}

// 更新url
func TestObject_UpdateFileUrl(t *testing.T) {
	Object := InitService()
	id := 10
	url := "123456"
	err := Object.UpdateFileUrl(uint(id), url)
	fmt.Println("err", err)
}
