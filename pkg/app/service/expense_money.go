package service

import (
	"context"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/pkg/app/dao"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseMoneyInterface interface {
	// Create ...
	Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.ExpenseMoneyBodyCreate) (expenseID string, err error)

	// Update ...
	Update(ctx context.Context, userID, expenseID primitive.ObjectID, payload requestmodel.ExpenseMoneyBodyUpdate) (result string, err error)

	// All ...
	All(ctx context.Context, q mgquerry.AppQuery, userID primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyAll, err error)

	// Detail ...
	Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyInfo, err error)
}

// ExpenseMoney ...
func ExpenseMoney() ExpenseMoneyInterface {
	return expenseMoneyImplement{}
}

// expenseMoneyImplement ...
type expenseMoneyImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s expenseMoneyImplement) Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.ExpenseMoneyBodyCreate) (expenseID string, err error) {
	// Convert payload string to ObjectId
	categoryID, _ := mongodb.NewIDFromString(payload.Category)
	subCategoryID, _ := mongodb.NewIDFromString(payload.SubCategory)

	// Find doc category is available
	categorySvc := categoryImplement{}
	category, err := categorySvc.findDocCategoryTypeExpenseAvailableByID(ctx, categoryID)
	if err != nil {
		return
	}

	// Find doc subCategory is available
	subCategorySvc := subCategoryImplement{}
	subCategory, err := subCategorySvc.findDocSubCategoryTypeExpenseAvailableByID(ctx, subCategoryID)
	if err != nil {
		return
	}

	var (
		d   = dao.ExpenseMoney()
		doc = payload.ConvertToBSON(userID, category, subCategory)
	)

	// Insert doc expense money
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	expenseID = doc.ID.Hex()
	return
}

// Update ...
func (s expenseMoneyImplement) Update(ctx context.Context, userID, expenseID primitive.ObjectID, payload requestmodel.ExpenseMoneyBodyUpdate) (result string, err error) {
	//TODO implement me
	panic("implement me")
}

// All ...
func (s expenseMoneyImplement) All(ctx context.Context, q mgquerry.AppQuery, userID primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyAll, err error) {
	//TODO implement me
	panic("implement me")
}

// Detail ...
func (s expenseMoneyImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyInfo, err error) {
	//TODO implement me
	panic("implement me")
}
