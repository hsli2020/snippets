package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logger := SetupLogger("error.log")

	logger.Info("info message")
	logger.Infof("%s", "info message")

	logger.Debug("debug message")
	logger.Debugf("%s", "debug message")

	logger.Error("error message")
	logger.Errorf("%s", "error message")
}

func SetupLogger(filename string) *Logger {
	logFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//defer logFile.Close()

	out := io.MultiWriter(os.Stdout, logFile)
	logger := NewLogger(out, "", log.LstdFlags) // |log.Lshortfile)

	return logger
}

// 自定义 logger
type Logger struct {
	stdlog *log.Logger
}

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "[DEBUG] "
	case LevelInfo:
		return "[INFO ] "
	case LevelError:
		return "[ERROR] "
	}
	return ""
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{stdlog: l}
}

func (l *Logger) Debug(v ...interface{}) {
	l.stdlog.Print(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.stdlog.Print(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.stdlog.Print(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.stdlog.Print(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.stdlog.Print(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.stdlog.Print(LevelError, fmt.Sprintf(format, v...))
}
