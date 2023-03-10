package service

import (
	"context"
	"errors"
	"expense-tracker-server/external/constant"
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/external/util/format"
	"expense-tracker-server/external/util/ptime"
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
	staff := d.FindOneByCondition(ctx, bson.M{"phone": format.PhoneFormatCommon(payload.Phone)})

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
	// - If success password --> Generate token
	accessToken := staff.GenerateAccessToken()
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
	var (
		d = dao.Staff()
	)

	// Find staff
	staff := d.FindOneByCondition(ctx, bson.M{"_id": staffID})

	if staff.ID.IsZero() {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	if staff.Status == constant.StatusInactive {
		err = errors.New(errorcode.StaffStatusInactive)
		return
	}

	// Response
	result = responsemodel.ResponseStaffMe{
		ID:     staff.ID.Hex(),
		Name:   staff.Name,
		Email:  staff.Email,
		Gender: staff.Gender,
		Phone:  format.PhoneFormatCommon(staff.Phone),
	}

	return
}

// Update ...
func (s staffImplement) Update(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyUpdate) (result responsemodel.ResponseUpdate, err error) {
	var (
		d             = dao.Staff()
		payloadUpdate = bson.M{
			"name":         payload.Name,
			"searchString": mongodb.NonAccentVietnamese(payload.Name),
			"phone":        format.PhoneFormatCommon(payload.Phone),
			"email":        payload.Email,
			"gender":       payload.Gender,
			"updatedAt":    ptime.Now(),
		}
		cond = bson.M{"_id": staffID}
	)

	// Find staff
	staff := d.FindOneByCondition(ctx, bson.M{"_id": staffID})

	if staff.ID.IsZero() {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	// Check staff status inactive
	if staff.Status == constant.StatusInactive {
		err = errors.New(errorcode.StaffStatusInactive)
		return
	}

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	result = responsemodel.ResponseUpdate{ID: staff.ID.Hex()}
	return
}

// ChangePassword ...
func (s staffImplement) ChangePassword(ctx context.Context, staffID primitive.ObjectID, payload requestmodel.StaffBodyChangePassword) (result responsemodel.ResponseUpdate, err error) {
	var (
		d             = dao.Staff()
		payloadUpdate = bson.M{
			"password":  auth.GenerateHashedPassword(payload.NewPassword),
			"updatedAt": ptime.Now(),
		}
		cond = bson.M{"_id": staffID}
	)

	// Find staff
	staff := d.FindOneByCondition(ctx, cond)
	if staff.ID.IsZero() {
		err = errors.New(errorcode.StaffNotFound)
		return
	}

	// Check oldPassword is correct
	if isValidOldPassword := auth.CompareHashedPassword(payload.OldPassword, staff.Password); !isValidOldPassword {
		err = errors.New(errorcode.StaffWrongOldPassword)
		return
	}

	// Compare oldPassword vs newPassword
	if isComparePassword := auth.CompareHashedPassword(payload.NewPassword, staff.Password); isComparePassword {
		err = errors.New(errorcode.StaffNewPasswordSameOldPassword)
		return
	}

	// Update
	if err = d.UpdateOneByCondition(ctx, cond, bson.M{"$set": payloadUpdate}); err != nil {
		return
	}

	// Response
	result = responsemodel.ResponseUpdate{
		ID: staffID.Hex(),
	}
	return
}

//
// PRIVATE METHODS ...
//
