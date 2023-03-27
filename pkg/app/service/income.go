package service

import (
	"context"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/pkg/app/dao"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeInterface ...
type IncomeInterface interface {
	// Create ...
	Create(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyCreate) (incomeID string, err error)

	// Update ...
	Update(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyUpdate) (incomeID string, err error)
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
func (s incomeImplement) Update(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyUpdate) (incomeID string, err error) {
	return
}

//
// PRIVATE METHODS
