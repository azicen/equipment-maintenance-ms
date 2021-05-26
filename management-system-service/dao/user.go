package dao

import (
	"management-system-server/core"
	"management-system-server/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	//PassWordCost 密码加密难度
	PassWordCost = 12
)

//AddUser 添加用户
func (d *Dao) AddUser(c *core.Context, name string, password string, status uint8) (m model.User, err error) {
	m = model.User{
		Name:     name,
		Password: password,
		Status:   status,
	}
	err = d.GetDB().Create(&m).Error
	return
}

//GetUser 用ID获取员工
func (d *Dao) GetUser(c *core.Context, id interface{}) (m model.User, err error) {
	err = d.GetDB().First(&m, id).Error
	return
}

//SetUserPassword 设置密码
func (d *Dao) SetUserPassword(c *core.Context, id interface{}, password string) (err error) {
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
func (d *Dao) CheckUserPassword(c *core.Context, id interface{}, password string) bool {
	u, err := d.GetUser(c, id)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//SetUserStatus 设置账号状态
func (d *Dao) SetUserStatus(c *core.Context, id interface{}, status uint8) (err error) {
	u, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	u.Status = status
	err = d.GetDB().Save(u).Error
	return
}

//DelUser 删除账号
func (d *Dao) DelUser(c *core.Context, id interface{}) (err error) {
	u, err := d.GetUser(c, id)
	if err != nil {
		return
	}
	err = d.GetDB().Save(u).Error
	return
}

//GetUserGroups 获取用户所在的权限组
func (d *Dao) GetUserGroups(c *core.Context, id interface{}) (m []model.UserInGroup, err error) {
	//// SELECT * FROM user_in_group WHERE user_id = id;
	err = d.GetDB().Where("user_id = ?", id).Find(&m).Error
	return
}

//BindHTTPAddUserInfo 
func (d *Dao) BindHTTPAddUserInfo(c *core.Context) (info model.HTTPAddUserInfo, err error) {
    err = c.BindJSON(&info)
    if err != nil {
        return
    }
    return
}
