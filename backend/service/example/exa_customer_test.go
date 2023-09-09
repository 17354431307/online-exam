package example_test

import (
	"backend/core"
	"backend/global"
	"backend/initialize"
	"backend/model/example"
	"backend/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCustomerService_CreateExaCustomer(t *testing.T) {
	customerService := service.ServiceGroupApp.ExampleServiceGroup.CustomerService

	testCases := []struct {
		name     string
		customer example.ExaCusmoter
	}{
		{
			name: "create customer",
			customer: example.ExaCusmoter{
				CustomerName:       "何小文",
				CustomerPhoneData:  "12345678",
				SysUserID:          1,
				SysUserAuthorityID: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := customerService.CreateExaCustomer(tc.customer)
			assert.NoError(t, err)
		})
	}
}

func TestCustomerService_DeleteExCustomer(t *testing.T) {
	customerService := service.ServiceGroupApp.ExampleServiceGroup.CustomerService

	testCases := []struct {
		name     string
		customer example.ExaCusmoter
	}{
		{
			name: "delete customer by id",
			customer: func() example.ExaCusmoter {
				res := example.ExaCusmoter{}
				res.ID = 1
				return res
			}(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := customerService.DeleteExCustomer(tc.customer)
			assert.NoError(t, err)
		})
	}
}

func TestCustomerService_GetExaCustomer(t *testing.T) {
	customerService := service.ServiceGroupApp.ExampleServiceGroup.CustomerService

	testCases := []struct {
		name    string
		id      uint
		wantErr error
	}{
		{
			name:    "no exists",
			id:      1,
			wantErr: gorm.ErrRecordNotFound,
		},
		{
			name: "get by id",
			id:   2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customer, err := customerService.GetExaCustomer(tc.id)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.id, customer.ID)
		})
	}
}

func TestCustomerService_UpdateExaCustomer(t *testing.T) {
	customerService := service.ServiceGroupApp.ExampleServiceGroup.CustomerService

	testCases := []struct {
		name     string
		id       uint
		changeFn func(curtomer *example.ExaCusmoter)
	}{
		{
			name: "update",
			id:   2,
			changeFn: func(curtomer *example.ExaCusmoter) {
				curtomer.CustomerName = "王辰"
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customer, err := customerService.GetExaCustomer(tc.id)
			assert.NoError(t, err)
			tc.changeFn(&customer)
			err = customerService.UpdateExaCustomer(&customer)
			assert.NoError(t, err)
		})
	}
}

func init() {
	global.OE_VIPER = core.InitViper("../../etc/config.yaml")
	global.OE_Log = core.InitializeZap()
	global.OE_DB = initialize.Gorm()
	if global.OE_DB != nil {
		initialize.RegisterTables()
	}
}
