package types

import "time"

type RegisterLogbookRequest struct {
	Logbook
}

type RegisterLogbookResponse struct {
	RemoteLogbookID int64      `json:"remote_logbook_id"`
	APIKey          string     `json:"api_key"`
	CreatedAt       time.Time  `json:"created_at"`
	ExpiresAt       *time.Time `json:"expires_at"`
}
