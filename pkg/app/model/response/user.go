package responsemodel

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
}
