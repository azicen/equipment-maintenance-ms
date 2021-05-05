package service

import (
	"github.com/gin-gonic/gin"
	"management-system-server/api"
	"management-system-server/middleware"
	"os"
)

var Router *gin.Engine

//Init 初始化服务
func Init() {
	InitRouter()
}

func GetRouter() *gin.Engine {
	return Router
}

//InitRouter 初始化路由
func InitRouter() {
	router := gin.Default()

	//中间件, 顺序不能改
	router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.InitCors())
	router.Use(middleware.CurrentUser())

	basisapi := router.Group("api/basis")

	basisapi.POST("ping", api.Ping)

	router.Run(":3000")

	Router = router
}