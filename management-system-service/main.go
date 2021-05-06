package main

import (
	"management-system-server/conf"
	"management-system-server/model"
	"management-system-server/service"
	"management-system-server/util"
)

func main() {
	//从配置文件读取配置
	conf.Init()

	util.InitLog()
	model.InitDatabase()
	service.Init()
}
