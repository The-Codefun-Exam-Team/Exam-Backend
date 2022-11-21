package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func DisplayLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func DisplayTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("1999-12-31 23:59:59"))
}

func DevelopmentLogging() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		DisableCaller: false,
		DisableStacktrace: false,
		Sampling: nil,
		Encoding: "console",
		OutputPaths: []string{"stdout"},
	}

	cfg.EncoderConfig.EncodeLevel = DisplayLevel
	cfg.EncoderConfig.EncodeTime = DisplayTime

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

func ProductionLogging() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		DisableCaller: false,
		DisableStacktrace: false,
		Sampling: nil,
		Encoding: "console",
		OutputPaths: []string{"stdout"},
	}

	cfg.EncoderConfig.EncodeLevel = DisplayLevel
	cfg.EncoderConfig.EncodeTime = DisplayTime

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

func InitializeLogging(mode string) (*zap.SugaredLogger, error) {
	if mode == "DEVELOPMENT" {
		return DevelopmentLogging()
	} else {
		return ProductionLogging()
	}
}