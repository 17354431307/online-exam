package example

import "backend/global"

type ExaCusmoter struct {
	global.OE_MODEL
	CustomerName       string `json:"customerName" form:"customerName" gorm:"comment:客户名"`
	CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`
	SysUserID          uint   `json:"sysUserID" form:"sysUserID" gorm:"comment:管理ID"`
	SysUserAuthorityID uint   `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
}
