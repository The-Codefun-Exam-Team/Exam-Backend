package main

import (
	"log"
	
	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/debug_problem"
)

func main(){
	e := echo.New()

	db, err := db.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	e.GET("/ping", general.Ping)
	e.GET("/debug_problem", general.TempDebug)
	e.GET("/debug_submission/:id", general.TempSubmission)

	if _, err := debugproblem.New(db, e.Group("/problems")); err != nil {
		log.Fatal(err)
	}

	e.Start(":80")
}