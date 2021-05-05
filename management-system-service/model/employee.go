package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Employee 员工数据模型
type Employee struct {
	gorm.Model
	Id       string
	Name     string
	GroupId  int
	Password string
}

const (
	//PassWordCost 密码加密难度
	PassWordCost = 12
	//Active 激活员工
	Active string = "active"
	//Inactive 未激活员工
	Inactive string = "inactive"
	//Suspend 被封禁员工，或离职
	Suspend string = "suspend"
)

//GetEmployee 用ID获取员工
func GetEmployee(ID interface{}) (Employee, error) {
	var e Employee
	result := GetDB().First(&e, ID)
	return e, result.Error
}

//SetPassword 设置密码
func (e *Employee) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	e.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (e *Employee) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err == nil
}