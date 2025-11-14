package types

type Logbook struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Callsign    string `json:"callsign"`
	Description string `json:"description"`
	ApiKey      string `json:"api_key,omitempty"` // Not part of the server models
	UserID      int64  `json:"user_id"`           // FK
}

type LogbookList []Logbook
