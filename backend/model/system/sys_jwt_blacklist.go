package system

import "backend/global"

type JwtBlacklist struct {
	global.OE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
