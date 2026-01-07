package routers

import (
	"demo10/controllers/admin_controller"
	"demo10/middlewares"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 是gin框架的对象
func StudentRoutersInit(r *gin.Engine) {
	//r := gin.Default()
	//把中间件集成到路由组中
	studentRouters := r.Group("/student", middlewares.AdminAuth, middlewares.SetValue)
	studentactul := admin_controller.StudentController{}

	{
		studentRouters.GET("/", studentactul.GetUser)          // 获取用户 (查询)
		studentRouters.GET("/add", studentactul.AddUser)       // 新增用户 (创建)
		studentRouters.GET("/edit", studentactul.UpdateUser)   // 更新用户 (修改)
		studentRouters.GET("/delete", studentactul.DeleteUser) // 删除用户 (删除)
	}
}
