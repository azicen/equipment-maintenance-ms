package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Equipment 设备数据库模型
type Equipment struct {
	gorm.Model
	Name         string     `gorm:"not null"`
	Location     string     `gorm:"not null"` //地点
	Status       uint8      `gorm:"not null"` //状态
	Date         *time.Time `gorm:"not null"` //近期维护时间
	StartDate    *time.Time `gorm:"not null"` //开始服役时间
	Deadline     *time.Time `gorm:"not null"` //结束服役时间
	TypeID       uint       `gorm:"not null"` //设备类型id
	UserID       uint       `gorm:"not null"` //创建用户id
	CreationDate *time.Time `gorm:"not null"` //创建时间
}

//HttpEquipmentInfo 设备http消息模型
type HttpEquipmentInfo struct {
	ID        uint         `form:"id" json:"id" binding:"required"`
	Name      string       `form:"name" json:"name"`
	GroupID   []uint       `form:"group_id" json:"group_id"`
	Location  string       `form:"location" json:"location"`
	Status    uint8        `form:"status" json:"status"`
	Date      uint64       `form:"date" json:"date"`
	StartDate uint64       `form:"start_date" json:"start_date"`
	Deadline  uint64       `form:"deadline" json:"deadline"`
	TypeID    uint         `form:"type_id" json:"type_id"`
	UserID    uint         `form:"user_id" json:"user_id"`
	Tpye      HttpTpyeEnum `form:"type" json:"type" binding:"required"`
}
