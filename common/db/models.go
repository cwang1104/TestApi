package db

import (
	"TestApi/common/config"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func init() {

	db, err := gorm.Open(config.Db_driver, config.Db_source)
	if err != nil {
		logs.Error("err open db = %s", err.Error())
		return
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(config.DbMaxOpenConns)
	db.DB().SetMaxIdleConns(config.DbMaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.DbMaxLifeTime) * time.Second)
	DB = db
	logs.Info("DB init success...")
}

type User struct {
	Phone      string    `gorm:"column:phone"`
	UID        int32     `gorm:"column:uid"`
	NickName   string    `gorm:"column:nick_name"`
	HeadURL    string    `gorm:"column:head_url"`
	Gender     int32     `gorm:"column:gender"`
	Position   string    `gorm:"column:position"`
	Birthday   string    `gorm:"column:birthday"`
	Email      string    `gorm:"column:email"`
	Introduce  string    `gorm:"column:introduce"`
	Love       int32     `gorm:"column:love"`
	CreateTime time.Time `gorm:"column:create_time"`
}
