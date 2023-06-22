package main

import (
	"fmt"
	"net/http"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/create_problem"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/debug_problem"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/debug_submission"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/submit"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/envlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Create HTTP client
	env.Client = http.DefaultClient

	env.Log.Info("Environment created")

	// Create the echo.Echo object
	e := echo.New()

	// Use middleware
	e.Pre(middleware.AddTrailingSlash())

	// Attach the route to /api/problems
	_ = debugproblem.NewModule(e.Group("/api/problems"), &env)

	// Attach the route to /api/submissions
	_ = debugsubmission.NewModule(e.Group("/api/submissions"), &env)

	// Attach the route to /api/submit
	_ = submit.NewModule(e.Group("/api/submit"), &env)

	// Attach the route to /api/new_problem
	_ = create.NewModule(e.Group("/api/new_problem"), &env)

	if err = e.Start(fmt.Sprintf(":%v", env.Config.ServerPort)); err != nil {
		env.Log.Fatalf("Cannot start server: %v", err)
	}
}
