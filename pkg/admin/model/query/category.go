package query

import (
	"expense-tracker-server/external/constant"
	"expense-tracker-server/external/response"
	internalconstant "expense-tracker-server/internal/constant"
	"expense-tracker-server/pkg/admin/errorcode"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CategoryAll ...
type CategoryAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
	Status  string `query:"status"`
	Type    string `query:"type"`
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
		validation.Field(&m.Page, validation.Min(0).Error(response.CommonPageInvalid)),
		validation.Field(&m.Limit, validation.Min(0).Error(response.CommonLimitInvalid)),
		validation.Field(&m.Type, validation.In(types...).Error(errorcode.CategoryTypeIsInvalid)),
		validation.Field(&m.Status, validation.In(statuses...).Error(errorcode.CategoryStatusIsInvalid)),
	)
}
