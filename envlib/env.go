package envlib

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Env is a struct containing most objects that are designed to be pass around in the program.
type Env struct {
	Config *Config
	Log    *zap.SugaredLogger
	DB     *sqlx.DB
}
