package admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c BaseController) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}

func (c BaseController) Error(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "error",
	})
}
