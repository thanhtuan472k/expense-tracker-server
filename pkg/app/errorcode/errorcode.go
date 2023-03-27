package errorcode

import "expense-tracker-server/external/response"

// Init ...
func Init() {
	// Init common code first
	response.Init()

	// Code from 200 -> 299
	response.AddListCodes(user)

	// Code from 300 -> 399
	response.AddListCodes(category)

	// Code from 400 -> 499
	response.AddListCodes(incomeMoney)
}
