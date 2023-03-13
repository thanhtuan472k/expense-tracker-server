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

// UserInterface ...
type UserInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)

	// FindOneByCondition ...
	FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.User)

	// FindByCondition ...
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.User)

	// CountByCondition ...
	CountByCondition(ctx context.Context, cond interface{}) int64

	// UpdateOneByCondition ...
	UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error)
}

// userImplement ...
type userImplement struct{}

// User ...
func User() UserInterface {
	return userImplement{}
}

// InsertOne ...
func (d userImplement) InsertOne(ctx context.Context, payload interface{}) (err error) {
	var (
		col = database.UserCol()
	)

	if _, err := col.InsertOne(ctx, payload); err != nil {
		logger.Error("dao.User - InsertOne", logger.LogData{
			Data: bson.M{
				"payload": payload,
				"error":   err.Error(),
			},
		})
	}

	return
}

// FindOneByCondition ...
func (d userImplement) FindOneByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOneOptions) (doc mgexpense.User) {
	var col = database.UserCol()
	col.FindOne(ctx, cond, opts...).Decode(&doc)
	return
}

// FindByCondition ...
func (d userImplement) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []mgexpense.User) {
	var col = database.UserCol()

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
func (d userImplement) CountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.UserCol()
	)

	total, _ := col.CountDocuments(ctx, cond)
	return total
}

// UpdateOneByCondition ...
func (d userImplement) UpdateOneByCondition(ctx context.Context, cond, payload interface{}) (err error) {
	var col = database.UserCol()

	_, err = col.UpdateOne(ctx, cond, payload)
	return
}
