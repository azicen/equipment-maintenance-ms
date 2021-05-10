package model

import "github.com/jinzhu/gorm"

//Group 权限组数据库模型
type Group struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//HttpGroupInfo 权限组http消息模型
type HttpGroupInfo struct {
	ID   uint         `form:"id" json:"id" binding:"required"`
	Name string       `form:"name" json:"name"`
	Tpye HttpTpyeEnum `form:"type" json:"type" binding:"required"`
}
