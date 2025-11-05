package types

const AppConfigServiceName = "appconfig"

type AppConfig struct {
	DatastoreConfigs []DatastoreConfig
	LoggingConfig    LoggingConfig
}
