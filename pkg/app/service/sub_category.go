package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/pkg/app/dao"
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

// findDocSubCategoryTypeIncomeAvailableByID ...
func (s subCategoryImplement) findDocSubCategoryTypeIncomeAvailableByID(ctx context.Context, id primitive.ObjectID) (doc mgexpense.SubCategory, err error) {
	var (
		d    = dao.SubCategory()
		cond = bson.D{{"_id", id}}
	)

	doc = d.FindOneByCondition(ctx, cond)
	if !doc.IsSubCategoryAvailableByTypeIncome() {
		err = errors.New("danh mucj con khong howpj lej") // TODO: refactor later
	}
	return
}
