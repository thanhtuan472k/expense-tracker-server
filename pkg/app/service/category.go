package service

import (
	"context"
	"expense-tracker-server/external/util/mgquerry"
	responsemodel "expense-tracker-server/pkg/app/model/response"
)

// CategoryInterface ...
type CategoryInterface interface {
	// All ...
	All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseUserRegister, err error)
}

// Category ...
func Category() CategoryInterface {
	return categoryImplement{}
}

// categoryImplement ...
type categoryImplement struct{}

//
// PUBLIC METHOD
//

// All ...
func (c categoryImplement) All(ctx context.Context, q mgquerry.AppQuery) (result responsemodel.ResponseUserRegister, err error) {
	//TODO implement me
	panic("implement me")
}
