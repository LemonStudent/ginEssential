package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orangezoom.cn/ginessential/common"
	"orangezoom.cn/ginessential/model"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenValue := context.GetHeader("Authorization")

		if tokenValue == "" || !strings.HasPrefix(tokenValue, "Bearer ") {
			context.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"msg": "权限不足"})
			context.Abort()
			return
		}

		tokenValue = tokenValue[7:]

		token, claims, err := common.ParseToken(tokenValue)

		if err != nil || !token.Valid {
			context.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"msg": "权限不足"})
			context.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			context.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"msg": "权限不足"})
			context.Abort()
			return
		}

		context.Set("user", user)
	}
}
