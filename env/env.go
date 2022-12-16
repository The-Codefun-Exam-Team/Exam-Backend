package env

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Env struct {
	Config *Config
	Log *zap.SugaredLogger
	DB *sqlx.DB
}