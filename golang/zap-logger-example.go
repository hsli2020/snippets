package main

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func main() {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339TimeEncoder,

			//CallerKey:    "caller",
			//EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "./log.txt"},
		ErrorOutputPaths: []string{"stderr"},
	}

	//config.EncoderConfig.EncodeLevel = CustomLevelEncoder
	//config.EncoderConfig.EncodeTime = CustomTimeEncoder

	zaplog, _ := config.Build()
	logger := zaplog.Sugar()
	logger.Info("this is a test config")

	url := "http://google.com"
	status := "200"
	err := errors.New("Just a test")

	logger.Debugf("Error fetching URL %s : Error = %s", url, err)
	logger.Warnf("Error fetching URL %s : Error = %s", url, err)
	logger.Errorf("Error fetching URL %s : Error = %s", url, err)
	logger.Infof("Success! statusCode = %s for URL %s", status, url)
}
