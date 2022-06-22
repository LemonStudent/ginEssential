package main

import (
	"github.com/gin-gonic/gin"
	"orangezoom.cn/ginessential/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
