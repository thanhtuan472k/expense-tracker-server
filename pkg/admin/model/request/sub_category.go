package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/ptime"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubCategoryBodyCreate ...
type SubCategoryBodyCreate struct {
	Name string `json:"name"`
}

// Validate ...
func (m SubCategoryBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m SubCategoryBodyCreate) ConvertToBSON(category mgexpense.Category) mgexpense.SubCategory {
	return mgexpense.SubCategory{
		ID:           primitive.NewObjectID(),
		Name:         m.Name,
		SearchString: mongodb.NonAccentVietnamese(m.Name),
		Status:       constant.StatusInactive,
		Type:         category.Type,
		Category:     category.ID,
		CreatedAt:    ptime.Now(),
		UpdatedAt:    ptime.Now(),
	}
}

// SubCategoryBodyUpdate ...
type SubCategoryBodyUpdate struct {
	Name string `json:"name"`
}

// Validate ...
func (m SubCategoryBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}

// SubCategoryChangeStatus ...
type SubCategoryChangeStatus struct {
	Status string `json:"status"`
}

// Validate ...
func (m SubCategoryChangeStatus) Validate() error {
	return validation.ValidateStruct(&m)
}

// SubCategoryAll ...
type SubCategoryAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
	Status  string `query:"status"`
}

// Validate ...
func (m SubCategoryAll) Validate() error {
	return validation.ValidateStruct(&m)
}
