package database

type Config struct {
	Driver          string `json:"driver" validate:"required"`
	DSN             string `json:"dsn" validate:"required"`
	MaxOpenConns    int    `json:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `json:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" validate:"required"` // Number of seconds
	Debug           bool   // Enable SQLBoiler query logging
}
