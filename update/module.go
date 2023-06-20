package update

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
)

type Module struct {
	env *envlib.Env
}

func NewModule(env *envlib.Env) *Module {
	m := &Module{
		env: env,
	}

	return m
}
