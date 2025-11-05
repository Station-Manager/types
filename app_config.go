package types

const AppConfigServiceName = "appconfig"

type AppConfig struct {
	DatastoreConfig DatastoreConfig `json:"datastore_config"`
	LoggingConfig   LoggingConfig   `json:"logging_config"`
	RequiredConfigs RequiredConfigs `json:"required_configs"`
}
