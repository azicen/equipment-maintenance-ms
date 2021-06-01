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

//SaveEquipmentType 保存设备类型
func (d *Dao) SaveEquipmentType(c *core.Context, e model.EquipmentType) (err error) {
	err = d.GetDB().Save(&e).Error
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
	err = d.GetDB().Delete(&et).Error
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

//GetAllEquipmentType 获取所有设备类型信息
func (d *Dao) GetAllEquipmentType(c *core.Context) (m []model.EquipmentType, err error) {
	err = d.GetDB().Find(&m).Error
	return
}


//GetEquipmentTypeGroups 获取设备类型所在的权限组
func (d *Dao) GetEquipmentTypeGroups(c *core.Context, id interface{}) (m []model.EquipmentInGroup, err error) {
	//// SELECT * FROM user_in_group WHERE user_id = id;
	err = d.GetDB().Where("equipment_id = ?", id).Find(&m).Error
	return
}

//AddEquipmentInGroup 添加设备类型与权限组
func (d *Dao) AddEquipmentInGroup(c *core.Context, groupID uint, equipmentID uint) (err error) {
	urg := model.EquipmentInGroup{
		GroupID:     groupID,
		EquipmentID: equipmentID,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetEquipmentInGroup 用ID获取设备与权限组的关系
func (d *Dao) GetEquipmentInGroup(c *core.Context, id interface{}) (m model.EquipmentInGroup, err error) {
	// TODO bug
	err = d.GetDB().First(&m, id).Error
	return
}

//DelEquipmentInGroup 删除设备类型与权限组
func (d *Dao) DelEquipmentInGroup(c *core.Context, id interface{}) (err error) {
	erg, err := d.GetUserInGroup(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(erg).Error
	return
}

//BindHTTPAddEquipmentTypeInfo 解析添加设备类型http消息
func (d *Dao) BindHTTPAddEquipmentTypeInfo(c *core.Context) (info model.HTTPAddEquipmentTypeInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}

//BindHTTPUpdateEquipmentTypeBasisInfo 解析更新设备类型基础属性http消息
func (d *Dao) BindHTTPUpdateEquipmentTypeBasisInfo(c *core.Context) (info model.HTTPUpdateEquipmentTypeBasisInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
