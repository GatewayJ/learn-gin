package views

import (
	form "learn-go/form"
	rbac "learn-go/rbac"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登录接口
// @Summary 登录接口
// @Description 登录接口
// @Produce  json
// @Param user body form.LoginForm true "user"
// @Success 200 {object}  gin.H
// @Router /login/ [post]
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
			c.SetCookie("token", token, 1000, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"msg": token})
		}
	}
}

// Index 主页接口
// @Summary 主页接口
// @Produce  json
// @Security ApiKeyAuth
// @in header
// @ token Cookie int false "token"
// @Success 200 {object} gin.H
// @Router /index/ [GET]
func Index(c *gin.Context) {
	ua2, err := c.Request.Cookie("token")
	if err != nil {
		log.Println("ParesToken: ", ua2, "err: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求token失效"})
	} else {
		newToken, err := rbac.RefreshToken(ua2.Value)
		if err != nil {
			log.Println("ParesToken: ", ua2, "err: ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"msg": "请求token失效"})
		} else {
			c.SetCookie("token", newToken, 1000, "/", "localhost", false, false)
			c.JSON(http.StatusOK, gin.H{"index": "this is index"})
		}
	}

}
