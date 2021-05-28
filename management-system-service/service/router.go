package service

import (
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
	s.dao = dao.New()
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
	//s.router.Use(core.Handle(middleware.CurrentUser))

	//s.router.Use(core.Handle(middleware.Return))
	s.router.Use(core.Handle(middleware.Error))

	//建立路由表
	apiV1 := s.router.Group("api/v1")
	apiV1.POST("ping", core.Handle(s.Ping))
	{
		//用户数据
		userName := "user"
		apiV1.POST(userName, core.Handle(s.AddUser))
		apiV1.GET(userName+"/:id", core.Handle(s.GetUserBasis))
		apiV1.PUT(userName+"/:id", core.Handle(s.UpdateUserBasis))

		apiV1.GET("users/:id", core.Handle(s.GetUserList))
	}
	{
		//设备类型
		equipmentTypeName := "equipment_type"
		apiV1.POST(equipmentTypeName, core.Handle(s.AddEquipmentType))
		apiV1.GET(equipmentTypeName+"/:id", core.Handle(s.GetEquipmentTypeBasis))
		apiV1.PUT(equipmentTypeName+"/:id", core.Handle(s.UpdateEquipmentTypeBasis))

		apiV1.GET("equipment_types", core.Handle(s.GetEquipmentType))
	}
	{
		//设备
		equipmentName := "equipment"
		apiV1.POST(equipmentName, core.Handle(s.AddEquipment))
		apiV1.GET(equipmentName+"/:id", core.Handle(s.GetEquipment))
		apiV1.PUT(equipmentName+"/:id", core.Handle(s.UpdateEquipment))

		apiV1.GET("equipments/:id", core.Handle(s.GetEquipmentList))
	}
	{
		//权限组
		group := "group"
		apiV1.POST(group, core.Handle(s.AddGroup))
		apiV1.GET(group+"/:id", core.Handle(s.GetGroup))
		apiV1.PUT(group+"/:id", core.Handle(s.UpdateGroup))
		//权限组列表
		apiV1.GET("groups", core.Handle(s.GetGroupList))
	}
	{
		//维护信息
		maintain := "maintain"
		apiV1.POST(maintain, core.Handle(s.AddMaintain))
		apiV1.GET(maintain+"/:id", core.Handle(s.GetMaintain))
	}

	s.router.Run(":3000")
}
