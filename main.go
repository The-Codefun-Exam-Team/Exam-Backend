package main

import (
	"log"
	
	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
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

	e.Start(":1700")
}