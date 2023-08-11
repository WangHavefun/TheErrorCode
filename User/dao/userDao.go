package dao

import (
	"User/constant"
	"User/model"
)

func AddUser(user *model.User) {
	constant.DB.Create(user)
}
