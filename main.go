package main

import (
	_ "learn-go/connect"
	views "learn-go/view"
	"net/http"

	_ "learn-go/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}

}

// @title Gin swagger  展示在web端的title上
// @version 1.0 定义接口的版本
// @description Gin swagger 示例项目 首页展示
// @securityDefinitions.apikey ApiKeyAuth  API的认证方式
// @in cookie 发送认证的方式
// @name token  后端获取认证值得方式
func main() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/login/", views.Login)
	r.GET("/index/", views.Index)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
