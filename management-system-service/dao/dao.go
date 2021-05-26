package dao

import (
	"management-system-server/model"
	log "management-system-server/util/logger"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" 
)

var dao *Dao

type Dao struct {
	db *gorm.DB
}

// New
func New() (d *Dao, err error) {
	db, err := NewMySQL()
	d = &Dao{db}
	dao = d

	return
}

//初始化数据库
func NewMySQL() (db *gorm.DB, err error) {
	sqldsn := os.Getenv("SQL_DSN")

	db, err = gorm.Open("mysql", sqldsn)
	db.LogMode(true)
	// Error
	if err != nil {
		log.Panic(err.Error(), err)
	}
	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(50)
	// 打开
	db.DB().SetMaxOpenConns(100)
	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Maintain{})
	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.Equipment{})
	db.AutoMigrate(&model.EquipmentType{})
	db.AutoMigrate(&model.EquipmentInGroup{})
	db.AutoMigrate(&model.UserInGroup{})

	return
}

func GetDao() *Dao {
	return dao
}

func (d *Dao) GetDB() *gorm.DB {
	return d.db
}
