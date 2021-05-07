package model

import "github.com/jinzhu/gorm"

//EquipmentRelationalGroup 权限组拥有的设备类型
type EquipmentRelationalGroup struct {
	gorm.Model
	GroupId     uint `gorm:"not null"`
	EquipmentId uint `gorm:"not null"`
}
