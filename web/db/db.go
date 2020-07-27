package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DbName = "NewSQL"

// 初始化数据库连接
func init() {
	//CREATE DATABASE IF NOT EXISTS " + gm.dbConfig.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
	dbConnect := NewDB("mysql")
	if err := dbConnect.Exec("CREATE DATABASE IF NOT EXISTS " + DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;").Error; err != nil {
		fmt.Println("create err:", err)
		return
	}
}

func NewDB(name string) *gorm.DB {
	con, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/"+name+"?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		fmt.Println("open err:", err.Error())
		return nil
	}
	return con
}
