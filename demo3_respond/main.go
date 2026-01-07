package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Shop struct {
	Price   float64 `json: price`
	Foods   string  `json: foods`
	Context string  `json: context`
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	//初始化一个默认路由
	r := gin.Default()
	//模板加载
	r.LoadHTMLGlob(func() string {
		if _, err := os.Stat("templates"); err == nil {
			return "templates/*"
		}
		if _, err := os.Stat(filepath.Join("demo3_respond", "templates")); err == nil {
			return filepath.Join("demo3_respond", "templates/*")
		}
		return "templates/*"
	}())

	//String
	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "这是字符串类型")
	})

	//json-1
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	//map[string]interface{}等价写法
	r.GET("/gin_H", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("/struct", func(c *gin.Context) {

		a := Shop{
			Price:   100,
			Foods:   "牛肉",
			Context: "测试测试内容",
		}

		c.JSON(http.StatusOK, a)
	})

	//响应Jsonp请求,跟回调函数实现跨域请求，?callback=func()
	// http://localhost:8080/jsonp?callback=xxxx
	//返回效果：// xxxx({"title":"我是一个标题-jsonp","desc":"描述","content":"测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题-jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(http.StatusOK, a)
	})

	//html
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "这是一个html",
		})
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin 我是一个xml",
		})
	})
	r.Run(":8091")
}
