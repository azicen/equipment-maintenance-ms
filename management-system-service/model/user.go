package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//UserStatusEnum 账号状态枚举
type UserStatusEnum int8

const (
	//Inactivu 未激活员工
	Inactivu UserStatusEnum = iota
	//Activu 激活员工
	Active
	//Suspend 被封禁员工，或离职
	Suspend
)

//User 员工数据模型
type User struct {
	gorm.Model
	Id       int    `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Password string
	Status   int8 `gorm:"not null"` //账号状态
}

const (
	//PassWordCost 密码加密难度
	PassWordCost = 12
)

//AddUser 添加用户
func AddUser(name string, password string, status int8) (User, error) {
	u := User{
		Name:     name,
		Password: password,
		Status:   status,
	}
	result := GetDB().Create(&u)
	return u, result.Error
}

//GetUser 用ID获取员工
func GetUser(id interface{}) (User, error) {
	var u User
	result := GetDB().First(&u, id)
	return u, result.Error
}

//SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	result := GetDB().Save(u)
	return result.Error
}

//CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

//SetStatus 设置账号状态
func (u *User) SetStatus(status UserStatusEnum) error {
	u.Status = int8(status)
	result := GetDB().Save(u)
	return result.Error
}

//DeleteUser 删除账号
func (u *User) DeleteUser() error {
	result := GetDB().Delete(u)
	return result.Error
}
