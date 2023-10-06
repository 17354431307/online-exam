package middleware

import (
	"backend/global"
	"backend/model/common/response"
	"backend/service"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 这里 jwt 鉴取头部信息 x-token 登录时返回 token 信息，这里前端需要把 token 存储到 cookie 或者本地 localStorage,
		// 不过需要跟后端协商过期时间，可以约定刷新令牌或重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的账户异地登录或令牌失效", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// ParseToken 解析 token 包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}

			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// token 续命，这里的意思是，claims 中的剩余时间如果要小于 claims 的缓存时间的话，就会进行 token 续命
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.OE_CONFIG.Jwt.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))

			// TODO 多点登录拦截待实现
			if global.OE_CONFIG.App.UseMultipoint {
				// ...
			}
		}

		c.Set("claims", c)
		c.Next()
	}
}
