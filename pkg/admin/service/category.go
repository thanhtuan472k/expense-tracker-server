package service

import (
	"context"
	"expense-tracker-server/pkg/admin/dao"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	"fmt"
)

// CategoryInterface ...
type CategoryInterface interface {
	// Create new category ...
	Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error)

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
func (c categoryImplement) Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error) {
	var (
		d   = dao.Category()
		doc = payload.ConvertToBSON()
	)

	// Create
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	categoryID = doc.ID.Hex()
	fmt.Println("tesst")
	return
}
