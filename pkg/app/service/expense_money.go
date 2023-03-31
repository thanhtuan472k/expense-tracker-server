package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/external/util/pagetoken"
	"expense-tracker-server/external/util/ptime"
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
	expenseMoney, err := s.FindByID(ctx, expenseID)
	if err != nil {
		return
	}

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
	result = expenseMoney.ID.Hex()
	return
}

// All ...
func (s expenseMoneyImplement) All(ctx context.Context, q mgquerry.AppQuery, userID primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyAll, err error) {
	var (
		d    = dao.ExpenseMoney()
		cond = bson.D{
			{"user", userID},
		}
	)

	// Assign query
	s.assignQueryExpenseMoney(&q, &cond)

	// Assign sort
	s.assignQuerySort(&q)

	// Init data
	var list = make([]responsemodel.ResponseExpenseMoneyInfo, 0)

	// Find docs
	docs := d.FindByCondition(ctx, cond, q.GetFindOptionsWithPage())
	for _, doc := range docs {
		list = append(list, s.getExpenseMoneyInfo(ctx, doc))
	}

	// Page token
	total := len(list)
	endData := total < int(q.Limit)
	var nextPageToken = ""
	if total == int(q.Limit) {
		nextPageToken = pagetoken.PageTokenUsingPage(int(q.Page) + 1)
	}

	// Response
	result = responsemodel.ResponseExpenseMoneyAll{
		List:          list,
		EndData:       endData,
		NextPageToken: nextPageToken,
		Total:         int64(total),
	}
	return
}

// Detail ...
func (s expenseMoneyImplement) Detail(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseExpenseMoneyInfo, err error) {
	// Find expenseMoney
	expenseMoney, err := s.FindByID(ctx, id)
	if err != nil {
		return
	}

	// Response
	result = s.getExpenseMoneyInfo(ctx, expenseMoney)
	return
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

//
// PRIVATE METHODS
//

// assignQueryExpenseMoney ...
func (s expenseMoneyImplement) assignQueryExpenseMoney(q *mgquerry.AppQuery, cond *bson.D) {
	q.ExpenseTracker.AssignFromToAt(cond)
}

// assignQuerySort ...
func (s expenseMoneyImplement) assignQuerySort(q *mgquerry.AppQuery) {
	switch q.SortString {
	case "created_at_first":
		q.SortInterface = bson.D{
			{"createdAt", 1},
			{"_id", 1},
		}
	default:
		q.SortInterface = bson.D{
			{"createdAt", 1},
			{"id", 1},
		}
	}
}

// getExpenseMoneyInfo ...
func (s expenseMoneyImplement) getExpenseMoneyInfo(ctx context.Context, doc mgexpense.ExpenseMoney) responsemodel.ResponseExpenseMoneyInfo {
	return responsemodel.ResponseExpenseMoneyInfo{
		ID: doc.ID.Hex(),
		Category: mgexpense.CategoryShort{
			ID:   doc.Category.ID,
			Name: doc.Category.Name,
		},
		SubCategory: mgexpense.SubCategoryShort{
			ID:   doc.SubCategory.ID,
			Name: doc.SubCategory.Name,
		},
		Money:     doc.Money,
		Note:      doc.Note,
		CreatedAt: ptime.TimeResponseInit(doc.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(doc.UpdatedAt),
	}
}
