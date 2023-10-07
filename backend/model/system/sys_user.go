package system

import (
	"backend/global"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.OE_MODEL
	UUID        uuid.UUID
	Username    string
	Password    string
	Nickname    string
	AuthorityId uint
}
