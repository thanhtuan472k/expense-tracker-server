package config

import "os"

// ENV ...
type ENV struct {
	Env                     string
	ZookeeperPrefixCommon   string
	ZookeeperPrefixExternal string
	ZookeeperPrefixApp      string
	ZookeeperPrefixAdmin    string
	MongoDB                 struct {
		URI                 string
		DBName              string
		ReplicaSet          string
		CAPem               string
		CertPem             string
		CertKeyFilePassword string
		ReadPrefMode        string
	}
	Admin struct {
		Server string
		Port   string
	}
	App struct {
		Server string
		Port   string
	}
	Nats       NatsConfig
	SecretKey  string
	MongoAudit MongoConfig `env:",prefix=MONGO_AUDIT_"`

	FileHost string

	EnableAuthenticationService string

	// Redis
	Redis RedisCfg
}

// RedisCfg ...
type RedisCfg struct {
	URI      string
	Password string
}

// NatsConfig ...
type NatsConfig struct {
	URL      string
	Username string
	Password string
	APIKey   string
}

// MongoConfig ...
type MongoConfig struct {
	Host   string
	DBName string
}

var env ENV

// GetENV ...
func GetENV() *ENV {
	return &env
}

// Init ...
func Init() {
	env = ENV{
		Env:                     os.Getenv("ENV"),
		ZookeeperPrefixExternal: os.Getenv("ZOOKEEPER_PREFIX_EXTERNAL"),
		ZookeeperPrefixCommon:   os.Getenv("ZOOKEEPER_PREFIX_CAMPAIGN_COMMON"),
		ZookeeperPrefixApp:      os.Getenv("ZOOKEEPER_PREFIX_CAMPAIGN_APP"),
		ZookeeperPrefixAdmin:    os.Getenv("ZOOKEEPER_PREFIX_CAMPAIGN_ADMIN"),
	}
}
