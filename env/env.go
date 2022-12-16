package env

import (
	"go.uber.org/zap"
)

type Env struct {
	config *Config
	log *zap.SugaredLogger
}