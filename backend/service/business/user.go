package business

import (
	"backend/global"
	"backend/model/business"
)

type UserService struct {
}

func (u *UserService) CreateUser(user business.User) (err error) {
	err = global.OE_DB.Create(&user).Error
	return err
}
