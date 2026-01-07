package routers

import "github.com/gin-gonic/gin"

//外部引用函数需要大写
//创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")

	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "这是一个接口")
		})
		apiRouters.GET("/List", func(c *gin.Context) {
			c.String(200, "这是一个接口列表")
		})
		apiRouters.GET("/userList", func(c *gin.Context) {
			c.String(200, "这是一个用户列表")
		})
	}
}
