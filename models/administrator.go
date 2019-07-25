package models

type Administrator struct {
	ID       int    `json:"ID" gorm:"column:id;primary_key"`
	HallID   int    `json:"hallID" gorm:"column:hallID;"`
	SalesID  string `json:"salesID" gorm:"column:salesID;"`
	Account  string `json:"account" gorm:"column:account;"`
	Password string `json:"password" gorm:"column:password;"`
	Name     string `json:"name" gorm:"column:name;"`
}

func (Administrator) TableName() string {
	return "administrator_list"
}
