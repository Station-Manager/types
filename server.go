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

type ServerConfig struct {
	Name         string `json:"name"` // AppName for goFiber
	Host         string `json:"host"`
	Port         int    `json:"port"`
	TLSEnabled   bool   `json:"tls_enabled"`
	TLSCertFile  string `json:"tls_cert_file"`
	TLSKeyFile   string `json:"tls_key_file"`
	ReadTimeout  int    `json:"read_timeout"`  // Seconds
	WriteTimeout int    `json:"write_timeout"` // Seconds
	IdleTimeout  int    `json:"idle_timeout"`  // Seconds
}
