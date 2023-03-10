package requestmodel

import (
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//
// Register
//

// UserBodyRegister ...
type UserBodyRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Validate ...
func (m UserBodyRegister) Validate() error {
	return validation.ValidateStruct(&m)
}

// ConvertToBSON ...
func (m UserBodyRegister) ConvertToBSON() mgexpense.User {
	return mgexpense.User{
		ID:           primitive.NewObjectID(),
		Name:         m.Name,
		SearchString: m.Name,
		Email:        m.Email,
		Gender:       m.Gender,
		Phone:        m.Phone,
		Status:       constant.StatusActive,
		Code:         "tuan",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

//
// Login
//

// UserBodyLogin ...
type UserBodyLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Validate ...
func (m UserBodyLogin) Validate() error {
	return validation.ValidateStruct(&m)
}
