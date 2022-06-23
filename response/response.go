package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"msg":  msg,
	})
}

func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	})
}
