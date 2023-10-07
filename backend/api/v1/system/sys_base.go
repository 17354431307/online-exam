package system

import (
	"backend/global"
	"backend/model/common/response"
	"backend/model/system"
	systemReq "backend/model/system/request"
	systemRes "backend/model/system/response"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type BaseApi struct {
}

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// TODO 判断验证码
	//key := c.ClientIP()
	//v, ok := global.BlackCahe.Get(key)
	//if !ok {
	//	global.BlackCahe.Set(key, 1, time.Duration())
	//}

	if l.Username == "he" && l.Password == "123456" {
		user := system.SysUser{
			UUID:        uuid.Must(uuid.NewV4()),
			Password:    "123456",
			Username:    "he",
			Nickname:    "摸鱼",
			AuthorityId: 1,
		}
		b.TokenNext(c, user)
		return
	} else {
		response.FailWithMessage("用户名不存在或密码错误", c)
		return
	}
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	// 颁发 token
	j := &utils.JWT{SigningKey: []byte(global.OE_CONFIG.Jwt.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.OE_Log.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败!", c)
		return
	}

	// TODO 多端

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.OE_Log.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.OE_Log.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		// 这里就是把老token给拉黑，生成新 token
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if jwtService.JoinInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
