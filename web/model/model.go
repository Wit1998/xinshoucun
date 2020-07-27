package model

// 数据库orders表
type Order struct {
	ID       uint    `json:"id"gorm:"type:varchar(100)"`
	UserName string  `json:"user_name"gorm:"type:varchar(100)"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"gorm:"type:varchar(100)"`
	FileUrl  string  `json:"file_url"gorm:"type:varchar(100)"`
}

//查询列表
type SearchList struct {
	UserName string `form:"user_name"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

// 更新列表
type NewUpdate map[string]interface{}
