package types

import (
	"go.bug.st/serial"
	"time"
)

type SerialConfig struct {
	PortName       string
	BaudRate       int
	DataBits       int
	Parity         serial.Parity
	StopBits       serial.StopBits
	ReadTimeoutms  time.Duration // Milliseconds
	WriteTimeoutms time.Duration // Milliseconds
	RTS            bool
	DTR            bool
	LineDelimiter  byte
}
