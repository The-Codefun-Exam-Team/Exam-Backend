package main

import (
	"fmt"
	envlib "github.com/The-Codefun-Exam-Team/Exam-Backend/env"
)

func main() {
	var err error

	// Creating the env
	env := envlib.Env{}

	// Load config
	env.Config, err = envlib.LoadConfig()

	if err != nil {
		panic(fmt.Sprintf("[cannot load config] %v", err))
	}

	// Initialize logger
	env.Log, err = envlib.InitializeLogger(env.Config.LoggingMode)

	if err != nil {
		panic(fmt.Sprintf("[cannot initialize logger] %v", err))
	}

	// Connect to database
	db_dsn := envlib.GetDSN(env.Config)
	env.DB, err = envlib.NewDB(db_dsn)

	if err != nil {
		panic(fmt.Sprintf("[cannot connect to database] %v", err))
	}
}
