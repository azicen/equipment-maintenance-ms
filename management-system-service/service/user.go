package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
)

//AddUser AddUser服务api逻辑处理
func (s *Service) AddUser(c *core.Context) {
	info, err := s.dao.BindHTTPAddUserInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.AddUser(c, info.Name, info.Password, uint8(model.UserStatusInactivu))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPAddUserResponse{
		ID: m.ID,
	}
	c.SetData(response)
}

//GetUserBasis 获取用户基础信息服务api逻辑处理
func (s *Service) GetUserBasis(c *core.Context) {
	m, err := s.dao.GetUser(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	groups, err := s.dao.GetUserGroups(c, m.ID)
	groupIdArray := []uint{}
	for i := 0; i < len(groups); i++ {
		groupIdArray = append(groupIdArray, groups[i].GroupID)
	}

	response := model.HTTPGetUserBasisResponse{
		Name:   m.Name,
		Status: m.Status,
		Groups: groupIdArray,
	}
	c.SetData(response)
}
