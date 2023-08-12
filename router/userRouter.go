package router

type User struct {
	ID        int64  `gorm:"column:id" form:"id"`               //id主键
	UserName  string `gorm:"column:username" form:"username"`   //用户名
	Password  string `gorm:"column:password" form:"password"`   //密码
	Name      string `gorm:"column:name" form:"name"`           //昵称
	Signature string `gorm:"column:signature" form:"signature"` //简介
}
