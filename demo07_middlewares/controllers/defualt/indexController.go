package defualt

import "github.com/gin-gonic/gin"

type IndexController struct{}

func (i *IndexController) Index(c *gin.Context) {
	c.String(200, "这是一个首页 - Controller")
}

func (i *IndexController) New(c *gin.Context) {
	c.String(200, "这是新闻页面 - Controller")
}
