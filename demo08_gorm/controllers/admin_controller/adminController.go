package admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 申明一个结构体用于后台控制器
type AdminController struct {
	BaseController //继承
}

func (a *AdminController) Index(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": "后台首页",
	})
	a.Success(g)
	//继承BaseController的Success方法

}

func (a *AdminController) User(g *gin.Context) {
	//从中间件中请求上下文获取值
	name, ok := g.Get("name")
	if ok {
		x, _ := name.(string)
		g.JSON(http.StatusOK, gin.H{
			"msg":  "用户列表",
			"name": x,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"msg": "未获取到用户信息",
		})
	}
}

func (a *AdminController) Article(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": "新闻列表",
	})
}
