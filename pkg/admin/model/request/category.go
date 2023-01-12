package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/ptime"
	internalconstant "expense-tracker-server/internal/constant"
	"expense-tracker-server/pkg/admin/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
		validation.Field(&m.Type, validation.In(types...).Error(errorcode.CategoryTypeIsInvalid)),
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

// CategoryAll ...
type CategoryAll struct {
	Page    int64  `json:"page"`
	Limit   int64  `json:"limit"`
	Keyword string `json:"keyword"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

// Validate ...
func (m CategoryAll) Validate() error {
	var (
		types = []interface{}{
			internalconstant.CategoryTypeExpense,
			internalconstant.CategoryTypeIncome,
		}

		statuses = []interface{}{
			constant.StatusActive,
			constant.StatusInactive,
		}
	)

	return validation.ValidateStruct(&m,
		validation.Field(&m.Page, is.UTFLetterNumeric.Error(response.CommonPageInvalid)),
		validation.Field(&m.Limit, is.UTFLetterNumeric.Error(response.CommonLimitInvalid)),
		validation.Field(&m.Type, validation.In(types...).Error(errorcode.CategoryIsInvalid)),
		validation.Field(&m.Status, validation.In(statuses...).Error(errorcode.CategoryStatusIsInvalid)),
	)
}

// CategoryChangeStatus ...
type CategoryChangeStatus struct {
	Status string `json:"status"`
}

// Validate ...
func (m CategoryChangeStatus) Validate() error {
	var statuses = []interface{}{
		constant.StatusActive,
		constant.StatusInactive,
	}

	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.In(statuses...).Error(errorcode.CategoryStatusIsInvalid)),
	)

}
