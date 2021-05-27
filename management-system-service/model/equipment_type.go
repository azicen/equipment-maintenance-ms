package model

import "github.com/jinzhu/gorm"

//EquipmentType 设备类型数据库模型
type EquipmentType struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Cycle uint64 `gorm:"not null"` //维护周期
}

//HTTPAddEquipmentTypeInfo HTTP消息模型AddEquipmentType
type HTTPAddEquipmentTypeInfo struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Cycle uint64 `form:"cycle" json:"cycle" binding:"required"`
}

//HTTPAddEquipmentTypeInfo HTTP回复模型AddEquipmentType
type HTTPAddEquipmentTypeResponse struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetEquipmentTypeBasisInfo 设备类型基础信息HTTP回复模型
type HTTPGetEquipmentTypeBasisResponse struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Cycle  uint64 `form:"cycle" json:"cycle" binding:"required"`
	Groups []uint `form:"groups" json:"groups" binding:"required"`
}

//HTTPUpdateEquipmentTypeBasisInfo 更新设备类型基础属性HTTP消息模型
type HTTPUpdateEquipmentTypeBasisInfo struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Cycle  uint64 `form:"cycle" json:"cycle" binding:"required"`
	Groups []uint `form:"groups" json:"groups" binding:"required"`
}
