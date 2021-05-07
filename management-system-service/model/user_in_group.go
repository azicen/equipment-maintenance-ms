package model

import "github.com/jinzhu/gorm"

//UserRelationalGroup 人员属于的权限组
type UserRelationalGroup struct {
	gorm.Model
	GroupId uint `gorm:"not null"`
	UserId  uint `gorm:"not null"`
}
