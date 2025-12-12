package types

type RequiredConfigs struct {
	DefaultRigID       int64  `json:"default_rig_id"`
	DefaultFreq        string `json:"default_freq"`
	DefaultMode        string `json:"default_mode"`
	OwnersCallsign     string `json:"owners_callsign"`
	DefaultIsRandomQso bool   `json:"default_is_random_qso"`
	PowerMultiplier    int    `json:"power_multiplier"`
	DefaultTxPower     int    `json:"default_tx_power"`
	UsePowerMultiplier bool   `json:"use_power_multiplier"`
	DefaultFwdEmail    string `json:"default_fwd_email"`
}
