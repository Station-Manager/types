package types

type Logbook struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Callsign    string `json:"callsign"`
	Description string `json:"description"`
}

type LogbookList []Logbook
