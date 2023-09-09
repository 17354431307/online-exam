package request

type CreateExaCustomerRequest struct {
	CustomerName       string `json:"customerName" form:"customerName"`             // 客户名
	CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData"`   // 客户手机号
	SysUserID          uint   `json:"sysUserID" form:"sysUserID"`                   // 管理ID
	SysUserAuthorityID uint   `json:"sysUserAuthorityID" form:"sysUserAuthorityID"` // 管理角色ID
}
