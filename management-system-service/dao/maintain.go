package dao

import (
	"management-system-server/core"
	"management-system-server/model"
	"time"
)

//AddMaintain 添加维护表单
func (d *Dao) AddMaintain(c *core.Context,
	userID uint, equipmentID uint, date time.Time, status uint8, remark string) (m model.Maintain, err error) {

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

//BindHTTPAddMaintainInfo
func (d *Dao) BindHTTPAddMaintainInfo(c *core.Context) (info model.HTTPAddMaintainInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}

//BindHTTPGetMaintainInfo
func (d *Dao) BindHTTPGetMaintainInfo(c *core.Context) (info model.HTTPGetMaintainInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	return
}
