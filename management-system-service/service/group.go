package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
)

//AddGroup AddGroup服务api逻辑处理
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

//GetGroupList GetGroupList服务api逻辑处理
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
