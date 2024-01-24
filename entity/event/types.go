package event

// LogLevel correlates to a log level for MAAS events.
type LogLevel string

// Log levels referring from MAAS server https://github.com/maas/maas/blob/master/src/maasserver/models/eventtype.py#L25
const (
	AUDIT    LogLevel = "AUDIT"
	CRITICAL LogLevel = "CRITICAL"
	DEBUG    LogLevel = "DEBUG"
	ERROR    LogLevel = "ERROR"
	WARNING  LogLevel = "WARNING"
	INFO     LogLevel = "INFO"
)
