package system

import (
	v1 "backend/api/v1"
	"backend/model/system/request"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")

	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)

		baseRouter.POST("register", func(c *gin.Context) {
			var form request.Register

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
