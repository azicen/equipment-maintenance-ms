package dao

import (
	"management-system-server/model"

	"github.com/gin-gonic/gin"
)

//AddEquipmentType 添加设备类型
func (d *Dao) AddEquipmentType(name string, cycle uint) (err error) {
	et := model.EquipmentType{
		Name:  name,
		Cycle: cycle,
	}
	err = d.GetDB().Create(&et).Error
	return
}

//GetEquipmentType 用ID获取设备类型
func (d *Dao) GetEquipmentType(c *gin.Context, id interface{}) (et model.EquipmentType, err error) {
	err = d.GetDB().First(&et, id).Error
	return
}

//DelEquipmentType 删除设备类型
func (d *Dao) DelEquipmentType(c *gin.Context, id interface{}) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Create(&et).Error
	return
}

//SetCycle 设置设备维护周期
func (d *Dao) SetCycle(c *gin.Context, id interface{}, cycle uint) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	et.Cycle = cycle
	err = d.GetDB().Create(&et).Error
	return
}
