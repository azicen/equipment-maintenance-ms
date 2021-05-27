package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
)

//AddGroup 添加权限组服务api逻辑处理
func (s *Service) AddGroup(c *core.Context) {
	info, err := s.dao.BindHTTPAddGroupInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.AddGroup(c, info.Name)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPAddGroupResponse{
		ID: m.ID,
	}
	c.SetData(response)
}

//GetGroup 获取权限组信息服务api逻辑处理
func (s *Service) GetGroup(c *core.Context) {
	m, err := s.dao.GetGroup(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPGetGroupResponse{
		Name: m.Name,
	}
	c.SetData(response)
}

//GetGroupList 获取权限组列表服务api逻辑处理
func (s *Service) GetGroupList(c *core.Context) {
	m, err := s.dao.GetAllGroup(c)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	group_all := make(map[uint]string)
	for i := 0; i < len(m); i++ {
		group_all[m[i].ID] = m[i].Name
	}

	response := model.HTTPGetGroupListResponse{
		Groups: group_all,
	}
	c.SetData(response)
}

//UpdateGroup 更新权限组服务api逻辑处理
func (s *Service) UpdateGroup(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateGroupInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.GetGroup(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	m.Name = info.Name
	err = s.dao.SaveGroup(c, m)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	c.SetData(true)
}
