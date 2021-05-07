package model

import (
	"github.com/jinzhu/gorm"
)

//UserStatusEnum 账号状态枚举
type UserStatusEnum uint8

const (
	//Inactivu 未激活员工
	Inactivu UserStatusEnum = iota
	//Activu 激活员工
	Active
	//Suspend 被封禁员工，或离职
	Suspend
)

//User 员工数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Password string
	Status   uint8 `gorm:"not null"` //账号状态
}
