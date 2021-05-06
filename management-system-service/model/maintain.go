package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Maintain struct {
	gorm.Model
	Id          int64      `gorm:"primary_key"`
	UserId      int        `gorm:"not null"` //维护表单发起用户
	EquipmentId int        `gorm:"not null"` //维护设备id
	Date        *time.Time `gorm:"not null"` //维护日期
	Status      string     `gorm:"not null"` //设备状态
	Remark      string     `gorm:"not null"` //异常信息
}

//AddMaintain 添加维护表单
func AddMaintain(userId int, equipmentId int, date *time.Time, status string, remark string) (Maintain, error) {
	m := Maintain{
		UserId:      userId,
		EquipmentId: equipmentId,
		Date:        date,
		Status:      status,
		Remark:      remark,
	}
	result := GetDB().Create(&m)
	return m, result.Error
}
