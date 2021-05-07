package dao

import (
	"management-system-server/model"
	"time"

	"github.com/gin-gonic/gin"
)

//AddMaintain 添加维护表单
func (d *Dao) AddMaintain(c *gin.Context, userId uint, equipmentId uint, date *time.Time, status string, remark string) (err error) {
	m := model.Maintain{
		UserId:      userId,
		EquipmentId: equipmentId,
		Date:        date,
		Status:      status,
		Remark:      remark,
	}
	err = d.GetDB().Create(&m).Error
	return
}
