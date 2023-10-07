package utils

import (
	"backend/global"
	"backend/model/system/request"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenMalformed   = errors.New("That's not even a token")
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.OE_CONFIG.Jwt.SigningKey),
	}
}

// CreateClaims 创建 jwt claims 自定义声明部分
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.OE_CONFIG.Jwt.BufferTime)
	ep, _ := ParseDuration(global.OE_CONFIG.Jwt.ExpiresTime)

	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间为1天，缓冲时间内会获得新的token令牌，此时一个用户会存在两个有效的令牌，但是前端只留一个，另一个丢弃
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"OE"},                    // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间，配置文件
			Issuer:    global.OE_CONFIG.Jwt.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个 token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.OE_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})

	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, TokenMalformed
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, TokenExpired
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, TokenNotValidYet
		}
		return nil, TokenInvalid
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, err
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
