package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubCategoryInterface ...
type SubCategoryInterface interface {
}

// SubCategory ...
func SubCategory() SubCategoryInterface {
	return subCategoryImplement{}
}

// subCategoryImplement ...
type subCategoryImplement struct{}

//
// PUBLIC METHOD
//

//
// PRIVATE METHODS
//

// findDocSubCategoryTypeExpenseAvailableByID ...
func (s subCategoryImplement) findDocSubCategoryTypeExpenseAvailableByID(ctx context.Context, subCategoryID primitive.ObjectID) (doc mgexpense.SubCategory, err error) {
	var (
		d    = dao.SubCategory()
		cond = bson.D{{"_id", subCategoryID}}
	)

	doc = d.FindOneByCondition(ctx, cond)
	if !doc.IsSubCategoryAvailableByTypeExpense() {
		err = errors.New(errorcode.SubCategoryIsInvalid)
	}
	return
}
