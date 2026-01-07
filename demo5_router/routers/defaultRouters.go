package routers

import "github.com/gin-gonic/gin"

//外部引用函数需要大写
//创建函数的时候把gin框架对象引入，gin.Engine 相当于——> r := gin.Default()
func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")

	{
		defaultRouters.GET("/index", func(c *gin.Context) {
			c.String(200, "这是一个首页")
		})

		defaultRouters.GET("/new", func(c *gin.Context) {
			c.String(200, "这是新闻页面")
		})

	}
}
