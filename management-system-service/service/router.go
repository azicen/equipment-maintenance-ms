package service

import (
	"management-system-server/api"
	"management-system-server/core"
	"management-system-server/core/middleware"
	"management-system-server/dao"
	"os"

	"github.com/gin-gonic/gin"
)

type Service struct {
	dao    *dao.Dao
	router *gin.Engine
}

//New 初始化服务
func New() (s *Service) {
	s = &Service{}
	s.NewRouter()
	return
}

func (s *Service) GetRouter() *gin.Engine {
	return s.router
}

//NewRouter 初始化路由
func (s *Service) NewRouter() {
	s.router = gin.Default()

	//中间件, 顺序不能改
	s.router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	s.router.Use(middleware.InitCors())
	s.router.Use(core.Handle(middleware.CurrentUser))

	s.router.Use(core.Handle(middleware.Return))
	s.router.Use(core.Handle(middleware.Error))

	basisapi := s.router.Group("api")

	basisapi.POST("ping", core.Handle(api.Ping))
	//basisapi.POST("user", core.Handle(s.ChangeUser))

	s.router.Run(":3000")
}
