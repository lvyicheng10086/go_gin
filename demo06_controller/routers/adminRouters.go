package routers

import (
	"demo06_controller/controllers/admin_controller"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 是gin框架的对象
func AdminRoutersInit(r *gin.Engine) {
	//r := gin.Default()
	adminRouters := r.Group("/admin")

	adminactul := admin_controller.AdminController{}

	{
		adminRouters.GET("/", adminactul.Index)
		adminRouters.GET("/user", adminactul.User)
		adminRouters.GET("/article", adminactul.Article)
	}
}
