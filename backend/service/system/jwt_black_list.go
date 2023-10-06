package system

import (
	"backend/global"
)

type JwtService struct {
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
