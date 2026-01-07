package main

import (
	"demo5_router/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认路由
	r := gin.Default()

	routers.DefaultRouters(r)
	routers.AdminRoutersInit(r)
	routers.ApiRouters(r)

	r.Run(":8080")

}
