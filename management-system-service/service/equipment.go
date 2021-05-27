package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
	"time"
)

//AddEquipment 添加设备服务api逻辑处理
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

//GetEquipment 获取设备信息服务api逻辑处理
func (s *Service) GetEquipment(c *core.Context) {
	m, err := s.dao.GetEquipment(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPGetEquipmentResponse{
		Name:         m.Name,
		Location:     m.Location,
		Status:       m.Status,
		Date:         uint64(m.Date.Unix()),
		StartDate:    uint64(m.StartDate.Unix()),
		Deadline:     uint64(m.Deadline.Unix()),
		TypeID:       m.TypeID,
		UserID:       m.UserID,
		CreationDate: uint64(m.CreationDate.Unix()),
	}
	c.SetData(response)
}

//UpdateEquipment 更新设备服务逻辑处理
func (s *Service) UpdateEquipment(c *core.Context) {
	info, err := s.dao.BindHTTPUpdateEquipmentInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.GetEquipment(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}
	m.Name = info.Name
	m.Location = info.Location
	m.Status = info.Status
	m.Date = time.Unix(int64(info.Date), 0)
	m.UserID = info.UserID

	err = s.dao.SaveEquipment(c, m)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	c.SetData(true)
}
