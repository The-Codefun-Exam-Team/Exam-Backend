package main

import (
	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
)

func main(){
	e := echo.New()

	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("exam.codefun.vn")
	// e.AutoTLSManager.Cache = autocert.DirCache("/autocert/.cache")
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	e.GET("/ping", general.Ping)

	e.Start(":1700")
	// e.StartAutoTLS(":21700")
}