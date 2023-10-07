package response

import "backend/model/system"

type SysUserResponse struct {
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expires_at"`
}
