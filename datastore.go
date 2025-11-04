package types

type DatastoreConfig struct {
	Driver          string `json:"driver" validate:"required,oneof=postgres sqlite3"`
	Path            string `json:"path" validate:"required_if=Driver sqlite3,min=1"`    // Used for sqlite3 only
	Options         string `json:"options" validate:"required_if=Driver sqlite3,min=1"` // Used for sqlite3 only
	Host            string `json:"host" validate:"required,hostname|ip"`
	Port            int    `json:"port" validate:"required,min=1,max=65535"`
	User            string `json:"user" validate:"required,min=1,max=63,alphanum|contains=_|contains=-"`
	Password        string `json:"pass" validate:"required,min=1"`
	Database        string `json:"database" validate:"required,min=1,max=63,alphanum|contains=_|contains=-"`
	SSLMode         string `json:"ssl_mode" validate:"required,oneof=disable require verify-ca verify-full"`
	MaxOpenConns    int    `json:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns    int    `json:"max_idle_conns" validate:"required,min=1"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" validate:"required,min=0"`  // Number of minutes: 1-15 minutes
	ConnMaxIdleTime int    `json:"conn_max_idle_time" validate:"required,min=0"` // Number of minutes: 1-5 minutes

	ContextTimeout            int `json:"context_timeout" validate:"required,min=5"`             // Number of seconds, with a minimum of 5 seconds
	TransactionContextTimeout int `json:"transaction_context_timeout" validate:"required,min=5"` // Number of seconds, with a minimum of 5 seconds

	Debug bool // Enable SQLBoiler query logging

	Params map[string]string `json:"params"`
}
