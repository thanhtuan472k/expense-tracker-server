package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson"
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
	// Find expenseMoney
	//expenseMoney, err := s.FindByID(ctx, expenseID)
	//if err != nil {
	//	return
	//}

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
		d             = dao.ExpenseMoney()
		data          = payload.ConvertToBSON(category, subCategory)
		payloadUpdate = bson.M{
			"money":       data.Money,
			"category":    data.Category,
			"subCategory": data.SubCategory,
			"note":        data.Note,
			"updatedAt":   data.UpdatedAt,
		}
		cond = bson.D{{"_id", expenseID}}
	)

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	result = expenseID.Hex()
	return
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

// FindByID ...
func (s expenseMoneyImplement) FindByID(ctx context.Context, id primitive.ObjectID) (result mgexpense.ExpenseMoney, err error) {
	var (
		d    = dao.ExpenseMoney()
		cond = bson.D{{"_id", id}}
	)

	expenseMoney := d.FindOneByCondition(ctx, cond)
	if expenseMoney.ID.IsZero() {
		err = errors.New(errorcode.ExpenseMoneyNotFound)
		return
	}
	result = expenseMoney
	return
}
