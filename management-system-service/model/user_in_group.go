package model

import "github.com/jinzhu/gorm"

//UserInGroup 人员属于的权限组 数据库模型
type UserInGroup struct {
	gorm.Model
	GroupID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}
