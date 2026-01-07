package routers

import "github.com/gin-gonic/gin"

//外部引用函数需要大写
//创建函数的时候把gin框架对象引入，gin.Engine 是gin框架的对象
func AdminRoutersInit(r *gin.Engine) {
	//r := gin.Default()
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "新闻列表")
		})
	}
}
