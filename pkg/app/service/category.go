package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryInterface ...
type CategoryInterface interface {
	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseUserRegister, err error)
}

// Category ...
func Category() CategoryInterface {
	return categoryImplement{}
}

// categoryImplement ...
type categoryImplement struct{}

//
// PUBLIC METHOD
//

// All ...
func (s categoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseUserRegister, err error) {
	//TODO implement me
	panic("implement me")
}

//
// PRIVATE METHODS
//

// findDocCategoryTypeIncomeAvailableByID ...
func (s categoryImplement) findDocCategoryTypeIncomeAvailableByID(ctx context.Context, id primitive.ObjectID) (doc mgexpense.Category, err error) {
	var (
		d    = dao.Category()
		cond = bson.D{{"_id", id}}
	)

	doc = d.FindOneByCondition(ctx, cond)
	if !doc.IsCategoryAvailableByTypeIncome() {
		err = errors.New(errorcode.CategoryIsInvalid)
	}
	return
}

// findDocCategoryTypeExpenseAvailableByID ...
func (s categoryImplement) findDocCategoryTypeExpenseAvailableByID(ctx context.Context, id primitive.ObjectID) (doc mgexpense.Category, err error) {
	var (
		d    = dao.Category()
		cond = bson.D{{"_id", id}}
	)

	doc = d.FindOneByCondition(ctx, cond)
	if !doc.IsCategoryAvailableByTypeExpense() {
		err = errors.New(errorcode.CategoryIsInvalid)
	}
	return
}
