package request

import "net/http"

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint         `json:"authority_id"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbin_infos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/base/login", Method: http.MethodPost},
	}
}
