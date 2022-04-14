package pool

import (
	"chat_rooms/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	username := config.GetConfig().MySQL.User
	password := config.GetConfig().MySQL.Password
	host := config.GetConfig().MySQL.Host
	port := config.GetConfig().MySQL.Port
	Dbname := config.GetConfig().MySQL.Name
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)

}

func GetDB() *gorm.DB {
	return db
}
