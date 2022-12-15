package database

import (
	"expense-tracker-server/external/mongodb"
	"expense-tracker-server/internal/config"
	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var expenseDB *mongo.Database

// ConnectMongoDBExpense ...
func ConnectMongoDBExpense() {
	var (
		cfg = config.GetENV().MongoDB
		err error
		tls *mongodb.ConnectTLSOpts
	)

	// Connect
	expenseDB, err = mongodb.Connect(mongodb.Config{
		Host:       cfg.URI,
		DBName:     cfg.DBName,
		Monitor:    apmmongo.CommandMonitor(),
		TLS:        tls,
		Standalone: &mongodb.ConnectStandaloneOpts{},
	})
	if err != nil {
		panic(err)
	}

	// Index ...
	index()
}

// UserCol ...
func UserCol() *mongo.Collection {
	return expenseDB.Collection(colUser)
}

// CategoryCol ...
func CategoryCol() *mongo.Collection {
	return expenseDB.Collection(colCategory)
}

//ExpenseMoneyCol ...
func ExpenseMoneyCol() *mongo.Collection {
	return expenseDB.Collection(colExpenseMoney)
}

//IncomeMoneyCol ...
func IncomeMoneyCol() *mongo.Collection {
	return expenseDB.Collection(colIncomeMoney)
}
