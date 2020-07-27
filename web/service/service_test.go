package service

import (
	"fmt"
	"testing"
	"xinshoucun/web/dao"
	"xinshoucun/web/db"
	"xinshoucun/web/model"
)

// 初始化数据库连接
func InitService() *Object {

	dbLink, _ := db.BuildDb()
	Dao := dao.Connect{Db: dbLink}
	Object := &Object{dao: Dao}
	return Object
}

// 添加
func TestObject_AddOrder(t *testing.T) {

	Object := InitService()

	req := model.Order{
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}

	err := Object.AddOrder(req)
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
		panic("id != rs.ID, err")
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

	req := map[string]interface{}{"Amount": 1, "Status": "1", "FileUrl": "1"}
	err := Object.UpdateOrder(req)
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
