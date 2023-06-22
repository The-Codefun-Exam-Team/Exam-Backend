package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// Use middleware and change settings
	e.Pre(middleware.AddTrailingSlash())
	e.HideBanner = true

	// Attach the route to /api/problems
	_ = debugproblem.NewModule(e.Group("/api/problems"), &env)

	// Attach the route to /api/submissions
	_ = debugsubmission.NewModule(e.Group("/api/submissions"), &env)

	// Attach the route to /api/submit
	_ = submit.NewModule(e.Group("/api/submit"), &env)

	// Attach the route to /api/new_problem
	_ = create.NewModule(e.Group("/api/new_problem"), &env)

	// Start the server in a goroutine
	go func() {
		if err = e.Start(fmt.Sprintf(":%v", env.Config.ServerPort)); err != nil && err != http.ErrServerClosed {
			env.Log.Fatalf("Cannot start server: %v", err)
		}
	}()

	// Receive os.Interrupt signal (Ctrl+C)
	interrupt_channel := make(chan os.Signal, 1)
	signal.Notify(interrupt_channel, os.Interrupt)
	<- interrupt_channel

	// Graceful shutdown
	// Timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		env.Log.Fatalf("Error shutting down server: %v", err)
	}
}
