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

//GetEquipmentTypeBasis GetEquipmentTypeBasis服务api逻辑处理
func (s *Service) GetEquipmentTypeBasis(c *core.Context) {
	m, err := s.dao.GetEquipmentType(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	groups, err := s.dao.GetEquipmentTypeGroups(c, c.Param("id"))
	groupIdArray := []uint{}
	for i := 0; i < len(groups); i++ {
		groupIdArray = append(groupIdArray, groups[i].GroupID)
	}

    response := model.HTTPGetEquipmentTypeBasisResponse{
        Name:  m.Name,
        Cycle: m.Cycle ,
        Groups: groupIdArray ,
    }
    c.SetData(response)
}