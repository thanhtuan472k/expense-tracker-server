package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserInterface ...
type UserInterface interface {
	// Register ...
	Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error)

	// Login ...
	Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error)

	// GetMe ...
	GetMe(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseUserMe, err error)
}

// User ...
func User() UserInterface {
	return userImplement{}
}

// userImplement ...
type userImplement struct{}

//
// PUBLIC METHODS
//

// Register ...
func (s userImplement) Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error) {
	return
}

// Login ...
func (s userImplement) Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error) {
	return
}

// GetMe ...
func (s userImplement) GetMe(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseUserMe, err error) {
	return
}

// PRIVATE METHODS
