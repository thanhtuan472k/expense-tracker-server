package service

import (
	"context"
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
	return
}

// Update ...
func (s incomeImplement) Update(ctx context.Context, userID primitive.ObjectID, payload requestmodel.IncomeBodyUpdate) (incomeID string, err error) {
	return
}
