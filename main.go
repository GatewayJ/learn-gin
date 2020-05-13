package main

import (
	_ "learn-go/connect"
	views "learn-go/view"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login/", views.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
