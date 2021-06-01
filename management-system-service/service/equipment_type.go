package service

import (
	"management-system-server/core"
	"management-system-server/model"
	log "management-system-server/util/logger"
	"net/http"
	"strconv"
)

//AddEquipmentType 添加设备类型服务api逻辑处理
func (s *Service) AddEquipmentType(c *core.Context) {
	info, err := s.dao.BindHTTPAddEquipmentTypeInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	m, err := s.dao.AddEquipmentType(info.Name, info.Cycle)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPAddEquipmentTypeResponse{
		ID: m.ID,
	}
	c.ReturnSuccess(response)
}

//GetEquipmentTypeBasis 获取设备类型基础属性服务api逻辑处理
func (s *Service) GetEquipmentTypeBasis(c *core.Context) {
	m, err := s.dao.GetEquipmentType(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
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
	c.ReturnSuccess(response)
}

//UpdateEquipmentTypeBasis 更新设备类型基础属性服务api逻辑处理
func (s *Service) UpdateEquipmentTypeBasis(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateEquipmentTypeBasisInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	//基础信息更新
	m, err := s.dao.GetEquipmentType(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	m.Name = info.Name
	m.Cycle = info.Cycle
	err = s.dao.SaveEquipmentType(c, m)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	//权限组的更新
	oldGroups, err := s.dao.GetEquipmentTypeGroups(c, m.ID)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
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
				log.Error(err.Error(), err)
				c.SetCode(http.StatusInternalServerError)
			}
		}
	}
	for groupId := range newMap {
		if _, ok := oldMap[groupId]; !ok {
			err = s.dao.AddEquipmentInGroup(c, groupId, m.ID)
			if err != nil {
				log.Error(err.Error(), err)
				c.SetCode(http.StatusInternalServerError)
			}
		}
	}

	c.ReturnSuccess(true)
}

//GetEquipmentType 获取设备类型列表服务api逻辑处理
func (s *Service) GetEquipmentType(c *core.Context) {
	m, err := s.dao.GetAllEquipmentType(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	allEquipmentType := make([]map[string]interface{}, len(m))
	for i := 0; i < len(m); i++ {
		allEquipmentType[i] = map[string]interface{}{"id": m[i].ID, "name": m[i].Name}
	}

	response := model.HTTPGetEquipmentTypeListResponse{
		EquipmentTypes: allEquipmentType,
	}
	c.ReturnSuccess(response)
}

//DeleteEquipmentType 删除设备服务api逻辑
func (s *Service) DeleteEquipmentType(c *core.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	err = s.dao.DelEquipmentType(c, uint(id))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	c.ReturnSuccess(true)
}