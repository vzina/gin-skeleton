package controller

import (
	"github.com/vzina/gin-skeleton/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexController is the default controller
type IndexController struct{
	BaseController
}

// GetIndex home page
func (ctrl *IndexController) GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Gin Skeleton",
		"content": "This is a skeleton based on gin framework",
	})
}

// GetVersion version json
func (ctrl *IndexController) GetVersion(c *gin.Context) {
	//panic(errno.ErrInvalidArgs)
	//ctrl.JsonResponse(c, errno.ErrInvalidArgs)
	ctrl.JsonResponse(c, gin.H{
		"version": config.Server.Version,
	})
}
