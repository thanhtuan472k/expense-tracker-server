package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
)

// UserInterface ...
type UserInterface interface {
	// Register ...
	Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error)

	// Login ...
	Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error)
}

// User ...
func User() UserInterface {
	return userImplement{}
}

// userImplement ...
type userImplement struct{}

//
// PUBLIC
//

// Register ...
func (s userImplement) Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error) {
	return
}

// Login ...
func (s userImplement) Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error) {
	return
}
