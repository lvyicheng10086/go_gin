package main

import (
	"demo06_controller/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.DefaultRouters(r)
	routers.AdminRoutersInit(r)
	routers.ApiRouters(r)
	r.Run(":8099")
}
