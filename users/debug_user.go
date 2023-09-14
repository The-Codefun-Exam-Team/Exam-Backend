package debuguser

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/labstack/echo/v4"
)

// Module contains an Env struct and a Group for routing
type Module struct {
	env *envlib.Env
}

// NewModule creates a new module with URL paths
func NewModule(gr *echo.Group, env *envlib.Env) *Module {
	module := &Module{env: env}

	gr.GET("/:username/stats/", module.GetUserStats)

	return module
}
