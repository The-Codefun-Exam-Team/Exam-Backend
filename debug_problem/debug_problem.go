package debugproblem

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/labstack/echo/v4"
)

type Module struct {
	env   *envlib.Env
	group *echo.Group
}

func NewModule(gr *echo.Group, env *envlib.Env) *Module {
	module := &Module{
		env:   env,
		group: gr,
	}

	module.group.GET("/:code/", module.GetSingleProblem)

	return module
}
