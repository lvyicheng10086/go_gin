package main

import (
	"demo12/middlewares"
	"demo12/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//全局中引入Logger中间件
	r.Use(middlewares.Logger)

	// 注册银行相关路由

	routers.ApiRouters(r)

	r.Run(":8011")
}
