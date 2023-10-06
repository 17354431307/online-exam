package request

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
