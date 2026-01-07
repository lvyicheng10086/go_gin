package routers

import (
	"demo12/controllers/apis"
	"demo12/middlewares"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func ApiRouters(r *gin.Engine) {

	apiRouters := r.Group("/api", middlewares.ApiAuth)
	api := apis.BankController{}

	{

		apiRouters.GET("/bank", api.GetBank)

	}
}
