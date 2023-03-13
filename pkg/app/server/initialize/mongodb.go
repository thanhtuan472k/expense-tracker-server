package initialize

import "expense-tracker-server/internal/database"

func mongoDB() {
	database.ConnectMongoDBExpense()
}
