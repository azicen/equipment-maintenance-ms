package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Maintain struct {
	gorm.Model
	UserId      uint        `gorm:"not null"` //维护表单发起用户
	EquipmentId uint        `gorm:"not null"` //维护设备id
	Date        *time.Time `gorm:"not null"` //维护日期
	Status      string     `gorm:"not null"` //设备状态
	Remark      string     `gorm:"not null"` //异常信息
}
