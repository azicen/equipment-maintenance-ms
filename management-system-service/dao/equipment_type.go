package dao

import (
	"management-system-server/core"
	"management-system-server/model"
)

//AddEquipmentType 添加设备类型
func (d *Dao) AddEquipmentType(name string, cycle uint64) (et model.EquipmentType, err error) {
	et = model.EquipmentType{
		Name:  name,
		Cycle: cycle,
	}
	err = d.GetDB().Create(&et).Error
	return
}

//GetEquipmentType 用ID获取设备类型
func (d *Dao) GetEquipmentType(c *core.Context, id interface{}) (et model.EquipmentType, err error) {
	err = d.GetDB().First(&et, id).Error
	return
}

//DelEquipmentType 删除设备类型
func (d *Dao) DelEquipmentType(c *core.Context, id interface{}) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Create(&et).Error
	return
}

//SetCycle 设置设备维护周期
func (d *Dao) SetCycle(c *core.Context, id interface{}, cycle uint64) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	et.Cycle = cycle
	err = d.GetDB().Create(&et).Error
	return
}

//BindHTTPAddEquipmentTypeInfo
func (d *Dao) BindHTTPAddEquipmentTypeInfo(c *core.Context) (info model.HTTPAddEquipmentTypeInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
