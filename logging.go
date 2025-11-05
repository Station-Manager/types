package types

const LoggerServiceName = "logger"

type LoggingConfig struct {
	Level             string `json:"level" validate:"oneof=debug info warn error fatal panic"`
	SkipFrameCount    int    `json:"skip_frame_count" validate:"min=0"`
	WithTimestamp     bool   `json:"with_timestamp"`
	ConsoleLogging    bool   `json:"console_logging"`
	FileLogging       bool   `json:"file_logging"`
	RelLogFileDir     string `json:"rel_log_file_dir" validate:"dir"`
	LogFileMaxBackups int    `json:"log_file_max_backups" validate:"min=1"`
	LogFileMaxAgeDays int    `json:"log_file_max_age_days" validate:"min=7"`
	LogFileMaxSizeMB  int    `json:"log_file_max_size_mb" validate:"min=10"`
}
