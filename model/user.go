package model

import "gorm.io/gorm"

type User struct {
	ID              int64  `gorm:"column:id;"`                                   //id主键
	UserName        string `gorm:"column:username;NOT NULL;type:varchar(255)" `  //用户名
	Password        string `gorm:"column:password;NOT NULL;type:varchar(255)"`   //密码
	Name            string `gorm:"column:name;type:varchar(255);default:'ikun'"` //昵称
	Signature       string `gorm:"column:signature;type:varchar(255)"`           //简介
	Avatar          string `gorm:"column:avatar;"`                               //用户头像
	BackgroundImage string `gorm:"column:background_image"`                      //个人用户顶部大图
	gorm.Model
}

func (u *User) TableName() string {
	return "user"
}
