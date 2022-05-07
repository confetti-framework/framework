package inter

import (
	"github.com/confetti-framework/syslog/log_level"
)

// Logger This interface is the interface you should use to add a logger yourself.
type Logger interface {
	SetApp(app AppReader) Logger
	Clear() bool
	Group(group string) Logger
	Log(severity log_level.Level, message string, arguments ...interface{})
	LogWith(severity log_level.Level, message string, data interface{})
}

// LoggerFacade This interface ensures that the loggers can be used with
// the convenient leveling methods. That while the concrete
// logger only needs to implement inter.Logger.
type LoggerFacade interface {
	SetApp(app AppReader) LoggerFacade
	Group(group string) LoggerFacade
	Log(severity log_level.Level, message string, arguments ...interface{})
	LogWith(severity log_level.Level, message string, data interface{})

	// Emergency Log that the system is unusable
	Emergency(message string, arguments ...interface{})

	// EmergencyWith Log that the system is unusable
	EmergencyWith(message string, data interface{})

	// Alert A condition that should be corrected immediately.
	Alert(message string, arguments ...interface{})

	// AlertWith A condition that should be corrected immediately.
	AlertWith(message string, data interface{})

	// Critical conditions
	Critical(message string, arguments ...interface{})

	// CriticalWith Critical conditions
	CriticalWith(message string, data interface{})

	// Error conditions
	Error(message string, arguments ...interface{})

	// ErrorWith Error conditions
	ErrorWith(message string, data interface{})

	// Warning conditions
	Warning(message string, arguments ...interface{})

	// WarningWith Warning conditions
	WarningWith(message string, data interface{})

	// Notice Normal but significant conditions
	// Conditions that are not error conditions, but that may require special handling.
	Notice(message string, arguments ...interface{})

	// NoticeWith Normal but significant conditions
	// Conditions that are not error conditions, but that may require special handling.
	NoticeWith(message string, data interface{})

	// Info Informational messages
	Info(message string, arguments ...interface{})

	// InfoWith Informational messages
	InfoWith(message string, data interface{})

	// Debug Debug-level messages
	// Messages containing information that is normally only useful when debugging a program.
	Debug(message string, arguments ...interface{})

	// DebugWith Debug-level messages
	// Messages containing information that is normally only useful when debugging a program.
	DebugWith(message string, data interface{})
}
