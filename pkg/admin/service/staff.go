package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
)

// StaffInterface ...
type StaffInterface interface {
	// Login ...
	Login(ctx context.Context, payload requestmodel.StaffBodyLogin) (success responsemodel.ResponseLoginSuccess, err error)
}

// Staff ...
func Staff() StaffInterface {
	return staffImplement{}
}

// staffImplement ...
type staffImplement struct{}

//
// PUBLIC METHODS ...
//

// Login ...
func (s staffImplement) Login(ctx context.Context, payload requestmodel.StaffBodyLogin) (success responsemodel.ResponseLoginSuccess, err error) {
	// Check staff (email, phone) is existed in system or not

	// If staff is not existed --> User not found

	// If staff existed
	// - payload.Password (hashed) and compare with hasedPassword in DB
	// - If wrong password --> Password is incorrect
	// - If success password --> Generate token and send response

	// Return

	return
}

//
// PRIVATE METHODS ...
//
