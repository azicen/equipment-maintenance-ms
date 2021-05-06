package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

//AddGroup 添加权限组
func AddGroup(name string) (Group, error) {
	g := Group{
		Name: name,
	}
	result := GetDB().Create(&g)
	return g, result.Error
}

//DeleteGroup 删除权限组
func (g *Group) DeleteGroup() error {
	result := GetDB().Delete(g)
	return result.Error
}
