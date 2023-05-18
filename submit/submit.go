package submit

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/labstack/echo/v4"
)

type Module struct {
	env *envlib.Env
}

func NewModule(gr *echo.Group, env *envlib.Env) *Module {
	module := &Module{
		env: env,
	}

	gr.POST("/", module.SubmitCode)

	return module
}
