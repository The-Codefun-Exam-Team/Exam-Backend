package main

import (
	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
)

func main(){
	e := echo.New()

	e.GET("/ping", general.Ping)
	e.GET("/debug_problem", general.TempDebug)

	e.Start(":1700")
}