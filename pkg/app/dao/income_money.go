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

// IncomeMoneyInterface ...
type IncomeMoneyInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)

	// FindOneByCondition ...
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.IncomeMoney)

	// FindByCondition ...
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.IncomeMoney)

	// CountByCondition ...
	CountByCondition(ctx context.Context, cond interface{}) int64

	// UpdateOneByCondition ...
	UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error)
}

// incomeMoneyImplement ...
type incomeMoneyImplement struct{}

// IncomeMoney ...
func IncomeMoney() IncomeMoneyInterface {
	return incomeMoneyImplement{}
}

// InsertOne ...
func (i incomeMoneyImplement) InsertOne(ctx context.Context, payload interface{}) (err error) {
	var (
		col = database.IncomeMoneyCol()
	)

	if _, err := col.InsertOne(ctx, payload); err != nil {
		logger.Error("dao.IncomeMoney - InsertOne", logger.LogData{
			Data: bson.M{
				"payload": payload,
				"error":   err.Error(),
			},
		})
	}

	return
}

// FindOneByCondition ...
func (i incomeMoneyImplement) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.IncomeMoney) {
	var col = database.IncomeMoneyCol()
	col.FindOne(ctx, cond, opts...).Decode(&doc)
	return
}

// FindByCondition ...
func (i incomeMoneyImplement) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.IncomeMoney) {
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
func (i incomeMoneyImplement) CountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.IncomeMoneyCol()
	)

	total, _ := col.CountDocuments(ctx, cond)
	return total
}

// UpdateOneByCondition ...
func (i incomeMoneyImplement) UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error) {
	var col = database.IncomeMoneyCol()

	_, err = col.UpdateOne(ctx, cond, payload)
	return
}
