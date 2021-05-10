package dao

import (
	"bytes"
	"errors"
	"management-system-server/model"
	"time"

	"github.com/gin-gonic/gin"
)

//AddEquipment 添加设备
func (d *Dao) AddEquipment(c *gin.Context,
	name string, location string, status uint8,
	startDate *time.Time, deadline *time.Time,
	typeID uint, userID uint) (err error) {

	creationDate := time.Now()
	e := model.Equipment{
		Name:         name,
		Location:     location,
		Status:       status,
		Date:         &creationDate,
		StartDate:    startDate,
		Deadline:     deadline,
		TypeID:       typeID,
		UserID:       userID,
		CreationDate: &creationDate,
	}
	err = d.GetDB().Create(&e).Error
	return
}

//GetEquipment 用ID获取设备
func (d *Dao) GetEquipment(c *gin.Context, id interface{}) (e model.Equipment, err error) {
	err = d.GetDB().First(&e, id).Error
	return
}

//DelEquipment 删除设备
func (d *Dao) DelEquipment(c *gin.Context, id interface{}) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Save(e).Error
	return
}

//SetDate 设置维护日期
func (d *Dao) SetDate(c *gin.Context, id interface{}, date *time.Time) (err error) {
	e, err := d.GetEquipment(c, id)
	if err != nil {
		return
	}
	e.Date = date
	err = d.GetDB().Save(e).Error
	return
}

//GetHttpEquipmentInfo 获取设备http消息
func (d *Dao) GetHttpEquipmentInfo(c *gin.Context) (info model.HttpEquipmentInfo, err error) {
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
		if info.Location == "" {
			errB.WriteString("location, ")
		}
		if info.Status <= 0 {
			errB.WriteString("status, ")
		}
		if info.Date <= 0 {
			errB.WriteString("date, ")
		}
		if info.StartDate <= 0 {
			errB.WriteString("start_date, ")
		}
		if info.Deadline <= 0 {
			errB.WriteString("deadline, ")
		}
		if info.TypeID <= 0 {
			errB.WriteString("type_id, ")
		}
		if info.UserID <= 0 {
			errB.WriteString("user_id, ")
		}
		break
	case model.HttpTpyeDelete:
		break
	case model.HttpTpyeChange:
		if info.Name == "" && info.Location == "" && info.Status <= 0 &&
			info.Date <= 0 && info.StartDate <= 0 && info.Deadline <= 0 &&
			info.TypeID <= 0 && info.UserID <= 0 && len(info.GroupID) == 0 {

			errB.WriteString("name, location, status, date, start_date, deadline, type_id, user_id, group_id, ")
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
