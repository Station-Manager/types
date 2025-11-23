package types

type AppConfig struct {
	DatastoreConfig DatastoreConfig `json:"datastore_config"`
	LoggingConfig   LoggingConfig   `json:"logging_config"`
	RequiredConfigs RequiredConfigs `json:"required_configs"`
	ServerConfig    ServerConfig    `json:"server_config"`
	RigConfigs      []RigConfig     `json:"rig_configs"`
}
