package dao

import (
	"context"
	"expense-tracker-server/external/logger"
	"expense-tracker-server/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoryInterface interface {
	// InsertOne ...
	InsertOne(ctx context.Context, payload interface{}) (err error)
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
