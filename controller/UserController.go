package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"orangezoom.cn/ginessential/common"
	"orangezoom.cn/ginessential/model"
	"orangezoom.cn/ginessential/util"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不得少于6位"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 判断手机号是否存在

	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号存在"})
		return
	}

	// 创建用户

	var user = model.User{Name: name, Telephone: telephone, Password: password}
	DB.Create(user)

	//返回结果

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {

	var user model.User

	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
