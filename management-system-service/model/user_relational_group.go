package model

import "github.com/jinzhu/gorm"

//UserRelationalGroup 人员属于的权限组
type UserRelationalGroup struct {
	gorm.Model
	Id      int `gorm:"primary_key"`
	GroupId int `gorm:"not null"`
	UserId  int `gorm:"not null"`
}

//AddUserRelationalGroup 添加用户与权限组
func AddUserRelationalGroup(groupId int, userId int) (UserRelationalGroup, error) {
	urg := UserRelationalGroup{
		GroupId: groupId,
		UserId:  userId,
	}
	result := GetDB().Create(&urg)
	return urg, result.Error
}

//DeleteUserRelationalGroup 删除用户与权限组
func (urg *UserRelationalGroup) DeleteUserRelationalGroup() error {
	result := GetDB().Delete(urg)
	return result.Error
}
