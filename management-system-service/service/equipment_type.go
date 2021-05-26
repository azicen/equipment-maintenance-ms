package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
)

//AddEquipmentType AddEquipmentType服务api逻辑处理
func (s *Service) AddEquipmentType(c *core.Context) {
	info, err := s.dao.BindHTTPAddEquipmentTypeInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.AddEquipmentType(info.Name, info.Cycle)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPAddEquipmentTypeResponse{
		ID: m.ID,
	}
	c.SetData(response)
}
