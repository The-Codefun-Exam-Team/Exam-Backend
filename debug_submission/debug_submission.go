package debugsubmission

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/labstack/echo/v4"
)

// Module contains an Env struct and a Group for routing
type Module struct {
	env   *envlib.Env
	group *echo.Group
}

// NewModule creates a new module with URL paths
func NewModule(gr *echo.Group, env *envlib.Env) *Module {
	module := &Module{
		env:   env,
		group: gr,
	}

	module.group.GET("/:id/", module.GetSingleSubmission)

	return module
}
