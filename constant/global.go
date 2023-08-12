package constant

import (
	"TheErrorCode/config"
	"gorm.io/gorm"
)

var (
	MYSQLCONFIG *config.MySqlConfig
	DB          *gorm.DB
)
