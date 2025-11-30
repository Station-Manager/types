package types

type AppConfig struct {
	DatastoreConfig      DatastoreConfig `json:"datastore_config"`
	LoggingConfig        LoggingConfig   `json:"logging_config"`
	RequiredConfigs      RequiredConfigs `json:"required_configs"`
	ServerConfig         *ServerConfig   `json:"server_config,omitempty"`
	RigConfigs           []RigConfig     `json:"rig_configs,omitempty"`
	LookupServiceConfigs []LookupConfig  `json:"lookup_service_configs,omitempty"`

	LoggingStation LoggingStation `json:"logging_station"`
}

type UiConfig struct {
	DefaultRigID int64   `json:"default_rig_id"`
	Logbook      Logbook `json:"logbook"`
	RigName      string  `json:"rig_name"`
	// DefaultFreq is in khz. This is used when the CAT system is unavailable.
	DefaultFreq string `json:"default_freq"`
	// DefaultMode is the CAT mode when the CAT system is unavailable.
	DefaultMode string `json:"default_mode"`
}
