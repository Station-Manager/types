package types

type RequiredConfigs struct {
	DefaultRigID       int64  `json:"default_rig_id"`
	DefaultFreq        string `json:"default_freq"`
	DefaultMode        string `json:"default_mode"`
	PowerMultiplier    string `json:"power_multiplier"`
	OwnersCallsign     string `json:"owners_callsign"`
	DefaultIsRandomQso bool   `json:"default_is_random_qso"`
	DefaultTxPower     int    `json:"default_tx_power"`
}
