package errorcode

import "expense-tracker-server/external/response"

// Init ...
func Init() {
	// Init common code first
	response.Init()

	// Code from 200 -> 299
	response.AddListCodes(user)
}
