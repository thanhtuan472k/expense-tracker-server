package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/ptime"
	internalconstant "expense-tracker-server/internal/constant"
	"expense-tracker-server/pkg/admin/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryBodyCreate ...
type CategoryBodyCreate struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Validate ...
func (m CategoryBodyCreate) Validate() error {
	var types = []interface{}{
		internalconstant.CategoryTypeExpense,
		internalconstant.CategoryTypeIncome,
	}
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error(errorcode.CategoryIsRequiredName)),
		validation.Field(&m.Type, validation.In(types...).Error(errorcode.CategoryIsInvalid)),
	)
}

// ConvertToBSON ...
func (m CategoryBodyCreate) ConvertToBSON() mgexpense.Category {
	return mgexpense.Category{
		ID:           primitive.NewObjectID(),
		Name:         m.Name,
		SearchString: mongodb.NonAccentVietnamese(m.Name),
		Status:       constant.StatusInactive,
		Type:         m.Type,
		CreatedAt:    ptime.Now(),
		UpdatedAt:    ptime.Now(),
	}
}

// CategoryChangeStatus ...
type CategoryChangeStatus struct {
	Status string `json:"status"`
}
