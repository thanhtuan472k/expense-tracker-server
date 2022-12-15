package initialize

import (
	"expense-tracker-server/internal/config"
	"expense-tracker-server/internal/redis"
	"expense-tracker-server/internal/zk"
	"expense-tracker-server/pkg/admin/errorcode"
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

	// Authentication ...
	//authentication()

	// Audit ...
	//InitAudit()

	// Nats connect ...
	//nats.Connect()

	// Redis connect ...
	redis.InitRedis()
}
