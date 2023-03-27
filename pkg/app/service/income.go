package service

import (
	"context"
	"errors"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeInterface ...
type IncomeInterface interface {
	// Create ...
	Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyCreate) (incomeID string, err error)

	// Update ...
	Update(ctx context.Context, userID, incomeID primitive.ObjectID, payload requestmodel.IncomeBodyUpdate) (result string, err error)
}

// Income ...
func Income() IncomeInterface {
	return incomeImplement{}
}

// incomeImplement ...
type incomeImplement struct{}

//
// PUBLIC METHODS
//

// Create ...
func (s incomeImplement) Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyCreate) (incomeID string, err error) {
	// Convert payload string to ObjectId
	categoryID, _ := mongodb.NewIDFromString(payload.Category)

	// Find doc category is available
	categorySvc := categoryImplement{}
	category, err := categorySvc.findDocCategoryTypeIncomeAvailableByID(ctx, categoryID)
	if err != nil {
		return
	}

	var (
		d   = dao.IncomeMoney()
		doc = payload.ConvertToBSON(userID, category)
	)

	// Insert doc income money
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	incomeID = doc.ID.Hex()
	return
}

// Update ...
func (s incomeImplement) Update(ctx context.Context, userID, incomeID primitive.ObjectID, payload requestmodel.IncomeBodyUpdate) (result string, err error) {
	// Find incomeMoney
	incomeMoney, err := s.FindByID(ctx, incomeID)
	if err != nil {
		return
	}

	// Convert payload string to ObjectId
	categoryID, _ := mongodb.NewIDFromString(payload.Category)

	// Find doc category is available
	categorySvc := categoryImplement{}
	category, err := categorySvc.findDocCategoryTypeIncomeAvailableByID(ctx, categoryID)
	if err != nil {
		return
	}

	fmt.Println("categoryDoc", category)

	var (
		d             = dao.IncomeMoney()
		data          = payload.ConvertToBSON(category)
		payloadUpdate = bson.M{
			"money":     data.Money,
			"note":      data.Note,
			"category":  data.Category,
			"updatedAt": data.UpdatedAt,
		}
		cond = bson.D{{"_id", incomeMoney.ID}}
	)

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	result = incomeMoney.ID.Hex()
	return
}

// FindByID ...
func (s incomeImplement) FindByID(ctx context.Context, id primitive.ObjectID) (result mgexpense.IncomeMoney, err error) {
	var (
		d    = dao.IncomeMoney()
		cond = bson.D{{"_id", id}}
	)

	incomeMoney := d.FindOneByCondition(ctx, cond)
	if incomeMoney.ID.IsZero() {
		err = errors.New(errorcode.IncomeMoneyNotFound)
		return
	}
	result = incomeMoney
	return
}

//
// PRIVATE METHODS
