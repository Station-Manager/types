package types

import "time"

type LookupConfig struct {
	Name        string        `json:"name"`
	Enabled     bool          `json:"enabled"`
	URL         string        `json:"url"`
	Username    string        `json:"username,omitempty"`
	Password    string        `json:"password,omitempty"`
	UserAgent   string        `json:"useragent"`
	HttpTimeout time.Duration `json:"timeout"`
	ViewUrl     string        `json:"view_url,omitempty"`
}
