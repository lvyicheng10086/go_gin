package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义一个日志中间件，用于记录请求的路径和参数
func Logger(c *gin.Context) {
	url := c.Request.URL.Path
	params := c.Request.URL.Query()
	fmt.Println("这是全局Logger中间件")
	fmt.Println("获取的路径是：", url)
	fmt.Println("获取的参数是：", params)
}
