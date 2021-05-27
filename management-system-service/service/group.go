package service

import (
	"management-system-server/core"
	"management-system-server/model"
	log "management-system-server/util/logger"
	"net/http"
)

//AddGroup 添加权限组服务api逻辑处理
func (s *Service) AddGroup(c *core.Context) {
	info, err := s.dao.BindHTTPAddGroupInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	m, err := s.dao.AddGroup(c, info.Name)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPAddGroupResponse{
		ID: m.ID,
	}
	c.ReturnSuccess(response)
}

//GetGroup 获取权限组信息服务api逻辑处理
func (s *Service) GetGroup(c *core.Context) {
	m, err := s.dao.GetGroup(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPGetGroupResponse{
		Name: m.Name,
	}
	c.ReturnSuccess(response)
}

//GetGroupList 获取权限组列表服务api逻辑处理
func (s *Service) GetGroupList(c *core.Context) {
	m, err := s.dao.GetAllGroup(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	allGroup := make([]map[string]interface{}, len(m))
	for i := 0; i < len(m); i++ {
		allGroup[i] = map[string]interface{}{"id": m[i].ID, "name": m[i].Name}
	}

	response := model.HTTPGetGroupListResponse{
		Groups: allGroup,
	}
	c.ReturnSuccess(response)
}

//UpdateGroup 更新权限组服务api逻辑处理
func (s *Service) UpdateGroup(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateGroupInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	m, err := s.dao.GetGroup(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	m.Name = info.Name
	err = s.dao.SaveGroup(c, m)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	c.ReturnSuccess(true)
}
