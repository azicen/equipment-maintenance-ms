package main

import (
	"management-system-server/conf"
	"management-system-server/util"
	"management-system-server/service"
)

func main() {
	//从配置文件读取配置
	conf.Init()

	util.InitLog()
	service.Init()
}
