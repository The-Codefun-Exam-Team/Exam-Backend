package envlib

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/models"
)

// Env is a struct containing most objects that are designed to be pass around in the program.
type Env struct {
	Config *Config
	Log    *zap.SugaredLogger
	DB     *sqlx.DB
	Client models.Client
}
