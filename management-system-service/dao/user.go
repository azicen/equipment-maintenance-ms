package dao

import (
	"bytes"
	"errors"
	"management-system-server/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	//PassWordCost 密码加密难度
	PassWordCost = 12
)

//AddUser 添加用户
func (d *Dao) AddUser(c *gin.Context, name string, password string, status uint8) (err error) {
	u := model.User{
		Name:     name,
		Password: password,
		Status:   status,
	}
	err = d.GetDB().Create(&u).Error
	return
}

//GetUser 用ID获取员工
func (d *Dao) GetUser(c *gin.Context, id interface{}) (m model.User, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//SetUserPassword 设置密码
func (d *Dao) SetUserPassword(c *gin.Context, id interface{}, password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return
	}
	u, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	u.Password = string(bytes)
	err = d.GetDB().Save(u).Error
	return
}

//CheckUserPassword 校验密码
func (d *Dao) CheckUserPassword(c *gin.Context, id interface{}, password string) bool {
	u, err := d.GetUser(c, id)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//SetUserStatus 设置账号状态
func (d *Dao) SetUserStatus(c *gin.Context, id interface{}, status uint8) (err error) {
	u, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	u.Status = status
	err = d.GetDB().Save(u).Error
	return
}

//DelUser 删除账号
func (d *Dao) DelUser(c *gin.Context, id interface{}) (err error) {
	u, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Save(u).Error
	return
}

//GetHttpUserInfo 获取用户http消息
func (d *Dao) GetHttpUserInfo(c *gin.Context) (info model.HttpUserInfo, err error) {
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
		if info.Password == "" {
			errB.WriteString("password, ")
		}
		if info.Status <= 0 {
			errB.WriteString("status, ")
		}
		break
	case model.HttpTpyeDelete:
		break
	case model.HttpTpyeChange:
		if info.Name == "" && info.Password == "" && info.Status <= 0 && len(info.GroupID) == 0 {
			errB.WriteString("name, password, status, group_id, ")
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
