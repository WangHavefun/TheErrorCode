package constant

import (
	"User/config"
	"gorm.io/gorm"
)

var (
	MYSQLCONFIG *config.MySqlConfig
	DB          *gorm.DB
)
