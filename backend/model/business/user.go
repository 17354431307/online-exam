package business

import "backend/global"

type UserStatus int

const (
	OFFLIEN = 0 // 离线
	ONLINE  = 1
	INEXAM  = 2
)

type User struct {
	global.OE_MODEL
	Name      string
	Phone     string
	Photo     string
	LoginName string
	Password  string
	CardID    string `gorm:"unique"`
	Role      int
	Status    UserStatus
}
