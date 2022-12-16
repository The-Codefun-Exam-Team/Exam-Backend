package debugproblem

import (
	"github.com/labstack/echo/v4"
	envlib "github.com/The-Codefun-Exam-Team/Exam-Backend/env"
)

type Module struct {
	env *envlib.Env
	group *echo.Group
}

func NewModule(gr *echo.Group, env *envlib.Env) (*Module) {
	module := &Module{
		env: env,
		group: gr,
	}

	return module
}

func (m *Module) GetSingleProblem(c echo.Context) error {
	// TODO: Verify user

	return nil
}