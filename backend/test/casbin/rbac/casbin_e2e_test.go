//go:build e2e

package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// initEnforcer 用 grom 的 adapter 初始化 csbin enforcer
func initEnforcer() (*casbin.Enforcer, error) {
	// 初始化 grom 适配器，在 casbin 的执行器中使用
	// gorm 使用 mysql 中已经 存在的 casbin_demo_db 数据库，并会使用 casbin_rule 表，如果表不存在会创建
	// 第三个参数 true 就是代表着不存在就创建
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

// 检验是否登录
func CheekLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里的 token 只是为了方便测试，直接把用户名当成 token 了
		if c.Request.Header.Get("token") == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "token required",
			})
		} else {
			c.Set("user_name", c.Request.Header.Get("token"))
			c.Next()
		}
	}
}

func RBAC() gin.HandlerFunc {
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		user, _ := c.Get("user_name")

		fmt.Println(enforcer.GetPolicy())
		if has, err := enforcer.Enforce(user, c.Request.RequestURI, c.Request.Method); err != nil || !has {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"msg": "forbidden",
			})
		} else {
			c.Next()
		}
	}

}

func Test_Casbin_RBAC(t *testing.T) {

	r := gin.New()

	r.Use(CheekLogin(), RBAC())

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "所有文章",
		})
	})
	r.POST("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "批量修改文章",
		})
	})

	r.Run(":8081")
}
