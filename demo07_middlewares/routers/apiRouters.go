package routers

import (
	"demo07_middlewares/controllers/apis"
	"demo07_middlewares/middlewares"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func ApiRouters(r *gin.Engine) {

	//路由中引入ApiAuth中间件
	apiRouters := r.Group("/api", middlewares.ApiAuth)
	api := apis.ApiController{}

	{
		apiRouters.GET("/", api.Index)
		apiRouters.GET("/PList", api.PList)
		apiRouters.GET("/userList", api.UserList)
	}
}
