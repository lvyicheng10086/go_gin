package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminAuth(c *gin.Context) {
	url := c.Request.URL.Path

	fmt.Println("这是路由中间件获取的路径是：", url)
}

func ApiAuth(c *gin.Context) {
	url := c.Request.URL.Path

	fmt.Println("获取的路径是：", url)
}

// 定义一个中间件，用于设置请求上下文的值
func SetValue(c *gin.Context) {

	c.Set("name", "张三")
	// 调用后续的中间件和处理函数
	c.Next()
}
