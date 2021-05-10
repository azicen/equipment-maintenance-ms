package dao

import (
	"bytes"
	"errors"
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
func (d *Dao) AddUserInGroup(c *gin.Context, groupID uint, userID uint) (err error) {
	urg := model.UserInGroup{
		GroupID: groupID,
		UserID:  userID,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetUserInGroup 用ID获取用户与权限组的关系
func (d *Dao) GetUserInGroup(c *gin.Context, id interface{}) (m model.UserInGroup, err error) {
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
func (d *Dao) AddEquipmentInGroup(groupID uint, equipmentID uint) (err error) {
	urg := model.EquipmentInGroup{
		GroupID:     groupID,
		EquipmentID: equipmentID,
	}
	err = d.GetDB().Create(&urg).Error
	return
}

//GetEquipmentInGroup 用ID获取设备与权限组的关系
func (d *Dao) GetEquipmentInGroup(c *gin.Context, id interface{}) (m model.EquipmentInGroup, err error) {
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

//GetHttpGroupInfo 获取权限组http消息
func GetHttpGroupInfo(c *gin.Context) (info model.HttpGroupInfo, err error) {
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
	case model.HttpTpyeChange:
		if info.Name == "" {
			errB.WriteString("name, ")
		}
		break
	case model.HttpTpyeDelete:
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
