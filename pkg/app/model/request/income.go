package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/ptime"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeBodyCreate ...
type IncomeBodyCreate struct {
	Name     string  `json:"name"`
	Money    float64 `json:"money"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}

// Validate ...
func (m IncomeBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m IncomeBodyCreate) ConvertToBSON(userID primitive.ObjectID, category mgexpense.Category) mgexpense.IncomeMoney {
	return mgexpense.IncomeMoney{
		ID: primitive.NewObjectID(),
		Category: mgexpense.CategoryShort{
			ID:   category.ID,
			Name: category.Name,
		},
		User:      userID,
		Money:     m.Money,
		Status:    constant.StatusActive,
		Note:      m.Note,
		CreatedAt: ptime.Now(),
		UpdatedAt: ptime.Now(),
	}
}

// IncomeBodyUpdate ...
type IncomeBodyUpdate struct {
	Name     string  `json:"name"`
	Money    float64 `json:"money"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}

// Validate ...
func (m IncomeBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}
