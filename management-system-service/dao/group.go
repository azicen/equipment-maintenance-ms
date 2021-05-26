package dao

import (
	"management-system-server/core"
	"management-system-server/model"
)

//AddGroup 添加权限组
func (d *Dao) AddGroup(c *core.Context, name string) (m model.Group, err error) {
	m = model.Group{
		Name: name,
	}
	err = d.GetDB().Create(&m).Error
	return
}

//GetGroup 用ID获取权限组
func (d *Dao) GetGroup(c *core.Context, id interface{}) (m model.Group, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelGroup 删除权限组
func (d *Dao) DelGroup(c *core.Context, id interface{}) (err error) {
	g, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(g).Error
	return
}

//AddUserInGroup 添加用户与权限组的关系
func (d *Dao) AddUserInGroup(c *core.Context, groupID uint, userID uint) (err error) {
	urg := model.UserInGroup{
		GroupID: groupID,
		UserID:  userID,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetUserInGroup 用ID获取用户与权限组的关系
func (d *Dao) GetUserInGroup(c *core.Context, id interface{}) (m model.UserInGroup, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelUserInGroup 删除用户与权限组的关系
func (d *Dao) DelUserInGroup(c *core.Context, id interface{}) (err error) {
	urg, err := d.GetUserInGroup(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(urg).Error
	return
}

//AddEquipmentInGroup 添加设备类型与权限组
func (d *Dao) AddEquipmentInGroup(groupID uint, equipmentID uint) (err error) {
	urg := model.EquipmentInGroup{
		GroupID:     groupID,
		EquipmentID: equipmentID,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetEquipmentInGroup 用ID获取设备与权限组的关系
func (d *Dao) GetEquipmentInGroup(c *core.Context, id interface{}) (m model.EquipmentInGroup, err error) {
	// TODO bug
	err = d.GetDB().First(&m, id).Error
	return
}

//DelEquipmentInGroup 删除设备类型与权限组
func (d *Dao) DelEquipmentInGroup(c *core.Context, id interface{}) (err error) {
	erg, err := d.GetUserInGroup(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(erg).Error
	return
}

//GetAllGroup 获取所有权限组信息
func (d *Dao) GetAllGroup(c *core.Context) (gs []model.Group, err error) {
	err = d.GetDB().Find(&gs).Error
	return
}

//BindHTTPAddGroupInfo
func (d *Dao) BindHTTPAddGroupInfo(c *core.Context) (info model.HTTPAddGroupInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
