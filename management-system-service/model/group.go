package model

import "github.com/jinzhu/gorm"

//Group 权限组数据库模型
type Group struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//HTTPAddGroupInfo HTTP消息模型AddGroup
type HTTPAddGroupInfo struct {
    Name string `form:"name" json:"name" binding:"required"`
}

//HTTPAddGroupInfo HTTP回复模型AddGroup
type HTTPAddGroupResponse struct {
    ID uint `form:"id" json:"id" binding:"required"`
}

//HTTPGetGroupListInfo HTTP消息模型GetGroupList
type HTTPGetGroupListInfo struct {
}

//HTTPGetGroupListInfo HTTP回复模型GetGroupList
type HTTPGetGroupListResponse struct {
    Groups map[uint]string `form:"groups" json:"groups" binding:"required"`
}