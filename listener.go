package types

type ListenerConfig struct {
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Protocol   string `json:"protocol"` // UDP or TCP
	BufferSize int    `json:"buffer_size"`
}
