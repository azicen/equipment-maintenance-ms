package model

import (
	"github.com/jinzhu/gorm"
)

//UserStatusEnum 账号状态枚举
type UserStatusEnum uint8

const (
	//UserStatusInactivu 未激活员工
	UserStatusInactivu UserStatusEnum = iota
	//Activu 激活员工
	UserStatusActive
	//UserStatusSuspend 被封禁员工，或离职
	UserStatusSuspend
)

//User 用户数据库模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Password string
	Status   uint8 `gorm:"not null"` //账号状态
}

//HTTPAddUserInfo HTTP消息模型AddUser
type HTTPAddUserInfo struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//HTTPAddUserInfo HTTP回复模型AddUser
type HTTPAddUserResponse struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetUserBasisResponse HTTP回复模型GetUserBasis
type HTTPGetUserBasisResponse struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Status uint8  `form:"status" json:"status" binding:"required"`
	Groups []uint `form:"groups" json:"groups" binding:"required"`
}
