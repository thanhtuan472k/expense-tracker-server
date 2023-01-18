package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
)

// StaffInterface ...
type StaffInterface interface {
	// Create new category ...
	Create(ctx context.Context, payload requestmodel.StaffBodyCreate) (staffID string, err error)
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

func (s staffImplement) Create(ctx context.Context, payload requestmodel.StaffBodyCreate) (staffID string, err error) {
	var ()

	// Check email existed

	// Check phone existed

	// Create staff

	// Response
	return
}

//
// PRIVATE METHODS ...
//
