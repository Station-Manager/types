package types

const LoggerServiceName = "logger"

type LoggingConfig struct {
	Level                  string `json:"level" validate:"oneof=trace debug info warn error fatal panic"`
	SkipFrameCount         int    `json:"skip_frame_count" validate:"min=0"`
	WithTimestamp          bool   `json:"with_timestamp"`
	ConsoleLogging         bool   `json:"console_logging"`
	FileLogging            bool   `json:"file_logging"`
	RelLogFileDir          string `json:"rel_log_file_dir" validate:"dir"`
	LogFileMaxBackups      int    `json:"log_file_max_backups" validate:"min=1"`
	LogFileMaxAgeDays      int    `json:"log_file_max_age_days" validate:"min=7"`
	LogFileMaxSizeMB       int    `json:"log_file_max_size_mb" validate:"min=10"`
	ShutdownTimeoutMS      int    `json:"shutdown_timeout_ms" validate:"omitempty,min=10,max=10000"` // Timeout for graceful shutdown (10ms-10s, 0=use default)
	ShutdownTimeoutWarning bool   `json:"shutdown_timeout_warning"`                                  // Log warning if shutdown timeout is exceeded
}
