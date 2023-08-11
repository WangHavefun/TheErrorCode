package model

type User struct {
	ID        int64  `gorm:"column:id"`        //id主键
	UserName  string `gorm:"column:username"`  //用户名
	Password  string `gorm:"column:password"`  //密码
	Name      string `gorm:"column:name"`      //昵称
	Signature string `gorm:"column:signature"` //简介
}

func (u *User) TableName() string {
	return "user"
}
