package initialize

import (
	"backend/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	PublicGroup := Router.Group("")

	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由
		exampleRouter.InitCustomerRouter(PublicGroup)
	}

	return Router
}
