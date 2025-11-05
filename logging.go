package types

const LoggerServiceName = "logger"

type LoggingConfig struct {
	Level             string
	SkipFrameCount    int
	WithTimestamp     bool
	ConsoleLogging    bool
	FileLogging       bool
	RelLogFileDir     string
	LogFileMaxBackups int
	LogFileMaxAgeDays int
	LogFileMaxSizeMB  int
}
