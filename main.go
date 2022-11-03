package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	debugproblem "github.com/The-Codefun-Exam-Team/Exam-Backend/debug_problem"
	debugsubmission "github.com/The-Codefun-Exam-Team/Exam-Backend/debug_submission"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/rankings"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/submit"
)

func main() {
	e := echo.New()

	// e.Pre(middleware.HTTPSRedirect())
	e.Pre(middleware.AddTrailingSlash())

	db, err := db.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	e.GET("/api/ping/", general.Ping)
	e.GET("/api/debug_submission/:id/", general.TempSubmission)

	if _, err := debugproblem.New(db, e.Group("/api/problems")); err != nil {
		log.Fatal(err)
	}

	if _, err := submit.New(db, e.Group("/api/submit")); err != nil {
		log.Fatal(err)
	}

	if _, err := debugsubmission.New(db, e.Group("/api/submission")); err != nil {
		log.Fatal(err)
	}

	if _, err := rankings.New(db, e.Group("/api/rankings")); err != nil {
		log.Fatal(err)
	}


	if err := e.Start(":80"); err != nil {
		log.Fatal(err)
	}
}
