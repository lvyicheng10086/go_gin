package routers

import (
	"demo10/controllers/admin_controller"

	"github.com/gin-gonic/gin"
)

// 外部引用函数需要大写
// 创建函数的时候把gin框架对象引入，gin.Engine 是gin框架的对象
func SalaryRoutersInit(r *gin.Engine) {
	salaryRouters := r.Group("/salary")

	navController := admin_controller.NavController{}

	{
		// salaryRouters.GET("/product", navController.AddProduct)
		// salaryRouters.GET("/product2", navController.OrderByProduct)
		// salaryRouters.GET("/limit", navController.Limit)
		// salaryRouters.GET("/query", navController.QuerySalary)
		// salaryRouters.GET("/delete", navController.DeleteSalary)
		// salaryRouters.GET("/update", navController.UpdateSalary)
		salaryRouters.GET("/select", navController.SelectSalary)
	}
}
