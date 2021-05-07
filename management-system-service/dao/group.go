package dao

import (
	"management-system-server/model"

	"github.com/gin-gonic/gin"
)

//AddGroup 添加权限组
func (d *Dao) AddGroup(c *gin.Context, name string) (err error) {
	g := model.Group{
		Name: name,
	}
	err = d.GetDB().Create(&g).Error
	return
}

//GetGroup 用ID获取权限组
func (d *Dao) GetGroup(c *gin.Context, id interface{}) (m model.Group, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelGroup 删除权限组
func (d *Dao) DelGroup(c *gin.Context, id interface{}) (err error) {
	g, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(g).Error
	return
}


//AddUserInGroup 添加用户与权限组的关系
func (d *Dao) AddUserInGroup(c *gin.Context, groupId uint, userId uint) (err error) {
	urg := model.UserRelationalGroup{
		GroupId: groupId,
		UserId:  userId,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetUserInGroup 用ID获取用户与权限组的关系
func (d *Dao) GetUserInGroup(c *gin.Context, id interface{}) (m model.UserRelationalGroup, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelUserInGroup 删除用户与权限组的关系
func (d *Dao) DelUserInGroup(c *gin.Context, id interface{}) (err error) {
	urg, err := d.GetUserInGroup(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(urg).Error
	return
}


//AddEquipmentInGroup 添加设备类型与权限组
func (d *Dao) AddEquipmentInGroup(groupId uint, equipmentId uint) (err error) {
	urg := model.EquipmentRelationalGroup{
		GroupId:     groupId,
		EquipmentId: equipmentId,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetEquipmentInGroup 用ID获取设备与权限组的关系
func (d *Dao) GetEquipmentInGroup(c *gin.Context, id interface{}) (m model.EquipmentRelationalGroup, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelEquipmentInGroup 删除设备类型与权限组
func (d *Dao) DelEquipmentInGroup(c *gin.Context, id interface{}) (err error) {
	erg, err := d.GetUserInGroup(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(erg).Error
	return 
}
