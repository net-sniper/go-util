package log

import (
	"context"
)

// LevelLogger represents the ability to log non-error messages, at a particular verbosity.
type LevelLogger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

type InfoLogger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	InfoDepth(depth int, args ...interface{})
}

// WarningLogger represents the ability to log warning messages.
type WarningLogger interface {
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	WarningDepth(depth int, args ...interface{})
}

// ErrorLogger represents the ability to log error messages.
type ErrorLogger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	ErrorDepth(depth int, args ...interface{})
}

// FatalLogger represents the ability to log fatal messages.
type FatalLogger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	FatalDepth(depth int, args ...interface{})
}

// ExitLogger Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
type ExitLogger interface {
	Exit(args ...interface{})
	Exitf(format string, args ...interface{})
	ExitDepth(depth int, args ...interface{})
}

// Logger represents the ability to log messages, both errors and not.
type Logger interface {
	// All Loggers implement InfoLogger.  Calling InfoLogger methods directly on
	// a Logger value is equivalent to calling them on a V(0) InfoLogger.  For
	// example, logger.Info() produces the same result as logger.V(0).Info.
	InfoLogger
	WarningLogger
	ErrorLogger
	FatalLogger
	ExitLogger

	V(level int) LevelLogger
	SetLevel(level int) error

	// WithContext returns a copy of context in which the log value is set.
	WithContext(ctx context.Context) Logger

	// Flush calls the underlying Core's Sync method, flushing any buffered
	// log entries. Applications should take care to call Sync before exiting.
	Flush()
}
