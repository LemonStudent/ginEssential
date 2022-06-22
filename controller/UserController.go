package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"orangezoom.cn/ginessential/common"
	"orangezoom.cn/ginessential/model"
	"orangezoom.cn/ginessential/util"
)

// Register 用户注册
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
	encryptionPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密异常"})
		return
	}

	var user = model.User{Name: name, Telephone: telephone, Password: string(encryptionPassword)}
	DB.Create(&user)

	//返回结果

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	var user model.User

	DB.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号不存在"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码错误"})
		return
	}

	token, err := common.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "token 生成失败"})
		return

	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": 200,
			"msg":  "登陆成功",
			"data": gin.H{"token": token},
		})
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code": 200,
			"msg":  "登陆成功",
			"data": gin.H{"user": user},
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
