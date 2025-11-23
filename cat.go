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
	RateLimiterInterval      time.Duration // Milliseconds
	ReadBufferSize           int64
	CmdChannelSize           int64
	ReplyChannelSize         int64
	StatusChannelSize        int64
	CommandTimeout           time.Duration // Millseconds
	RateLimiterCmdsPerSecond int
}
