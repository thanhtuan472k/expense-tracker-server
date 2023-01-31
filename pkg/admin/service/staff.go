package service

import (
	"context"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
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
func (s staffImplement) Login(ctx context.Context, payload requestmodel.StaffBodyLogin) (success responsemodel.ResponseLoginSuccess, err error) {
	// Check staff (email, phone) is existed in system or not

	// If staff is not existed --> User not found

	// If staff existed
	// - payload.Password (hashed) and compare with hasedPassword in DB
	// - If wrong password --> Password is incorrect
	// - If success password --> Generate token and send response

	// Return

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
