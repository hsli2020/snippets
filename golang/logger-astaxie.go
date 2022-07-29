// https://gist.github.com/finallyayo/7cd6460483c2e88708fa140fe378c0e4
// Adapted from https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/13.4.html
package logger

import (
	"log"
	"os"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

// logLevel controls the global log level used by the logger.
var level = LevelTrace

// LogLevel returns the global log level and can be used in
// a custom implementations of the logger interface.
func Level() int {
	return level
}

// SetLogLevel sets the global log level used by the simple
// logger.
func SetLevel(l int) {
	level = l
}

// logger references the used application logger.
var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

// SetLogger sets a new logger.
func SetLogger(l *log.Logger) {
	Logger = l
}

// Trace logs a message at trace level.
func Trace(v ...interface{}) {
	if level <= LevelTrace {
		Logger.Printf("[TRACE] %v\n", v...)
	}
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	if level <= LevelDebug {
		Logger.Printf("[DEBUG] %v\n", v...)
	}
}

// Info logs a message at info level.
func Info(v ...interface{}) {
	if level <= LevelInfo {
		Logger.Printf("[INFO] %v\n", v...)
	}
}

// Warning logs a message at warning level.
func Warn(v ...interface{}) {
	if level <= LevelWarning {
		Logger.Printf("[WARN] %v\n", v...)
	}
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	if level <= LevelError {
		Logger.Printf("[ERR] %v\n", v...)
	}
}

// Fatak logs a message at fatal level.
func Fatal(v ...interface{}) {
	if level <= LevelFatal {
		Logger.Printf("[FATAL] %v\n", v...)
		os.Exit(1)
	}
}