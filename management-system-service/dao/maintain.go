package dao

import (
	"errors"
	"management-system-server/core"
	"management-system-server/model"
	"time"
)

//AddMaintain 添加维护表单
func (d *Dao) AddMaintain(c *core.Context,
	userID uint, equipmentID uint, date time.Time, status string, remark string) (m model.Maintain, err error) {

	m = model.Maintain{
		UserID:      userID,
		EquipmentID: equipmentID,
		Date:        date,
		Status:      status,
		Remark:      remark,
	}
	err = d.GetDB().Create(&m).Error
	return
}

//GetMaintain 通过id获取维护表单
func (d *Dao) GetMaintain(c *core.Context, id interface{}) (m model.Maintain, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//DelMaintain 通过id删除维护表单
func (d *Dao) DelMaintain(c *core.Context, id interface{}) (err error) {
	u, err := d.GetMaintain(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Delete(u).Error
	return
}

//GetTenMaintainList 获取10个维护信息，从id为"nextId"开始
func (d *Dao) GetTenMaintainList(c *core.Context, nextId interface{}) (m []model.Maintain, err error) {
	if _, ok := nextId.(uint); !ok {
		err = errors.New("nextId not uint")
		return
	}
	err = d.GetDB().Limit(10).Where("id >= ?", nextId).Find(&m).Error
	return
}

//GetTenMaintainByUserList 获取10个指定用户的维护信息，从id为"nextId"开始
func (d *Dao) GetTenMaintainByUserList(c *core.Context, userId interface{}, nextId interface{}) (m []model.Maintain, err error) {
	if _, ok := userId.(uint); !ok {
		err = errors.New("userId not uint")
		return
	}
	if _, ok := nextId.(uint); !ok {
		err = errors.New("nextId not uint")
		return
	}
	err = d.GetDB().Limit(10).Where("user_id = ? AND id >= ?", userId, nextId).Find(&m).Error
	return
}

//BindHTTPAddMaintainInfo
func (d *Dao) BindHTTPAddMaintainInfo(c *core.Context) (info model.HTTPAddMaintainInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
