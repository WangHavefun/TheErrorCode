package initalize

import (
	"TheErrorCode/constant"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	host := constant.MYSQLCONFIG.Host
	password := constant.MYSQLCONFIG.Password
	username := constant.MYSQLCONFIG.Username
	port := constant.MYSQLCONFIG.Port
	dbName := constant.MYSQLCONFIG.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	constant.DB = db
}
