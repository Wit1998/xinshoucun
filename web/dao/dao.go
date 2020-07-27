package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"xinshoucun/web/db"
	"xinshoucun/web/model"
)

type Connect struct {
	DbConnect *gorm.DB
}

var dbConnect *gorm.DB

func NewMyConnect() *Connect {
	if dbConnect == nil {
		dbConnect = db.NewDB(db.DbName)
		dbConnect.AutoMigrate(&model.Order{})
	}

	return &Connect{DbConnect: dbConnect}
}

// 创建接口
type ConnImpl interface {
	// 创建
	CreateOrder(order model.Order) error

	// 查询
	SearchOrder(id uint) (model.Order, error)

	// 查询所有
	SearchAll() ([]model.Order, error)

	// 查询列表
	SearchOrders(username string, page, limit int) ([]model.Order, error)

	// 更新
	UpdateOrder(order map[string]interface{}) error

	// 更新Url
	UpdateUrl(id uint, url string) error
}

// 实例化接口
// 创建
func (c *Connect) CreateOrder(order model.Order) error {
	//开始创建
	rs := c.DbConnect.Begin()
	result := rs.Create(order)
	if err := result.Error; err != nil {
		//创建失败，回滚
		rs.Rollback()
		return err
	}
	rs.Commit()
	return nil
}

// 查询
func (c *Connect) SearchOrder(id uint) (model.Order, error) {

	order := model.Order{
		ID:       0,
		UserName: "",
		Amount:   0,
		Status:   "",
		FileUrl:  "",
	}
	Search := c.DbConnect.Where("id = ?", id)
	if err := Search.First(order).Error; err != nil {
		fmt.Println("err", err)
	}
	return order, nil
}

// 查询所有
func (c *Connect) SearchAll() ([]model.Order, error) {

	var orders []model.Order
	if err := c.DbConnect.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 查询列表
func (c *Connect) SearchOrders(username string, page, limit int) ([]model.Order, error) {

	var orders []model.Order
	Search := c.DbConnect.Where("user_name LIKE ?", username).Order("created_at desc, amount")
	//发生偏移
	Search = Search.Offset((page - 1) * limit).Limit(limit)
	if err := Search.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 更新
func (c *Connect) UpdateOrder(order map[string]interface{}) error {
	//开始更新
	rs := c.DbConnect.Begin()
	if err := rs.Updates(order).Error; err != nil {
		// 如果找不到，则使用Rollback回滚，返回到上一级
		rs.Rollback()
		return err
	}
	// 找到则提交
	rs.Commit()
	return nil
}

//// 更新Url
//func (c *Connect) UpdateUrl(id uint, url string) error {
//
//	order := model.Order{
//		ID:       0,
//		UserName: "",
//		Amount:   0,
//		Status:   "",
//		FileUrl:  "",
//	}
//	// 开始更新
//	rs := c.DbConnect.Begin()
//	if err := rs.Model(&order).Where("id = ?", id).Update("file_url", url).Error; err != nil {
//		// 如果找不到，则使用Rollback回滚,返回到上一级
//		rs.Rollback()
//		return err
//	}
//	// 找到则提交
//	rs.Commit()
//	return nil
//}
