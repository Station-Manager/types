package types

type RequiredConfigs struct {
	DefaultRigID int64  `json:"default_rig_id"`
	DefaultFreq  string `json:"default_freq"`
	DefaultMode  string `json:"default_mode"`
}
