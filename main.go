package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
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
			name = RandomString(10)
		}

		log.Println("name:" + name + "\t telephone:" + telephone + "\t password:" + password)

		// 判断手机号是否存在

		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号存在"})
			return
		}

		// 创建用户

		user := User{Name: name, Telephone: telephone, Password: password}

		db.Create(user)

		//返回结果

		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run())
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters)-1)]
	}

	return string(result)
}

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

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
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil
	}
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {

	var user User

	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
