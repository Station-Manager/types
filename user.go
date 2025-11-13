package types

type User struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	PassHash           string `json:"pass_hash"`
	Issuer             string `json:"issuer,omitempty"`
	Subject            string `json:"subject,omitempty"`
	Email              string `json:"email,omitempty"`
	BootstrapHash      string `json:"bootstrap_hash,omitempty"`
	BootstrapExpiresAt string `json:"bootstrap_expires_at,omitempty"`
	BootstrapUsedAt    string `json:"bootstrap_used_at,omitempty"`
}
