package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Name         string     `gorm:"not null"`
	Location     string     `gorm:"not null"` //地点
	Status       string     `gorm:"not null"` //状态
	Date         *time.Time `gorm:"not null"` //近期维护时间
	StartDate    *time.Time `gorm:"not null"` //开始服役时间
	Deadline     *time.Time `gorm:"not null"` //结束服役时间
	TypeId       uint        `gorm:"not null"` //设备类型id
	UserId       uint        `gorm:"not null"` //创建用户id
	CreationDate *time.Time `gorm:"not null"` //创建时间
}
