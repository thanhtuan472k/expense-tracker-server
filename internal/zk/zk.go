package zk

import (
	"expense-tracker-server/external/util/pstring"
	"expense-tracker-server/external/zk"
	"expense-tracker-server/internal/config"
	"fmt"
	"os"
)

// Connect ...
func Connect() {
	var (
		uri = os.Getenv("ZOOKEEPER_URI")
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
	envVars.Admin.Auth.SecretKey = zk.GetStringValue(fmt.Sprintf("%s/authentication/secret_key", adminPrefix))
	envVars.Admin.Auth.TimeExpiredToken = pstring.StringToInt64(zk.GetStringValue(fmt.Sprintf("%s/authentication/time_expired_token", adminPrefix)))

	// App
	appPrefix := getAppPrefix("")
	envVars.App.Server = zk.GetStringValue(fmt.Sprintf("%s/server", appPrefix))
	envVars.App.Port = zk.GetStringValue(fmt.Sprintf("%s/port", appPrefix))
	envVars.App.Auth.SecretKey = zk.GetStringValue(fmt.Sprintf("%s/authentication/secret_key", appPrefix))
	envVars.App.Auth.TimeExpiredToken = pstring.StringToInt64(zk.GetStringValue(fmt.Sprintf("%s/authentication/time_expired_token", appPrefix)))

	// MongoDB
	mongodbPrefix := getExternalPrefix("/mongodb/expense")
	envVars.MongoDB.URI = zk.GetStringValue(fmt.Sprintf("%s/uri", mongodbPrefix))
	envVars.MongoDB.DBName = zk.GetStringValue(fmt.Sprintf("%s/db_name", mongodbPrefix))

	envVars.MongoDB.ReplicaSet = zk.GetStringValue(fmt.Sprintf("%s/replica_set", mongodbPrefix))
	envVars.MongoDB.CAPem = zk.GetStringValue(fmt.Sprintf("%s/ca_pem", mongodbPrefix))
	envVars.MongoDB.CertPem = zk.GetStringValue(fmt.Sprintf("%s/cert_pem", mongodbPrefix))
	envVars.MongoDB.CertKeyFilePassword = zk.GetStringValue(fmt.Sprintf("%s/cert_key_file_password", mongodbPrefix))
	envVars.MongoDB.ReadPrefMode = zk.GetStringValue(fmt.Sprintf("%s/read_pref_mode", mongodbPrefix))

	// Redis
	redisPrefix := getExternalPrefix("/redis/expense")
	envVars.Redis.URI = zk.GetStringValue(fmt.Sprintf("%s/uri", redisPrefix))
	envVars.Redis.Password = zk.GetStringValue(fmt.Sprintf("%s/auth/password", redisPrefix))

}

func getCommonPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixCommon, group)
}

func getExternalPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixExternal, group)
}

func commonValues(envVars *config.ENV) {
}

func getAppPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixApp, group)
}

func getAdminPrefix(group string) string {
	return fmt.Sprintf("%s%s", config.GetENV().ZookeeperPrefixAdmin, group)
}
