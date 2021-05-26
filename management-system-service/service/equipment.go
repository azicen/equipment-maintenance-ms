package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
	"time"
)

//AddEquipment AddEquipment服务api逻辑处理
func (s *Service) AddEquipment(c *core.Context) {
	info, err := s.dao.BindHTTPAddEquipmentInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}
	time.Unix(int64(info.StartDate), 0)
	m, err := s.dao.AddEquipment(
		c, info.Name, info.Location, info.Status,
		time.Unix(int64(info.StartDate), 0), time.Unix(int64(info.Deadline), 0),
		info.TypeID, info.UserID)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPAddEquipmentResponse{
		ID: m.ID,
	}
	c.SetData(response)
}
