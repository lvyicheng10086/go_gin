package main

import (
	"net/http" //如果需要用 net/http 包中的常量表示响应码，请同时导入它

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	router.Run() // 默认监听 0.0.0.0:8080
// }

// Gin结合 net/http使用
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 默认监听 0.0.0.0:8080
}
