package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"orangezoom.cn/ginessential/common"
	_ "orangezoom.cn/ginessential/model"
	"os"
)

func main() {
	InitConfig()

	log.Println(viper.GetString("database.host"))
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run(":" + port))
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "\\config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
