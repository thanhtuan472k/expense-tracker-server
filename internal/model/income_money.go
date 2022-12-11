package model

import "time"

// IncomeMoney ...
type IncomeMoney struct {
	ID        AppID     `bson:"_id"`
	Category  AppID     `bson:"category"` // get list category have type = income
	Amount    float64   `bson:"amount"`
	Status    string    `bson:"status"`
	Note      string    `bson:"note"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
