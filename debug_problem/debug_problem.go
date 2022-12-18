package debugproblem

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

	module.group.GET("/:code/", module.GetSingleProblem)
	module.group.GET("/", module.GetAllProblem)

	return module
}
