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

	token, err := rbac.GenerateToken(up.User)
	if err != nil {
		log.Println(111111, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
	} else {
		log.Printf("%+v\n", up)
		c.SetCookie("token", "thisistoken", 1000, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"msg": token})
	}
}

func Index(c *gin.Context) {
	ua2 := c.Request.Header.Get("token")
	claim, err := rbac.ParesToken(ua2)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": claim})
	}

}
