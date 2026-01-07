package routers

import (
	"demo11/controllers/yukuaizheng"

	"github.com/gin-gonic/gin"
)

// YukuaizhengRouters 注册渝快政务APP相关路由
func YukuaizhengRouters(r *gin.Engine) {
	// 创建控制器实例
	ykController := yukuaizheng.YukuaizhengController{}

	// 创建路由组，统一前缀 /api/v1/yukuaizheng
	// 可以根据需要添加中间件，例如 middlewares.ApiAuth
	ykGroup := r.Group("/api/v1/yukuaizheng")

	// 这里假设部分接口需要登录认证，部分不需要，为了演示方便，这里先统一不加认证或根据需求加
	// ykGroup.Use(middlewares.ApiAuth)

	{
		// 1. 登录模块
		ykGroup.POST("/auth/login", ykController.Login)

		// 2. 医疗救助资金支出统计查询
		ykGroup.GET("/medical-assistance/stats", ykController.GetMedicalFundStats)

		// 3. 低收入人群信息录入 / 查询
		ykGroup.POST("/low-income-population", ykController.AddLowIncomeInfo)
		ykGroup.GET("/low-income-population", ykController.GetLowIncomeInfo)

		// 4. 预警信息推送与救助状态更新
		ykGroup.GET("/warnings", ykController.GetWarnings)
		ykGroup.PUT("/assistance-status/:id", ykController.UpdateAssistanceStatus)
	}
}
