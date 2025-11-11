package types

type Logbook struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Callsign    string `json:"callsign"`
	Description string `json:"description"`
	ApiKey      string `json:"api_key"`
}

type LogbookList []Logbook
