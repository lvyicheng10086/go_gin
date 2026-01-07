package routers

import (
	"demo06_controller/controllers/defualt"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	defualt := defualt.IndexController{}

	{
		defaultRouters.GET("/index", defualt.Index)
		defaultRouters.GET("/new", defualt.New)

	}
}
