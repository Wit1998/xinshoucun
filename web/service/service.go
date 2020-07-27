package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"xinshoucun/web/dao"
	"xinshoucun/web/model"
)

type Object struct {
	dao *dao.Connect
}

// 初始化
func NewObject(dao *dao.Connect) *Object {
	return &Object{dao: dao}
}

// 新增数据
func (s *Object) AddOrder(add model.Order) error {

	order := model.Order{
		ID:       0,
		UserName: add.UserName,
		Amount:   add.Amount,
		Status:   add.Status,
		FileUrl:  add.FileUrl,
	}
	err := s.dao.CreateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

// 查询数据
func (s *Object) OrderDetail(id uint) (model.Order, error) {

	order, err := s.dao.SearchOrder(id)
	if err != nil {
		panic(err)
	}
	return order, nil
}

// 查询数据列表
func (s *Object) OrderList(order model.SearchList) ([]model.Order, error) {

	if order.Page <= 0 {
		order.Page = 1
	}
	if order.Limit <= 0 {
		order.Limit = 10
	}
	list, err := s.dao.SearchOrders("%"+order.UserName+"%", order.Page, order.Limit)

	if err != nil {
		return nil, err
	}
	return list, nil
}

// 更新数据
func (s *Object) UpdateOrder(order map[string]interface{}) error {
	err := s.dao.UpdateOrder(order)
	if err != nil {
		return err
	}
	return nil
}

// 更新文件路径
func (s *Object) UpdateFileUrl(id uint, url string) error {

	err := s.dao.UpdateUrl(id, url)
	if err != nil {
		return err
	}
	return nil
}

// 下载文件
func (s *Object) DownloadFile(id uint) (string, error) {

	order, err := s.dao.SearchOrder(id)
	if err != nil {
		panic(err)
	}
	url := order.FileUrl

	rs, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rs.Body)
	defer rs.Body.Close()
	if err != nil {
		panic(err)
	}
	// 文件命名
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("download-%d.jpg", timeStamp)
	// 存放文件路径
	filePath := fmt.Sprintf("new/%s", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		panic(err)
	}
	return filePath, nil
}
