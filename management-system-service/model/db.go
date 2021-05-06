package model

import (
	"management-system-server/util"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

//InitDatabase 初始化数据库
func InitDatabase() {
	sqldsn := os.Getenv("SQL_DSN")

	db, err := gorm.Open("mysql", sqldsn)
	db.LogMode(true)
	//Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

//执行数据迁移
func migration() {
	//自动迁移模式
	GetDB().AutoMigrate(&User{})
	GetDB().AutoMigrate(&Maintain{})
	GetDB().AutoMigrate(&Group{})
	GetDB().AutoMigrate(&Equipment{})
	GetDB().AutoMigrate(&EquipmentType{})
	GetDB().AutoMigrate(&EquipmentRelationalGroup{})
	GetDB().AutoMigrate(&UserRelationalGroup{})
}

func GetDB() *gorm.DB {
	return DB
}
