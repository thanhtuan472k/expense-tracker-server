package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
)

// CategoryInterface ...
type CategoryInterface interface {
	// Create new category ...
	Create(ctx context.Context, payload requestmodel.CategoryBody) (categoryID string, err error)

	// All

	// Detail

	// Update

	// ChangeStatus
}

// Category ...
func Category() CategoryInterface {
	return categoryImplement{}
}

// categoryImplement ...
type categoryImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (c categoryImplement) Create(ctx context.Context, payload requestmodel.CategoryBody) (categoryID string, err error) {
	//TODO implement me
	panic("implement me")
}
