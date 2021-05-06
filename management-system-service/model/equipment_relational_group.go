package model

import "github.com/jinzhu/gorm"

//EquipmentRelationalGroup 权限组拥有的设备类型
type EquipmentRelationalGroup struct {
	gorm.Model
	Id          int `gorm:"primary_key"`
	GroupId     int `gorm:"not null"`
	EquipmentId int `gorm:"not null"`
}

//AddEquipmentRelationalGroup 添加设备类型与权限组
func AddEquipmentRelationalGroup(groupId int, equipmentId int) (EquipmentRelationalGroup, error) {
	urg := EquipmentRelationalGroup{
		GroupId:     groupId,
		EquipmentId: equipmentId,
	}
	result := GetDB().Create(&urg)
	return urg, result.Error
}

//DeleteEquipmentRelationalGroup 删除设备类型与权限组
func (urg *EquipmentRelationalGroup) DeleteEquipmentRelationalGroup() error {
	result := GetDB().Delete(urg)
	return result.Error
}
