package admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
//申明一个结构体用于后台控制器
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
	g.JSON(http.StatusOK, gin.H{
		"msg": "用户列表",
	})
}

func (a *AdminController) Article(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": "新闻列表",
	})
}
