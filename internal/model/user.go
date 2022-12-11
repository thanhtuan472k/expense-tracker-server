package model

import "time"

// UserRaw ...
type UserRaw struct {
	ID           AppID     `bson:"_id"`
	Name         string    `bson:"name"`
	SearchString string    `bson:"searchString"`
	Gender       string    `bson:"gender"`
	Banned       bool      `bson:"banned"`
	Phone        string    `bson:"phone"`
	Code         []string  `bson:"code"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
