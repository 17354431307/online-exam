package core

import (
	"backend/global"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	address := fmt.Sprintf(":%d", global.OE_CONFIG.App.Port)
	s := initServer(address, r)

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)

	fmt.Println(`address`, address)

	s.ListenAndServe() // 启动并监听 HTTP 服务
}

// initServer 实现 HTTP 的无缝重启和优雅关闭
func initServer(address string, router *gin.Engine) server {
	// endless 库主要的作用是实现 HTTP 的无缝重启和优雅关闭

	// 使用 endless 库创建一个 HTTP 服务器
	s := endless.NewServer(address, router)

	// 设置 HTTP 请求头读取时间为 20 s
	s.ReadHeaderTimeout = 20 * time.Second

	// 设置 HTTP 响应体的写入时间为 20 s
	s.WriteTimeout = 20 * time.Second

	// 设置 HTTP 请求头的最大字节数为 1Mb
	s.MaxHeaderBytes = 1 << 20

	return s
}
