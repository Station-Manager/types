package database

type Config struct {
	Driver          string `json:"driver" validate:"required"`
	Host            string `json:"host" validate:"required,hostname|ip"`
	Port            int    `json:"port" validate:"required,min=1,max=65535"`
	User            string `json:"user" validate:"required,min=1,max=63,alphanum|contains=_|contains=-"`
	Password        string `json:"pass" validate:"required,min=1"`
	Database        string `json:"database" validate:"required,min=1,max=63,alphanum|contains=_|contains=-"`
	SSLMode         string `json:"ssl_mode" validate:"required,oneof=disable require verify-ca verify-full"`
	MaxOpenConns    int    `json:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `json:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" validate:"required"` // Number of seconds
	Debug           bool   // Enable SQLBoiler query logging
}
