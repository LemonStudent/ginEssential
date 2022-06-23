package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"orangezoom.cn/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
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
