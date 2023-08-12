package main

import initalize "TheErrorCode/initialize"

func main() {

	initalize.Viper()
	initalize.Mysql()
	initalize.Gin()
}
