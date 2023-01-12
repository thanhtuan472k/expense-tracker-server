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

type CategoryInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)

	// FindOneByCondition ...
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.Category)

	// FindByCondition ...
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.Category)

	// CountByCondition ...
	CountByCondition(ctx context.Context, cond interface{}) int64
}

// categoryImplement ...
type categoryImplement struct{}

// Category ...
func Category() CategoryInterface {
	return categoryImplement{}
}

// InsertOne ...
func (d categoryImplement) InsertOne(ctx context.Context, payload interface{}) (err error) {
	var (
		col = database.CategoryCol()
	)

	if _, err := col.InsertOne(ctx, payload); err != nil {
		logger.Error("dao.Category - InsertOne", logger.LogData{
			Data: bson.M{
				"payload": payload,
				"error":   err.Error(),
			},
		})
	}

	return
}

// FindOneByCondition ...
func (d categoryImplement) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.Category) {
	var col = database.CategoryCol()

	if err := col.FindOne(ctx, cond, opts...).Decode(&doc); err != nil {
		logger.Error("dao.Category - FindOneByCondition err", logger.LogData{
			Data: bson.M{
				"cond":  cond,
				"opts":  opts,
				"error": err.Error(),
			},
		})
	}
	return
}

// FindByCondition ...
func (d categoryImplement) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.Category) {
	var col = database.CategoryCol()

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
func (d categoryImplement) CountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.CategoryCol()
	)

	total, _ := col.CountDocuments(ctx, cond)
	return total
}
