package model

import "github.com/jinzhu/gorm"

//EquipmentType 设备类型数据库模型
type EquipmentType struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Cycle uint64 `gorm:"not null"` //维护周期
}

//HttpEquipmentTypeInfo 设备类型http消息模型
type HttpEquipmentTypeInfo struct {
	ID    uint         `form:"id" json:"id" binding:"required"`
	Name  string       `form:"name" json:"name"`
	Cycle uint64       `form:"cycle" json:"cycle"`
	Tpye  HttpTpyeEnum `form:"type" json:"type" binding:"required"`
}
