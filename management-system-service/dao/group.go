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

//SaveGroup 保存权限组
func (d *Dao) SaveGroup(c *core.Context, e model.Group) (err error) {
	err = d.GetDB().Save(&e).Error
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

//BindHTTPUpdateGroupInfo
func (d *Dao) BindHTTPUpdateGroupInfo(c *core.Context) (info model.HTTPUpdateGroupInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
