package main

import (
	_ "learn-go/connect"
	views "learn-go/view"

	_ "learn-go/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

// @title Gin swagger  展示在web端的title上
// @version 1.0 定义接口的版本
// @description Gin swagger 示例项目 首页展示
// @securityDefinitions.apikey ApiKeyAuth  API的认证方式
// @in cookie 发送认证的方式
// @name token  后端获取认证值得方式
func main() {
	r := gin.Default()
	r.POST("/login/", views.Login)
	r.GET("/index/", views.Index)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
