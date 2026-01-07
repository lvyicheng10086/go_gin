package main

import (
	"demo11/middlewares"
	"demo11/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//全局中引入Logger中间件
	r.Use(middlewares.Logger)

	// 注册学生相关路由
	routers.ApiRouters(r)

	// 注册渝快政务相关路由
	routers.YukuaizhengRouters(r)

	r.Run(":8011")
}
