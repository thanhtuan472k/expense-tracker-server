package errorcode

import "expense-tracker-server/external/response"

func Init() {
	// Init common code first
	response.Init()

	// Code from 200 -> 299
	response.AddListCodes(staff)

	// Code from 300 -> 399
	response.AddListCodes(category)

	// Code from 400 -> 499
	response.AddListCodes(subCategory)
}
