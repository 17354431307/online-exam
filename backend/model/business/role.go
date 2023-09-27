package business

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	name string
}
