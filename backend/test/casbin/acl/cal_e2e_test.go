//go:build e2e

package acl

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

/*
	sub 访问资源的用户
	obj 将要访问的资源
	act 用户对资源的操作
*/

// initCasBinWithGorm 用 grom 的 adapter 初始化 csbin enforcer
func initCasBinWithGorm() (*casbin.Enforcer, error) {
	// 初始化 grom 适配器，在 casbin 的执行器中使用
	// gorm 使用 mysql 中已经 存在的 casbin_demo_db 数据库，并会使用 casbin_rule 表，如果表不存在会创建
	a, err := gormadapter.NewAdapter(
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/casbin_demo_db?charset=utf8mb4&parseTime=true&loc=Local",
		true,
	)
	if err != nil {
		return nil, fmt.Errorf("gormadapter.NewAdapter 调用失败, %v", err)
	}

	// 创建 casbin 适配器
	e, err := casbin.NewEnforcer("./model.conf.conf", a)
	if err != nil {
		return nil, fmt.Errorf("casbin.NewEnforcer 调用失败, %v", err)
	}

	// 从数据库加载 policy
	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("从 db 加载 policy 失败, %v", err)
	}

	return e, nil
}

// myAuth 拦截器
func myAuth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取资源路径
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户，一般会从 token 中解析出来的
		sub := c.Query("sub")

		// 使用 Enforce 来检查是否具有权限
		ok, err := e.Enforce(sub, obj, act)
		if !ok {
			log.Printf("权限检查失败, %v\n", err)
			c.Abort()
			c.JSON(http.StatusUnauthorized, "没有权限")
		} else {
			log.Println("权限检查成功!")
			c.Next()
		}
	}
}

func Test_Casbin_ACL(t *testing.T) {
	e, err := initCasBinWithGorm()
	assert.NoError(t, err)

	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		log.Println("添加一个策略!")

		sub := c.PostForm("sub")
		obj := c.PostForm("obj")
		act := c.PostForm("act")

		// 添加一个 policy
		if ok, err := e.AddPolicy(sub, obj, act); !ok {
			log.Printf("添加策略失败, %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "添加策略失败!",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"meg": "添加策略成功!",
		})
	})

	// 注册全局中间件, 后续所有的请求都会通过这个拦截器
	r.Use(myAuth(e))
	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello, world!")
	})

	r.Run(":8081")
}
