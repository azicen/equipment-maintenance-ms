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

//User 用户数据库模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Password string
	Status   uint8 `gorm:"not null"` //账号状态
}

//HttpUserInfo 用户http消息模型
type HttpUserInfo struct {
	ID       uint         `form:"id" json:"id" binding:"required"`
	Name     string       `form:"name" json:"name"`
	GroupID  []uint       `form:"group_id" json:"group_id"`
	Password string       `form:"password" json:"password"`
	Status   uint8        `form:"status" json:"status" binding:"required"`
	Tpye     HttpTpyeEnum `form:"type" json:"type" binding:"required"`
}
