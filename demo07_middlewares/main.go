package main

import (
	"demo07_middlewares/middlewares"
	"demo07_middlewares/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//全局中引入Logger中间件
	r.Use(middlewares.Logger)
	routers.DefaultRouters(r)
	routers.AdminRoutersInit(r)
	routers.ApiRouters(r)
	r.Run(":8011")
}
