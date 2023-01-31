package responsemodel

// ResponseLoginSuccess ...
type ResponseLoginSuccess struct {
	ID    string `json:"_id"`
	Token string `json:"token"`
}

// ResponseStaffMe ...
type ResponseStaffMe struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
}
