package views

import (
	form "learn-go/form"
	jwtUtil "learn-go/rbac"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var up form.LoginForm
	err := c.ShouldBind(&up)

	claims := jwtUtil.CustomClaims{
		"Name":     "name",
		"Password": "pd",
		myjwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "new",                           //签名的发行者
		},
	}
	jwtUtil.CreateToken(claims)
	if err != nil {
		log.Println(111111, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
	} else {
		log.Printf("%+v\n", up)
		c.SetCookie("token", "thisistoken", 1000, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
	}
}
