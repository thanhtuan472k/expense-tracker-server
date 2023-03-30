package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/ptime"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExpenseMoneyBodyCreate ...
type ExpenseMoneyBodyCreate struct {
	Money       float64 `json:"money"`
	Category    string  `json:"category"`
	SubCategory string  `json:"subCategory"`
	Note        string  `json:"note"`
}

// Validate ...
func (m ExpenseMoneyBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m ExpenseMoneyBodyCreate) ConvertToBSON(userID primitive.ObjectID, category mgexpense.Category, subCategory mgexpense.SubCategory) mgexpense.ExpenseMoney {
	return mgexpense.ExpenseMoney{
		ID: primitive.NewObjectID(),
		Category: mgexpense.CategoryShort{
			ID:   category.ID,
			Name: category.Name,
		},
		SubCategory: mgexpense.SubCategoryShort{
			ID:   subCategory.ID,
			Name: subCategory.Name,
		},
		User:      userID,
		Money:     m.Money,
		Status:    constant.StatusActive,
		Note:      m.Note,
		CreatedAt: ptime.Now(),
		UpdatedAt: ptime.Now(),
	}
}

// ExpenseMoneyBodyUpdate ...
type ExpenseMoneyBodyUpdate struct {
}

// Validate ...
func (m ExpenseMoneyBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}
