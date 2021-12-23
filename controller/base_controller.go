package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vzina/gin-skeleton/errno"
	"github.com/vzina/gin-skeleton/middleware"
	"net/http"
)

type BaseController struct {

}

func (b *BaseController) JsonResponse(c *gin.Context, data interface{})  {
	var resp errno.Message
	if sv, ok := data.(error); ok {
		resp = errno.GetMessageByErr(sv)
	} else {
		resp =  errno.GetMessageByErr(errno.OK)
		resp.WithData(data)
	}

	c.JSON(http.StatusOK, resp.WithRequestId(c.Value(middleware.TraceIdKey).(string)))
}
