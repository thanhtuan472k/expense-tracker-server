package zk

import (
	"expense-tracker-server/external/zk"
	"expense-tracker-server/internal/config"
	"fmt"
	"os"
)

// Connect ...
func Connect() {
	var (
		uri = os.Getenv("ZOOKEREPER_URI")
	)

	// Connect
	if err := zk.Connect(uri); err != nil {
		panic(err)
	}

	envVars := config.GetENV()
	server(envVars)
	commonValues(envVars)

}

// get value in zookeeper
func server(envVars *config.ENV) {
	// Admin
	adminPrefix := getAdminPrefix("")
	envVars.Admin.Server = zk.GetStringValue(fmt.Sprintf("%s/server", adminPrefix))
	envVars.Admin.Port = zk.GetStringValue(fmt.Sprintf("%s/port", adminPrefix))

	// App
	appPrefix := getAppPrefix("")
	envVars.App.Server = zk.GetStringValue(fmt.Sprintf("%s/server", appPrefix))
	envVars.App.Port = zk.GetStringValue(fmt.Sprintf("%s/port", appPrefix))

}

func getAppPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixApp, group)
}

func getAdminPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixAdmin, group)
}

func commonValues(envVars *config.ENV) {
}
