package views

import (
	form "learn-go/form"
	rbac "learn-go/rbac"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var up form.LoginForm
	err := c.ShouldBind(&up)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		token, err := rbac.GenerateToken(up.User)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
		} else {
			log.Printf("%+v\n", up)
			c.SetCookie("token", token, 1000, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"msg": token})
		}
	}
}

func Index(c *gin.Context) {
	ua2 := c.Request.Header.Get("token")
	println("asasasasa", ua2)
	a, err := rbac.RefreshToken(ua2)
	println(a)
	if err != nil {
		log.Println("ParesToken: ", err.Error())
		c.JSON(http.StatusOK, gin.H{"msg": "请求token失效1"})
	} else {
		token, err := rbac.RefreshToken(ua2)
		if err != nil {
			log.Println("RefreshToken: ", err.Error())
			c.JSON(http.StatusOK, gin.H{"msg": "请求token失效2"})
		} else {
			c.SetCookie("token", token, 1000, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"msg": token})
		}
	}

}
