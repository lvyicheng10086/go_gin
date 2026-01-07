package main

import (
	"net/http" //如果需要用 net/http 包中的常量表示响应码，请同时导入它

	"github.com/gin-gonic/gin"
)

// Gin结合 net/http使用
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好gin")
	})
	//POST
	r.POST("/new", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个POST请求")
	})
	r.DELETE("/del", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个Delete请求")
	})

	// r.Run()  启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run(":8080")

	r.Run() // 默认监听 0.0.0.0:8080
}
