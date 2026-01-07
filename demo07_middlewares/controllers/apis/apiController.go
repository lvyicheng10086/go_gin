package apis

import (
	"demo07_middlewares/controllers/admin_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	admin_controller.BaseController
}

func (a *ApiController) Index(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": "接口首页",
	})
}

func (a *ApiController) UserList(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": "用户接口列表",
	})
}

func (a *ApiController) PList(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"msg": " 我是一个api接口-plist ",
	})
}
