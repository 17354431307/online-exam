package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	r := gin.New()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world!")
	})
	r.Run(":8081")
}
