package errorcode

import "expense-tracker-server/external/response"

func Init() {
	// Init common code first
	response.Init()

	// Code from 200 -> 299
	response.AddListCodes(category)
}
