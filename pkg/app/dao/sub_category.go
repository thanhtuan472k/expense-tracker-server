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

// SubCategoryInterface ...
type SubCategoryInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)

	// FindOneByCondition ...
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.SubCategory)

	// FindByCondition ...
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.SubCategory)

	// CountByCondition ...
	CountByCondition(ctx context.Context, cond interface{}) int64

	// UpdateOneByCondition ...
	UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error)
}

// subCategoryImplement ...
type subCategoryImplement struct{}

// SubCategory ...
func SubCategory() SubCategoryInterface {
	return subCategoryImplement{}
}

// InsertOne ...
func (d subCategoryImplement) InsertOne(ctx context.Context, payload interface{}) (err error) {
	var (
		col = database.SubCategoryCol()
	)

	if _, err := col.InsertOne(ctx, payload); err != nil {
		logger.Error("dao.SubCategory - InsertOne", logger.LogData{
			Data: bson.M{
				"payload": payload,
				"error":   err.Error(),
			},
		})
	}

	return
}

// FindOneByCondition ...
func (d subCategoryImplement) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.SubCategory) {
	var col = database.SubCategoryCol()

	if err := col.FindOne(ctx, cond, opts...).Decode(&doc); err != nil {
		logger.Error("dao.SubCategory - FindOneByCondition err", logger.LogData{
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
func (d subCategoryImplement) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.SubCategory) {
	var col = database.SubCategoryCol()

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
func (d subCategoryImplement) CountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.SubCategoryCol()
	)

	total, _ := col.CountDocuments(ctx, cond)
	return total
}

// UpdateOneByCondition ...
func (d subCategoryImplement) UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error) {
	var col = database.SubCategoryCol()

	_, err = col.UpdateOne(ctx, cond, payload)
	return
}
