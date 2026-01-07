package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// `json:"username" form:"username" `
// `json:"password" from:"password"`
type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"page":     page,
		})
	})

	//默认值，id。不需要传参数
	r.GET("/defualt", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "ok",
		})
	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	//url中的动态参数
	r.GET("/List/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.String(200, "uid=%v", uid)
	})

	//解析xml格式
	r.POST("/Article", func(c *gin.Context) {
		article := &Article{}
		//Gin自动读取 request body并解析XML
		if err := c.ShouldBindXML(&article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})

	r.Run(":8080")
}
