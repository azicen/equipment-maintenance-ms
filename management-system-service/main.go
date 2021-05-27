package main

import (
	"management-system-server/conf"
	"management-system-server/service"
	log "management-system-server/util/logger"
)

func main() {
	//从配置文件读取配置
	conf.Init()

	log.Init()
	
	service.New()
}
