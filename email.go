package types

type EmailConfig struct {
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	From     string `json:"from" validate:"required,email"`
	To       string `json:"to" validate:"required,email"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}
