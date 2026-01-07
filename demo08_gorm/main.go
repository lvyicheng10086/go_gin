package main

import (
	"demo08_gorm/middlewares"
	"demo08_gorm/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//全局中引入Logger中间件
	r.Use(middlewares.Logger)
	routers.DefaultRouters(r)
	routers.AdminRoutersInit(r)
	// 注册学生相关路由
	routers.StudentRoutersInit(r)
	routers.ApiRouters(r)

	r.Run(":8011")
}
