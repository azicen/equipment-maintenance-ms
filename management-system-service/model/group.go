package model

import "github.com/jinzhu/gorm"

//Group 权限组数据库模型
type Group struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

//HTTPAddGroupInfo 添加权限组HTTP消息模型
type HTTPAddGroupInfo struct {
	Name string `form:"name" json:"name" binding:"required"`
}

//HTTPAddGroupInfo 添加权限组HTTP回复模型
type HTTPAddGroupResponse struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetGroupListInfo 获取权限组列表HTTP回复模型
type HTTPGetGroupListResponse struct {
	Groups map[uint]string `form:"groups" json:"groups" binding:"required"`
}

//HTTPUpdateGroupInfo 更新权限组信息HTTP消息模型
type HTTPUpdateGroupInfo struct {
	Name string `form:"name" json:"name" binding:"required"`
}

//HTTPGetGroupInfo 获取权限组HTTP回复模型
type HTTPGetGroupResponse struct {
    Name string `form:"name" json:"name" binding:"required"`
}