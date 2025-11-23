package types

import "time"

type CatCommand struct {
	Name string
	Cmd  string
}

type CatState struct {
	Prefix  string // The CAT command prefix.
	Markers []Marker
}

type Marker struct {
	Tag           string
	Index         int
	Length        int
	ValueMappings []ValueMapping
}

func (c *CatState) New(prefix string) CatState {
	return CatState{
		Prefix:  prefix,
		Markers: make([]Marker, 0),
	}
}

type ValueMapping struct {
	Key   string
	Value string
}

// CatStatus is a map of CAT status values keyed by state tag. These are related to CAT commands; a particular
// command when issued will return a status value linked to that command. The status value is a string that
// represents the status of the CAT command in a format usable by the frontend.
//
// See the cat.StatusChannel in the cat/service.go file for more details.
type CatStatus map[string]string

// StateValues is a map of state values keyed by state tag. These are related to CAT commands; a particular command
// when issued will return a state value linked to that command. The state value is a string that represents the
// state of the CAT command in user-readable format.
//
// See the cat.StateValues() method in the cat/service.go file for more details.
type StateValues map[string]map[string]string

type CatConfig struct {
	RateLimiterInterval      time.Duration   `koanf:"rate_limiter_interval" validate:"min=1000000,max=10000000000"` // 1ms to 10s in nanoseconds
	ReadBufferSize           int64           `koanf:"read_buffer_size" validate:"min=64,max=65536"`
	CmdChannelSize           int64           `koanf:"cmd_channel_size" validate:"min=1,max=1000"`
	ReplyChannelSize         int64           `koanf:"reply_channel_size" validate:"min=1,max=1000"`
	StatusChannelSize        int64           `koanf:"status_channel_size" validate:"min=1,max=100"`
	CommandTimeout           time.Duration   `koanf:"command_timeout" validate:"min=0,max=30000000000"` // 0 to 30s in nanoseconds (0 = no timeout)
	RateLimiterCmdsPerSecond int             `koanf:"rate_limiter_cmds_per_second" validate:"min=1,max=20"`
	Reconnection             ReconnectConfig `koanf:"reconnection"`
}

type ReconnectConfig struct {
	Enabled             bool          `koanf:"enabled"`                                                         // Enable automatic reconnection
	MaxRetries          int           `koanf:"max_retries" validate:"min=0,max=100"`                            // 0 = infinite retries
	InitialBackoff      time.Duration `koanf:"initial_backoff" validate:"min=100000000,max=30000000000"`        // 100ms to 30s
	MaxBackoff          time.Duration `koanf:"max_backoff" validate:"min=1000000000,max=300000000000"`          // 1s to 5min
	BackoffMultiplier   float64       `koanf:"backoff_multiplier" validate:"min=1.0,max=10.0"`                  // Exponential backoff factor
	HealthCheckInterval time.Duration `koanf:"health_check_interval" validate:"min=1000000000,max=60000000000"` // 1s to 60s
	FailureThreshold    int           `koanf:"failure_threshold" validate:"min=1,max=20"`                       // Consecutive failures before disconnect
	RecoveryThreshold   int           `koanf:"recovery_threshold" validate:"min=1,max=10"`                      // Consecutive successes before considering healthy
}
