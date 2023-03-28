package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/ptime"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeMoneyBodyCreate ...
type IncomeMoneyBodyCreate struct {
	Money    float64 `json:"money"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}

// Validate ...
func (m IncomeMoneyBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m IncomeMoneyBodyCreate) ConvertToBSON(userID primitive.ObjectID, category mgexpense.Category) mgexpense.IncomeMoney {
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

// IncomeMoneyBodyUpdate ...
type IncomeMoneyBodyUpdate struct {
	Money    float64 `json:"money"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}

// Validate ...
func (m IncomeMoneyBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m IncomeMoneyBodyUpdate) ConvertToBSON(category mgexpense.Category) mgexpense.IncomeMoney {
	return mgexpense.IncomeMoney{
		Category: mgexpense.CategoryShort{
			ID:   category.ID,
			Name: category.Name,
		},
		Money:     m.Money,
		Note:      m.Note,
		UpdatedAt: ptime.Now(),
	}
}
