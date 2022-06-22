package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orangezoom.cn/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	database := "essential"
	username := "root"
	password := "root"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
