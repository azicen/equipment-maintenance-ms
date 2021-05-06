package model

import "github.com/jinzhu/gorm"

type EquipmentType struct {
	gorm.Model
	Id    int    `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Cycle int    `gorm:"not null"` //维护周期
}

//AddEquipmentType 添加设备类型
func AddEquipmentType(name string, cycle int) (EquipmentType, error) {
	et := EquipmentType{
		Name:  name,
		Cycle: cycle,
	}
	result := GetDB().Create(&et)
	return et, result.Error
}

//DeleteEquipmentType 删除设备类型
func (et *EquipmentType) DeleteEquipmentType() error {
	result := GetDB().Delete(et)
	return result.Error
}

//SetCycle 设置设备维护周期
func (et *EquipmentType) SetCycle(cycle int) error {
	et.Cycle = cycle
	result := GetDB().Save(et)
	return result.Error
}
