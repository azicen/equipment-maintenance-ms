package dao

import (
	"management-system-server/core"
	"management-system-server/model"
	"time"
)

//AddEquipment 添加设备
func (d *Dao) AddEquipment(c *core.Context,
	name string, location string, status string,
	startDate time.Time, deadline time.Time,
	typeID uint, userID uint) (e model.Equipment, err error) {

	creationDate := time.Now()
	e = model.Equipment{
		Name:         name,
		Location:     location,
		Status:       status,
		Date:         creationDate,
		StartDate:    startDate,
		Deadline:     deadline,
		TypeID:       typeID,
		UserID:       userID,
	}
	err = d.GetDB().Create(&e).Error
	return
}

//SaveEquipment 保存设备信息
func (d *Dao) SaveEquipment(c *core.Context, e model.Equipment) (err error) {
	err = d.GetDB().Save(&e).Error
	return
}

//GetEquipment 用ID获取设备
func (d *Dao) GetEquipment(c *core.Context, id interface{}) (e model.Equipment, err error) {
	err = d.GetDB().First(&e, id).Error
	return
}

//DelEquipment 删除设备
func (d *Dao) DelEquipment(c *core.Context, id interface{}) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(e).Error
	return
}

//SetDate 设置维护日期
func (d *Dao) SetDate(c *core.Context, id interface{}, date time.Time) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	e.Date = date
	err = d.GetDB().Save(e).Error
	return
}

//BindHTTPAddEquipmentInfo
func (d *Dao) BindHTTPAddEquipmentInfo(c *core.Context) (info model.HTTPAddEquipmentInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}

//BindHTTPUpdateEquipmentInfo 解析更新设备HTTP消息模型
func (d *Dao) BindHTTPUpdateEquipmentInfo(c *core.Context) (info model.HTTPUpdateEquipmentInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
