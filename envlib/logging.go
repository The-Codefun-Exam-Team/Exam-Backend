package envlib

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Display the level as [INFO] or [WARN]
func displayLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// Display the time according to ISO 8601
func displayTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// developmentLogging creates a logger with options for development purposes.
func developmentLogging() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel), // Starts logging at debug level
		Development:       true,                                 // Turns DPanic into Panic
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
	}

	// Set the EncoderConfig
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		LineEnding:  "\n",
		EncodeLevel: displayLevel,
		EncodeTime:  displayTime,
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

// productionLogging creates a logger with options for production.
func productionLogging() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel), // Starts logging at info level
		Development:       false,                               // Turns DPanic into Error
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
	}

	// Set the EncoderConfig
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		LineEnding:  "\n",
		EncodeLevel: displayLevel,
		EncodeTime:  displayTime,
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

// testingLogging creates a logger with options for testing.
func testingLogging() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.FatalLevel), // Starts logging at info level
		Development:       false,                                // Turns DPanic into Error
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		OutputPaths:       []string{},
	}

	// Set the EncoderConfig
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		LineEnding:  "\n",
		EncodeLevel: displayLevel,
		EncodeTime:  displayTime,
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

// InitializeLogger takes in a mode and return an appropriate logger
// The modes allowed are "DEVELOPMENT" and "PRODUCTION"
func InitializeLogger(mode string) (*zap.SugaredLogger, error) {
	if mode == "DEVELOPMENT" {
		return developmentLogging()
	} else if mode == "PRODUCTION" {
		return productionLogging()
	} else if mode == "TESTING" {
		return testingLogging()
	} else {
		return nil, errors.New("mode not found")
	}
}
