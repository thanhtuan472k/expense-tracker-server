package service

import (
	"context"
	"errors"
	"expense-tracker-server/internal/auth"
	"expense-tracker-server/pkg/admin/dao"
	"expense-tracker-server/pkg/admin/errorcode"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StaffInterface ...
type StaffInterface interface {
	// Login ...
	Login(ctx context.Context, payload requestmodel.StaffBodyLogin) (success responsemodel.ResponseLoginSuccess, err error)

	// GetMe ...
	GetMe(ctx context.Context, staffID primitive.ObjectID) (result responsemodel.ResponseStaffMe, err error)

	// Update ...
	Update(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyUpdate) (result responsemodel.ResponseUpdate, err error)

	// ChangePassword ...
	ChangePassword(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyChangePassword) (result responsemodel.ResponseUpdate, err error)
}

// Staff ...
func Staff() StaffInterface {
	return staffImplement{}
}

// staffImplement ...
type staffImplement struct{}

//
// PUBLIC METHODS ...
//

// Login ...
func (s staffImplement) Login(ctx context.Context, payload requestmodel.StaffBodyLogin) (result responsemodel.ResponseLoginSuccess, err error) {
	var (
		d = dao.Staff()
	)
	// Check phone is existed in system or not
	staff := d.FindOneByCondition(ctx, bson.M{"phone": payload.Phone})

	// If staff is not existed --> User not found
	if staff.ID.IsZero() {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	// If staff existed in DB
	if isValidPassword := auth.CompareHashedPassword(payload.Password, staff.Password); !isValidPassword {
		err = errors.New(errorcode.StaffPasswordIncorrect)
		return
	}
	// - If success password --> Generate token and send response
	accessToken := staff.GenerateAccessToken()
	//refreshToken := staff.GenerateRefreshToken()
	result = responsemodel.ResponseLoginSuccess{
		ID:          staff.ID.Hex(),
		AccessToken: accessToken,
		//RefreshToken: refreshToken,
	}

	// Response
	return
}

// GetMe ...
func (s staffImplement) GetMe(ctx context.Context, staffID primitive.ObjectID) (result responsemodel.ResponseStaffMe, err error) {
	//TODO implement me
	panic("implement me")
}

// Update ...
func (s staffImplement) Update(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyUpdate) (result responsemodel.ResponseUpdate, err error) {
	//TODO implement me
	panic("implement me")
}

// ChangePassword ...
func (s staffImplement) ChangePassword(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyChangePassword) (result responsemodel.ResponseUpdate, err error) {
	//TODO implement me
	panic("implement me")
}

//
// PRIVATE METHODS ...
//
