package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
)

//AddEquipmentType 添加设备类型服务api逻辑处理
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

//GetEquipmentTypeBasis 获取设备类型基础属性服务api逻辑处理
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
		Name:   m.Name,
		Cycle:  m.Cycle,
		Groups: groupIdArray,
	}
	c.SetData(response)
}

//UpdateEquipmentTypeBasis 更新设备类型基础属性服务api逻辑处理
func (s *Service) UpdateEquipmentTypeBasis(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateEquipmentTypeBasisInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	//基础信息更新
	m, err := s.dao.GetEquipmentType(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	m.Name = info.Name
	m.Cycle = info.Cycle
	err = s.dao.SaveEquipmentType(c, m)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	//权限组的更新
	oldGroups, err := s.dao.GetEquipmentTypeGroups(c, m.ID)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	newMap := make(map[uint]struct{})
	oldMap := make(map[uint]struct{})
	for _, groupId := range info.Groups {
		newMap[groupId] = struct{}{}
	}
	for _, group := range oldGroups {
		oldMap[group.GroupID] = struct{}{}
		if _, ok := newMap[group.GroupID]; !ok {
			err = s.dao.DelEquipmentInGroup(c, group.ID)
			if err != nil {
				c.Error(http.StatusInternalServerError, err)
			}
		}
	}
	for groupId := range newMap {
		if _, ok := oldMap[groupId]; !ok {
			err = s.dao.AddEquipmentInGroup(c, groupId, m.ID)
			if err != nil {
				c.Error(http.StatusInternalServerError, err)
			}
		}
	}

	c.SetData(true)
}
