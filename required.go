package types

import "time"

type RequiredConfigs struct {
	DefaultRigID       int64  `json:"default_rig_id"`
	DefaultFreq        string `json:"default_freq"`
	DefaultMode        string `json:"default_mode"`
	DefaultIsRandomQso bool   `json:"default_is_random_qso"`
	PowerMultiplier    int    `json:"power_multiplier"`
	DefaultTxPower     int    `json:"default_tx_power"`
	UsePowerMultiplier bool   `json:"use_power_multiplier"`
	// The default TO email address.
	DefaultFwdEmail              string        `json:"default_fwd_email"`
	QsoForwardingIntervalSeconds time.Duration `json:"qso_forwarding_interval_seconds"`
}
