package main

import (
	"github.com/gin-gonic/gin"
	"orangezoom.cn/ginessential/common"
	_ "orangezoom.cn/ginessential/model"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
