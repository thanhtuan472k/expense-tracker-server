package service

import (
	"context"
	"errors"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/pkg/admin/dao"
	"expense-tracker-server/pkg/admin/errorcode"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubCategoryInterface ...
type SubCategoryInterface interface {
	// Create new subCategory ...
	Create(ctx context.Context, categoryID primitive.ObjectID, payload requestmodel.SubCategoryBodyCreate) (subCategoryID string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseSubCategoryAll)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseSubCategoryAdmin, err error)

	// Update ...
	Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryBodyUpdate) (subCategoryID string, err error)

	// ChangeStatus ...
	ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryChangeStatus) (subCategoryID string, err error)
}

// SubCategory ...
func SubCategory() SubCategoryInterface {
	return subCategoryImplement{}
}

// subCategoryImplement ...
type subCategoryImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s subCategoryImplement) Create(ctx context.Context, categoryID primitive.ObjectID, payload requestmodel.SubCategoryBodyCreate) (subCategoryID string, err error) {
	// Find category by id
	var categorySvc = categoryImplement{}
	category, _ := categorySvc.FindByID(ctx, categoryID)

	var (
		d   = dao.SubCategory()
		doc = payload.ConvertToBSON(category)
	)

	// Check category is active
	if !category.IsActiveCategory() {
		err = errors.New(errorcode.CategoryStatusInactive)
		return
	}

	// Create
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	subCategoryID = doc.ID.Hex()
	return

}

// All ...
func (s subCategoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseSubCategoryAll) {
	//TODO implement me
	panic("implement me")
}

// Detail ...
func (s subCategoryImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseSubCategoryAdmin, err error) {
	//TODO implement me
	panic("implement me")
}

// Update ...
func (s subCategoryImplement) Update(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryBodyUpdate) (categoryID string, err error) {
	//TODO implement me
	panic("implement me")
}

// ChangeStatus ...
func (s subCategoryImplement) ChangeStatus(ctx context.Context, id primitive.ObjectID, payload requestmodel.SubCategoryChangeStatus) (categoryID string, err error) {
	//TODO implement me
	panic("implement me")
}
