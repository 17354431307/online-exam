package system

import (
	"backend/global"
	"backend/model/system"
	"backend/utils"
	"context"
)

type JwtService struct {
}

// JoinInBlacklist 拉黑 jwt
func (j *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.OE_DB.Create(&jwtList).Error
	if err != nil {
		return
	}

	global.BlackCahe.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlacklist 判断 JWT 是否在黑名单中
func (j *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCahe.Get(jwt)
	return ok

	// 存储到数据库中
	//err := global.OE_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	//isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	//return !isNotFound
}

// GetRedisJWT 从 redis 中取 jwt
func (j *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.OE_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT 把 jwt 存储到 redis 中
func (j *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于 jwt 过期时间
	dr, err := utils.ParseDuration(global.OE_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		return err
	}

	timer := dr
	err = global.OE_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
