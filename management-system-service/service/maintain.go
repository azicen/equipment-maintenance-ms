package service

import (
	"management-system-server/core"
	"management-system-server/model"
	log "management-system-server/util/logger"
	"net/http"
	"strconv"
	"time"
)

//AddMaintain AddMaintain服务api逻辑处理
func (s *Service) AddMaintain(c *core.Context) {
	info, err := s.dao.BindHTTPAddMaintainInfo(c)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	m, err := s.dao.AddMaintain(c, info.UserID, info.EquipmentID, time.Unix(int64(info.Date), 0), info.Status, info.Remark)
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPAddMaintainResponse{
		ID: m.ID,
	}
	c.ReturnSuccess(response)
}

//GetMaintain GetMaintain服务api逻辑处理
func (s *Service) GetMaintain(c *core.Context) {
	m, err := s.dao.GetMaintain(c, c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}

	response := model.HTTPGetMaintainResponse{
		UserID:      m.UserID,
		EquipmentID: m.EquipmentID,
		Date:        uint64(m.Date.Unix()),
		Status:      m.Status,
		Remark:      m.Remark,
	}
	c.ReturnSuccess(response)
}

//GetMaintainList 获取维护信息列表服务api逻辑处理
func (s *Service) GetMaintainList(c *core.Context) {
	nextId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(c.DefaultQuery("user_id", "0"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}

	var m []model.Maintain
	if userId == 0 {
		m, err = s.dao.GetTenMaintainList(c, uint(nextId))
		if err != nil {
			log.Error(err.Error(), err)
			c.SetCode(http.StatusInternalServerError)
			return
		}
	} else {
		m, err = s.dao.GetTenMaintainByUserList(c, uint(userId), uint(nextId))
		if err != nil {
			log.Error(err.Error(), err)
			c.SetCode(http.StatusInternalServerError)
			return
		}
	}

	allMaintain := make([]map[string]interface{}, len(m))
	for i := 0; i < len(m); i++ {
		allMaintain[i] = map[string]interface{}{
			"id":           m[i].ID,
			"user_id":      m[i].UserID,
			"equipment_id": m[i].EquipmentID,
			"date":         m[i].Date,
			"status":       m[i].Status,
			"remark":       m[i].Remark,
		}
	}

	response := model.HTTPGetMaintainListResponse{
		Maintains: allMaintain,
	}
	c.ReturnSuccess(response)
}

//DeleteMaintain 删除用户服务api逻辑
func (s *Service) DeleteMaintain(c *core.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusBadRequest)
		return
	}
	err = s.dao.DelMaintain(c, uint(id))
	if err != nil {
		log.Error(err.Error(), err)
		c.SetCode(http.StatusInternalServerError)
		return
	}
	c.ReturnSuccess(true)
}