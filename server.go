package types

type RequestAction string

func (s RequestAction) String() string {
	return string(s)
}

const (
	// RegisterLogbookAction registers a new logbook with the server.
	RegisterLogbookAction RequestAction = "register_logbook"
	InsertQsoAction                     = "insert_qso"
)

type PostRequest struct {
	Callsign string        `json:"callsign"` // The callsign associated with the user account *NOT THE LOGBOOK CALLSIGN*. The logbook callsign is associated with Key.
	Key      string        `json:"key"`      // Logbook's API Key, or if registering, the Bootstrap Key
	Action   RequestAction `json:"action"`   // The action to perform
	Data     string        `json:"data"`     // The data to send
}
