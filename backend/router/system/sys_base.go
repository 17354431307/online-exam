package system

import (
	"backend/model/system"
	"backend/utils"
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

		baseRouter.POST("register", func(c *gin.Context) {
			var form system.Register

			if err := c.ShouldBindJSON(&form); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	return baseRouter
}
