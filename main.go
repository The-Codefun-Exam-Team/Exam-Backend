package main

import (
	"fmt"
	envlib "github.com/The-Codefun-Exam-Team/Exam-Backend/env"
)

func main() {
	var err error
	env := envlib.Env{}

	env.Config, err = envlib.LoadConfig()

	if err != nil {
		panic(fmt.Sprintf("[cannot load config] %v", err))
	}

	env.Log, err = envlib.InitializeLogging(env.Config.LoggingMode)

	if err != nil {
		panic(fmt.Sprintf("[cannot initialize logger] %v", err))
	}

	db_dsn := envlib.GetDSN(env.Config)
	env.DB, err = envlib.New(db_dsn)

	if err != nil {
		panic(fmt.Sprintf("[cannot connect to database] %v", err))
	}
}
