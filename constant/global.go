package constant

import (
	"TheErrorCode/config"
	"gorm.io/gorm"
)

var (
	MYSQLCONFIG   *config.MySqlConfig
	DB            *gorm.DB
	JWTCONFIG     *config.Jwt
	DEFAULTIMAGES = "http://8.130.66.162:9000/douyin/default.jpg" //默认头像和背景
	MINIO         = "http://8.130.66.162:9000/douyin/"
)
