package dao

import (
	"bytes"
	"errors"
	"management-system-server/model"
	"time"

	"github.com/gin-gonic/gin"
)

//AddMaintain 添加维护表单
func (d *Dao) AddMaintain(c *gin.Context,
	userID uint, equipmentID uint, date *time.Time, status uint8, remark string) (err error) {

	m := model.Maintain{
		UserID:      userID,
		EquipmentID: equipmentID,
		Date:        date,
		Status:      status,
		Remark:      remark,
	}
	err = d.GetDB().Create(&m).Error
	return
}

//GetHttpMaintainInfo 获取维护表单http消息
func GetHttpMaintainInfo(c *gin.Context) (info model.HttpMaintainInfo, err error) {
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
		if info.UserID <= 0 {
			errB.WriteString("user_id, ")
		}
		if info.EquipmentID <= 0 {
			errB.WriteString("equipment_id, ")
		}
		if info.Date <= 0 {
			errB.WriteString("date, ")
		}
		if info.Status <= 0 {
			errB.WriteString("status, ")
		}
		if info.Remark == "" {
			errB.WriteString("remark, ")
		}
		break
	case model.HttpTpyeDelete:
		break
	case model.HttpTpyeChange:
		if info.UserID <= 0 && info.EquipmentID <= 0 && info.Date <= 0 && info.Status <= 0 && info.Remark == "" {
			errB.WriteString("user_id, equipment_id, date, status, remark, ")
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
