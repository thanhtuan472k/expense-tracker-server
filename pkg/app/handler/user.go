package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"expense-tracker-server/pkg/app/service"
	"github.com/labstack/echo/v4"
)

// User ...
type User struct{}

// Register godoc
// @tags User
// @summary Login
// @id user-register
// @accept json
// @produce json
// @param payload body requestmodel.UserBodyRegister true "Payload"
// @success 200 {object} nil
// @router /users/register [post]
func (User) Register(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.UserBodyRegister)
		s       = service.User()
	)

	result, err := s.Register(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}

// Login godoc
// @tags User
// @summary Login
// @id user-login
// @accept json
// @produce json
// @param payload body requestmodel.UserBodyLogin true "Payload"
// @success 200 {object} nil
// @router /users/login [post]
func (User) Login(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.UserBodyLogin)
		s       = service.User()
	)

	result, err := s.Login(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}

// GetMe godoc
// @tags Staff
// @summary GetMe
// @id user-get-me
// @security ApiKeyAuth
// @accept json
// @produce json
// @success 200 {object} nil
// @router /users/me [get]
func (User) GetMe(c echo.Context) error {
	var (
		ctx    = echocontext.GetContext(c)
		userID = echocontext.GetCurrentUserID(c)
		s      = service.User()
	)

	result, err := s.GetMe(ctx, userID)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}
