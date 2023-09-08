package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")

	{
		baseRouter.POST("login", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return baseRouter
}
