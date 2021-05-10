package model

import "github.com/jinzhu/gorm"

//EquipmentInGroup 权限组拥有的设备类型 数据库模型
type EquipmentInGroup struct {
	gorm.Model
	GroupID     uint `gorm:"not null"`
	EquipmentID uint `gorm:"not null"`
}
