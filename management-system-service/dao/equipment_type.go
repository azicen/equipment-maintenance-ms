package dao

import (
	"bytes"
	"errors"
	"management-system-server/model"

	"github.com/gin-gonic/gin"
)

//AddEquipmentType 添加设备类型
func (d *Dao) AddEquipmentType(name string, cycle uint64) (err error) {
	et := model.EquipmentType{
		Name:  name,
		Cycle: cycle,
	}
	err = d.GetDB().Create(&et).Error
	return
}

//GetEquipmentType 用ID获取设备类型
func (d *Dao) GetEquipmentType(c *gin.Context, id interface{}) (et model.EquipmentType, err error) {
	err = d.GetDB().First(&et, id).Error
	return
}

//DelEquipmentType 删除设备类型
func (d *Dao) DelEquipmentType(c *gin.Context, id interface{}) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Create(&et).Error
	return
}

//SetCycle 设置设备维护周期
func (d *Dao) SetCycle(c *gin.Context, id interface{}, cycle uint64) (err error) {
	et, err := d.GetEquipmentType(c, id)
	if err != nil {
		return
	}
	et.Cycle = cycle
	err = d.GetDB().Create(&et).Error
	return
}

//GetHttpEquipmentTypeInfo 获取设备类型http消息
func (d *Dao) GetHttpEquipmentTypeInfo(c *gin.Context) (info model.HttpEquipmentTypeInfo, err error) {
	err = c.BindJSON(&info)
	if err != nil {
		return
	}
	errB := new(bytes.Buffer)
	if info.ID <= 0 {
		errB.WriteString("id, ")
	}
	switch info.Tpye {
	case model.HttpTpyeAdd:
		if info.Name == "" {
			errB.WriteString("name, ")
		}
		if info.Cycle <= 0 {
			errB.WriteString("cycle, ")
		}
		break
	case model.HttpTpyeDelete:
		break
	case model.HttpTpyeChange:
		if info.Name == "" && info.Cycle <= 0 {
			errB.WriteString("name, cycle, ")
		}
		break
	default:
		errB.WriteString("type, ")
		break
	}
	if errB.Len() > 0 {
		errB.WriteString("数据非法。")
		err = errors.New(errB.String())
	}
	return
}
