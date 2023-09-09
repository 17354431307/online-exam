package example

import (
	"backend/global"
	"backend/model/example"
)

type CustomerService struct {
}

// @author: 模板样例
// @function: CreateExaCustomer
// @description: 创建客户
// @param: e model.ExaCustomer
// @return: err error
func (c *CustomerService) CreateExaCustomer(e example.ExaCusmoter) (err error) {
	err = global.OE_DB.Create(&e).Error
	return err
}

// @author: 模板样例
// @function: DeleteExCustomer
// @description: 删除客户
// @param: e model.ExaCusmoter
// @return: err error
func (c *CustomerService) DeleteExCustomer(e example.ExaCusmoter) (err error) {
	err = global.OE_DB.Delete(&e).Error
	return err
}

// @author: 模板样例
// @function: UpdateExaCustomer
// @description: 更新用户
// @param: e *model.ExaCusmoter
// @return: err error
func (c *CustomerService) UpdateExaCustomer(e *example.ExaCusmoter) (err error) {
	err = global.OE_DB.Save(e).Error
	return err
}

// @author: 模板样例
// @function: GetExaCustomer
// @description: 获取客户信息
// @param: id uint
// @return: customer model.ExaCusmoter, err error
func (c *CustomerService) GetExaCustomer(id uint) (customer example.ExaCusmoter, err error) {
	err = global.OE_DB.Where("id = ?", id).First(&customer).Error
	return
}
