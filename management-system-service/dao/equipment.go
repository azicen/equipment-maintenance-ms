package dao

import (
	"management-system-server/model"
	"time"

	"github.com/gin-gonic/gin"
)

//AddEquipment 添加设备
func (d *Dao) AddEquipment(c *gin.Context,
	name string, location string, status string,
	startDate *time.Time, deadline *time.Time,
	typeId uint, userId uint) (err error) {

	creationDate := time.Now()
	e := model.Equipment{
		Name:         name,
		Location:     location,
		Status:       status,
		Date:         &creationDate,
		StartDate:    startDate,
		Deadline:     deadline,
		TypeId:       typeId,
		UserId:       userId,
		CreationDate: &creationDate,
	}
	err = d.GetDB().Create(&e).Error
	return
}

//GetEquipment 用ID获取设备
func (d *Dao) GetEquipment(c *gin.Context, id interface{}) (e model.Equipment, err error) {
	err = d.GetDB().First(&e, id).Error
	return
}

//DelEquipment 删除设备
func (d *Dao) DelEquipment(c *gin.Context, id interface{}) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Save(e).Error
	return
}

//SetDate 设置维护日期
func (d *Dao) SetDate(c *gin.Context, id interface{}, date *time.Time) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	e.Date = date
	err = d.GetDB().Save(e).Error
	return
}
