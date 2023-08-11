package initalize

import (
	"User/constant"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

func Viper() {
	viper.SetConfigType("yml")                 //设置配置文件类型
	viper.SetConfigFile("./config/config.yml") //设置配置文件路径
	err := viper.ReadInConfig()                //读取配置文件
	if err != nil {
		log.Panic("获取配置文件错误")
	}
	c := viper.AllSettings()                               //读取所有配置为map类型
	mapstructure.Decode(c["mysql"], &constant.MYSQLCONFIG) //反序列化
	viper.WatchConfig()
}
