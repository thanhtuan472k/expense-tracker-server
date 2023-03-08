package initialize

import (
	"expense-tracker-server/internal/config"
	"expense-tracker-server/internal/redis"
	"expense-tracker-server/internal/zk"
	"expense-tracker-server/pkg/app/errorcode"
)

// Init ...
func Init() {
	// Config ...
	config.Init()

	// Error code locale ...
	errorcode.Init()

	// Zookeeper connect ...
	zk.Connect()

	// Mongo db connect ...
	mongoDB()

	// Redis connect ...
	redis.InitRedis()
}
