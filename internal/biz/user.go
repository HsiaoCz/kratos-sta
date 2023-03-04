package biz

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Number   int64  `gorm:"column:number;type:int;"`
	Username string `gorm:"column:username;type:varchar(40);"`
	Password string `gorm:"column:password;type:varchar(100);"`
	Content  string `gorm:"column:content;type:varchar(200)"`
	Email    string `gorm:"column:email;type:varchar(40)"`
}

func (u User) TableName() string {
	return "user"
}
