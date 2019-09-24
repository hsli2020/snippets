package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	LogFile         = "logs"
	DebugFlag       = "Debug      "
	InformationFlag = "Information"
	WarningFlag     = "Warning    "
	ErrorFlag       = "Error      "
	FatalFlag       = "Fatal      "
	TimeFormat      = "2006-01-02 15:04:05"
	LineBreak       = "\n"
)

type logger struct {
	logFile      string
	mutex        *sync.Mutex
	exitFunction func()
}

func NewInstance() *logger {
	logger := &logger{
		logFile:      LogFile,
		mutex:        &sync.Mutex{},
		exitFunction: exit,
	}
	return logger
}

func (logger *logger) Debug(message string) {
	logger.log(DebugFlag, message)
}

func (logger *logger) Debugf(messageFormat string, values ...interface{}) {
	message := fmt.Sprintf(messageFormat, values...)
	logger.log(DebugFlag, message)
}

func (logger *logger) Information(message string) {
	logger.log(InformationFlag, message)
}

func (logger *logger) Informationf(messageFormat string, values ...interface{}) {
	message := fmt.Sprintf(messageFormat, values...)
	logger.log(InformationFlag, message)
}

func (logger *logger) Warning(message string) {
	logger.log(WarningFlag, message)
}

func (logger *logger) Warningf(messageFormat string, values ...interface{}) {
	message := fmt.Sprintf(messageFormat, values...)
	logger.log(WarningFlag, message)
}

func (logger *logger) Error(message string) {
	logger.log(ErrorFlag, message)
}

func (logger *logger) Errorf(messageFormat string, values ...interface{}) {
	message := fmt.Sprintf(messageFormat, values...)
	logger.log(ErrorFlag, message)
}

func (logger *logger) Fatal(message string) {
	logger.log(FatalFlag, message)
	if logger.exitFunction != nil {
		logger.exitFunction()
	}
}

func (logger *logger) Fatalf(messageFormat string, values ...interface{}) {
	message := fmt.Sprintf(messageFormat, values...)
	logger.log(FatalFlag, message)
	if logger.exitFunction != nil {
		logger.exitFunction()
	}
}

func (logger *logger) log(flag string, message string) {
	//synchronization
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	//create formatted message
	formattedMessage := fmt.Sprintf("[%s][%s]: %s%s", flag, time.Now().Format(TimeFormat), message, LineBreak)

	//open log file
	logFile, err := os.OpenFile(logger.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		printError("Opening log file failed!")
		return
	}
	defer func() {
		err := logFile.Close()
		if err != nil {
			printError("Closing log file failed!")
		}
	}()

	//write to log file
	_, err = logFile.WriteString(formattedMessage)
	if err != nil {
		printError("Writing to log file failed!")
	}
}

func exit() {
	os.Exit(1)
}

func printError(message string) {
	os.Stderr.WriteString(fmt.Sprintf("go_logger: %s%s", message, LineBreak))
}
