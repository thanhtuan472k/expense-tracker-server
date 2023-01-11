package service

import (
	"context"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/pkg/admin/dao"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryInterface ...
type CategoryInterface interface {
	// Create new category ...
	Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseCategoryAll)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseCategoryAdmin, err error)

	// Update ...
	Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryBodyCreate)

	// ChangeStatus ...
	ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryChangeStatus) (result responsemodel.ResponseChangeStatus)
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
func (s categoryImplement) Create(ctx context.Context, payload requestmodel.CategoryBodyCreate) (categoryID string, err error) {
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
	return
}

// All ...
func (s categoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseCategoryAll) {
	//TODO implement me
	panic("implement me")
}

// Detail ...
func (s categoryImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseCategoryAdmin, err error) {
	//TODO implement me
	panic("implement me")
}

// Update ...
func (s categoryImplement) Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryBodyCreate) {
	//TODO implement me
	panic("implement me")
}

// ChangeStatus ...
func (s categoryImplement) ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.CategoryChangeStatus) (result responsemodel.ResponseChangeStatus) {
	//TODO implement me
	panic("implement me")
}
