package model

import "github.com/jinzhu/gorm"

type EquipmentType struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Cycle uint    `gorm:"not null"` //维护周期
}
