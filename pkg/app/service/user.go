package service

import (
	"context"
	"errors"
	"expense-tracker-server/pkg/app/dao"
	"expense-tracker-server/pkg/app/errorcode"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	responsemodel "expense-tracker-server/pkg/app/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserInterface ...
type UserInterface interface {
	// Register ...
	Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error)

	// Login ...
	Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error)

	// GetMe ...
	GetMe(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseUserMe, err error)
}

// User ...
func User() UserInterface {
	return userImplement{}
}

// userImplement ...
type userImplement struct{}

//
// PUBLIC METHODS
//

// Register ...
func (s userImplement) Register(ctx context.Context, payload requestmodel.UserBodyRegister) (result responsemodel.ResponseUserRegister, err error) {
	var (
		d   = dao.User()
		doc = payload.ConvertToBSON()
	)

	// Check phone and email existed
	if isExisted := s.isExistedPhoneOrEmailByUser(ctx, payload.Phone, payload.Email); isExisted {
		err = errors.New(errorcode.UserExistedPhoneOrEmail)
		return
	}

	// Insert doc
	if err = d.InsertOne(ctx, doc); err != nil {
		return
	}

	// Response
	result.ID = doc.ID.Hex()
	return
}

// Login ...
func (s userImplement) Login(ctx context.Context, payload requestmodel.UserBodyLogin) (result responsemodel.ResponseUserLogin, err error) {
	return
}

// GetMe ...
func (s userImplement) GetMe(ctx context.Context, id primitive.ObjectID) (result responsemodel.ResponseUserMe, err error) {
	return
}

// PRIVATE METHODS

// isExistedPhoneOrEmailByUser ...
func (s userImplement) isExistedPhoneOrEmailByUser(ctx context.Context, phone, email string) bool {
	var (
		d    = dao.User()
		cond = bson.M{
			"$or": []bson.M{
				{"phone": phone},
				{"email": email},
			},
		}
	)

	// Find user
	if user := d.FindOneByCondition(ctx, cond); !user.ID.IsZero() {
		return true
	}

	return false

}
