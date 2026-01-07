package routers

import (
	"demo10/controllers/admin_controller"
	"demo10/middlewares"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 是gin框架的对象
func AdminRoutersInit(r *gin.Engine) {
	//r := gin.Default()
	//把中间件集成到路由组中
	adminRouters := r.Group("/admin", middlewares.AdminAuth, middlewares.SetValue)
	adminactul := admin_controller.AdminController{}

	{
		adminRouters.GET("/", adminactul.Index)
		adminRouters.GET("/user", adminactul.User)
		adminRouters.GET("/article", adminactul.Article)
	}
}
