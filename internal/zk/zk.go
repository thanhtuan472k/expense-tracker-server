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
		uri = os.Getenv("ZOOKEEPER_URI")
	)

	// Connect
	if err := zk.Connect(uri); err != nil {
		fmt.Println("errrrr", err)
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

	// MongoDB
	mongodbPrefix := getExternalPrefix("/mongodb/expense")
	envVars.MongoDB.URI = zk.GetStringValue(fmt.Sprintf("%s/uri", mongodbPrefix))
	envVars.MongoDB.DBName = zk.GetStringValue(fmt.Sprintf("%s/db_name", mongodbPrefix))

	envVars.MongoDB.ReplicaSet = zk.GetStringValue(fmt.Sprintf("%s/replica_set", mongodbPrefix))
	envVars.MongoDB.CAPem = zk.GetStringValue(fmt.Sprintf("%s/ca_pem", mongodbPrefix))
	envVars.MongoDB.CertPem = zk.GetStringValue(fmt.Sprintf("%s/cert_pem", mongodbPrefix))
	envVars.MongoDB.CertKeyFilePassword = zk.GetStringValue(fmt.Sprintf("%s/cert_key_file_password", mongodbPrefix))
	envVars.MongoDB.ReadPrefMode = zk.GetStringValue(fmt.Sprintf("%s/read_pref_mode", mongodbPrefix))

	// NATS
	//natsPrefix := getExternalPrefix("/nats/expense")
	//envVars.Nats.URL = zk.GetStringValue(fmt.Sprintf("%s/uri", natsPrefix))
	//envVars.Nats.Username = zk.GetStringValue(fmt.Sprintf("%s/user", natsPrefix))
	//envVars.Nats.Password = zk.GetStringValue(fmt.Sprintf("%s/password", natsPrefix))
	//envVars.Nats.APIKey = zk.GetStringValue(fmt.Sprintf("%s/api_key", natsPrefix))

	// MongoDB_AUDIT
	//mongoAuditPrefix := getExternalPrefix("/mongodb/campaign_audit")
	//envVars.MongoAudit.Host = zk.GetStringValue(fmt.Sprintf("%s/host", mongoAuditPrefix))
	//envVars.MongoAudit.DBName = zk.GetStringValue(fmt.Sprintf("%s/db_name", mongoAuditPrefix))

	// FileHost
	//commonPrefix := getCommonPrefix("")
	//envVars.FileHost = zk.GetStringValue(fmt.Sprintf("%s/file_host", commonPrefix))

	// Authentication
	//authPrefix := getAppPrefix("/authentication")
	//envVars.SecretKey = zk.GetStringValue(fmt.Sprintf("%s/auth_secretkey", authPrefix))

	// AUTH GG
	//envVars.EnableAuthenticationService = zk.GetStringValue(fmt.Sprintf("%s/authentication_google/enable", adminPrefix))

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
