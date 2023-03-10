package requestmodel

import (
	"context"
	"expense-tracker-server/external/constant"
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/internal/auth"
	"expense-tracker-server/pkg/app/dao"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
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
func (m UserBodyRegister) ConvertToBSON(ctx context.Context) mgexpense.User {
	return mgexpense.User{
		ID:           primitive.NewObjectID(),
		Name:         m.Name,
		SearchString: m.Name,
		Email:        m.Email,
		Gender:       m.Gender,
		Phone:        m.Phone,
		Password:     auth.GenerateHashedPassword(m.Password),
		Status:       constant.StatusActive,
		Code:         GenerateRandomUniqueUserCode(ctx),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GenerateRandomUniqueUserCode ...
func GenerateRandomUniqueUserCode(ctx context.Context) (result string) {
	randomString := funk.RandomString(5)
	result = "U" + "_" + randomString

	var d = dao.User()

	// Find user by code
	user := d.FindOneByCondition(ctx, bson.M{
		"code": result,
	})

	if !user.ID.IsZero() {
		return GenerateRandomUniqueUserCode(ctx)
	}

	return
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
