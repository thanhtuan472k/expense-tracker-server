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
	// Find user with payload.phone

	// Check valid password

	// Generate token

	// Return

	return
}

//
// PRIVATE METHODS ...
//

