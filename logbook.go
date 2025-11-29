package types

type Logbook struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`  // FK to users.id
	Name        string `json:"name"`     // Unique name (to the user) for the logbook
	Callsign    string `json:"callsign"` // The callsign associated with the logbook
	Description string `json:"description"`
}

type LogbookList []Logbook
