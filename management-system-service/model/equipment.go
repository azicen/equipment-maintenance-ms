package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Id           int        `gorm:"primary_key"`
	Name         string     `gorm:"not null"`
	Location     string     `gorm:"not null"` //地点
	Status       string     `gorm:"not null"` //状态
	Date         *time.Time `gorm:"not null"` //近期维护时间
	StartDate    *time.Time `gorm:"not null"` //开始服役时间
	Deadline     *time.Time `gorm:"not null"` //结束服役时间
	TypeId       int        `gorm:"not null"` //设备类型id
	UserId       int        `gorm:"not null"` //创建用户id
	CreationDate *time.Time `gorm:"not null"` //创建时间
}

//GetEquipment 用ID获取设备
func GetEquipment(id interface{}) (Equipment, error) {
	var e Equipment
	result := GetDB().First(&e, id)
	return e, result.Error
}

//AddEquipment 添加设备
func AddEquipment(
	name string, location string, status string,
	startDate *time.Time, deadline *time.Time,
	typeId int, userId int) (Equipment, error) {

	creationDate := time.Now()
	e := Equipment{
		Name:         name,
		Location:     location,
		Status:       status,
		Date:         &creationDate,
		StartDate:    startDate,
		Deadline:     deadline,
		TypeId:       typeId,
		UserId:       userId,
		CreationDate: &creationDate,
	}
	result := GetDB().Create(&e)
	return e, result.Error
}

//DeleteEquipment 删除设备
func (e *Equipment) DeleteEquipment() error {
	result := GetDB().Delete(e)
	return result.Error
}

//SetDate 设置维护日期
func (e *Equipment) SetDate(date *time.Time) error {
	e.Date = date
	result := GetDB().Save(e)
	return result.Error
}