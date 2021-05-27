package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Equipment 设备数据库模型
type Equipment struct {
	gorm.Model
	Name         string    `gorm:"not null"`
	Location     string    `gorm:"not null"` //地点
	Status       uint8     `gorm:"not null"` //状态
	Date         time.Time `gorm:"not null"` //近期维护时间
	StartDate    time.Time `gorm:"not null"` //开始服役时间
	Deadline     time.Time `gorm:"not null"` //结束服役时间
	TypeID       uint      `gorm:"not null"` //设备类型id
	UserID       uint      `gorm:"not null"` //创建用户id
	CreationDate time.Time `gorm:"not null"` //创建时间
}

//HTTPAddEquipmentInfo 添加设备HTTP消息模型
type HTTPAddEquipmentInfo struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Location     string `form:"location" json:"location" binding:"required"`
	Status       uint8  `form:"status" json:"status" binding:"required"`
	Date         uint64 `form:"date" json:"date" binding:"required"`
	StartDate    uint64 `form:"start_date" json:"start_date" binding:"required"`
	Deadline     uint64 `form:"deadline" json:"deadline" binding:"required"`
	TypeID       uint   `form:"type_id" json:"type_id" binding:"required"`
	UserID       uint   `form:"user_id" json:"user_id" binding:"required"`
	CreationDate uint64 `form:"creation_date" json:"creation_date" binding:"required"`
}

//HTTPAddEquipmentInfo 添加设备HTTP回复模型
type HTTPAddEquipmentResponse struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetEquipmentInfo 获取设备信息HTTP回复模型
type HTTPGetEquipmentResponse struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Location     string `form:"location" json:"location" binding:"required"`
	Status       uint8  `form:"status" json:"status" binding:"required"`
	Date         uint64 `form:"date" json:"date" binding:"required"`
	StartDate    uint64 `form:"start_date" json:"start_date" binding:"required"`
	Deadline     uint64 `form:"deadline" json:"deadline" binding:"required"`
	TypeID       uint   `form:"type_id" json:"type_id" binding:"required"`
	UserID       uint   `form:"user_id" json:"user_id" binding:"required"`
	CreationDate uint64 `form:"creation_date" json:"creation_date" binding:"required"`
}

//HTTPUpdateEquipmentInfo 更新设备HTTP消息模型
type HTTPUpdateEquipmentInfo struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Location string `form:"location" json:"location" binding:"required"`
	Status   uint8  `form:"status" json:"status" binding:"required"`
	Date     uint64 `form:"date" json:"date" binding:"required"`
	UserID   uint   `form:"user_id" json:"user_id" binding:"required"`
}
