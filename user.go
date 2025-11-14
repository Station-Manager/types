package types

import "time"

type User struct {
	ID                 int64     `json:"id"`
	Callsign           string    `json:"callsign" validate:"min=3,max=30,alphanum"`
	PassHash           string    `json:"pass_hash" validate:"required"`
	Issuer             string    `json:"issuer,omitempty"`
	Subject            string    `json:"subject,omitempty"`
	Email              string    `json:"email,omitempty"`
	EmailConfirmed     bool      `json:"email_confirmed"`
	BootstrapHash      string    `json:"bootstrap_hash,omitempty"`
	BootstrapExpiresAt time.Time `json:"bootstrap_expires_at,omitempty"`
	BootstrapUsedAt    time.Time `json:"bootstrap_used_at,omitempty"`
}
