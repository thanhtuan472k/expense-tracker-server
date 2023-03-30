package dao

import (
	"context"
	"expense-tracker-server/external/logger"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/internal/database"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExpenseMoneyInterface ...
type ExpenseMoneyInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)

	// FindOneByCondition ...
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.ExpenseMoney)

	// FindByCondition ...
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.ExpenseMoney)

	// CountByCondition ...
	CountByCondition(ctx context.Context, cond interface{}) int64

	// UpdateOneByCondition ...
	UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error)
}

// expenseMoneyImplement ...
type expenseMoneyImplement struct{}

// ExpenseMoney ...
func ExpenseMoney() ExpenseMoneyInterface {
	return expenseMoneyImplement{}
}

// InsertOne ...
func (i expenseMoneyImplement) InsertOne(ctx context.Context, payload interface{}) (err error) {
	var (
		col = database.ExpenseMoneyCol()
	)

	if _, err := col.InsertOne(ctx, payload); err != nil {
		logger.Error("dao.ExpenseMoney - InsertOne", logger.LogData{
			Data: bson.M{
				"payload": payload,
				"error":   err.Error(),
			},
		})
	}

	return
}

// FindOneByCondition ...
func (i expenseMoneyImplement) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.ExpenseMoney) {
	var col = database.IncomeMoneyCol()
	col.FindOne(ctx, cond, opts...).Decode(&doc)
	return
}

// FindByCondition ...
func (i expenseMoneyImplement) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.ExpenseMoney) {
	var col = database.IncomeMoneyCol()

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		fmt.Println("err FindByCondition 1: ", err.Error())
		return
	}

	// Close cursor
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		fmt.Println("err FindByCondition 2: ", err.Error())
	}
	return
}

// CountByCondition ...
func (i expenseMoneyImplement) CountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.ExpenseMoneyCol()
	)

	total, _ := col.CountDocuments(ctx, cond)
	return total
}

// UpdateOneByCondition ...
func (i expenseMoneyImplement) UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error) {
	var col = database.ExpenseMoneyCol()

	_, err = col.UpdateOne(ctx, cond, payload)
	return
}
