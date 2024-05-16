package ilogger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
	err       error
)

// LogData for tracking
type LogData map[string]interface{}

func Init() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if zapLogger, err = cfg.Build(); err != nil {
		panic(err)
	}
}

// Error ...
func Error(content string, data LogData) {
	jsonData, _ := json.Marshal(data)
	zapLogger.Error(content, zap.String("data", string(jsonData)))
}
