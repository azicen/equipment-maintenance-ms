package service

import (
	"management-system-server/core"
	"management-system-server/model"
	log "management-system-server/util/logger"
	"net/http"
	"strconv"
)

//AddUser AddUser服务api逻辑处理
func (s *Service) AddUser(c *core.Context) {
	info, err := s.dao.BindHTTPAddUserInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	m, err := s.dao.AddUser(c, info.Name, info.Password, uint8(model.UserStatusInactivu))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPAddUserResponse{
		ID: m.ID,
	}
	c.ReturnSuccess(response)
}

//GetUserBasis 获取用户基础信息服务api逻辑处理
func (s *Service) GetUserBasis(c *core.Context) {
	m, err := s.dao.GetUser(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
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
	c.ReturnSuccess(response)
}

//UpdateUserBasis UpdateUserBasis服务api逻辑处理
func (s *Service) UpdateUserBasis(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateUserBasisInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	//基础信息更新
	m, err := s.dao.GetUser(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	m.Name = info.Name
	m.Status = info.Status
	err = s.dao.SaveUser(c, m)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	//权限组的更新
	oldGroups, err := s.dao.GetUserGroups(c, m.ID)
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
			err = s.dao.DelUserInGroup(c, group.ID)
			if err != nil {
				log.Error(err.Error(), err)
				c.SetCode(http.StatusInternalServerError)
			}
		}
	}
	for groupId := range newMap {
		if _, ok := oldMap[groupId]; !ok {
			err = s.dao.AddUserInGroup(c, groupId, m.ID)
			if err != nil {
				log.Error(err.Error(), err)
				c.SetCode(http.StatusInternalServerError)
			}
		}
	}

	c.ReturnSuccess(true)
}

//GetUserList 获取用户列表服务api逻辑处理
func (s *Service) GetUserList(c *core.Context) {
	nextId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	m, err := s.dao.GetTenUserList(c, uint(nextId))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	allUser := make([]map[string]interface{}, len(m))
	for i := 0; i < len(m); i++ {
		allUser[i] = map[string]interface{}{"id": m[i].ID, "name": m[i].Name}
	}

	response := model.HTTPGetUserListResponse{
		Users: allUser,
	}
	c.ReturnSuccess(response)
}

//DeleteUser 删除用户服务api逻辑
func (s *Service) DeleteUser(c *core.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	err = s.dao.DelUser(c, uint(id))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	c.ReturnSuccess(true)
}