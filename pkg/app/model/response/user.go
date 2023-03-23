package responsemodel

import "expense-tracker-server/external/util/ptime"

// ResponseUserRegister ...
type ResponseUserRegister struct {
	ID string `json:"_id"`
}

// ResponseUserLogin ...
type ResponseUserLogin struct {
	ID          string `json:"_id"`
	AccessToken string `json:"accessToken"`
}

// ResponseUserMe ...
type ResponseUserMe struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Gender    string              `json:"gender"`
	Phone     string              `json:"phone"`
	Code      string              `json:"code"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
	UpdatedAt *ptime.TimeResponse `json:"updatedAt"`
}
