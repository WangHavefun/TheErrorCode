package constant

import (
	"TheErrorCode/config"
	"gorm.io/gorm"
)

var (
	MYSQLCONFIG   *config.MySqlConfig
	DB            *gorm.DB
	JWTCONFIG     *config.Jwt
	DEFAULTIMAGES = "http://8.130.66.162:9090/api/v1/buckets/douyin/objects/download?preview=true&prefix=ZGVmYWx1dC5qcGc=&version_id=null" //默认头像和背景
)
