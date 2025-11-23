package types

import (
	"go.bug.st/serial"
	"time"
)

type SerialConfig struct {
	PortName      string
	BaudRate      int
	DataBits      int
	Parity        serial.Parity
	StopBits      serial.StopBits
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	RTS           bool
	DTR           bool
	LineDelimiter byte
}
