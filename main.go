package main

import (
	"log"

	"flag"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/debug_problem"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/debug_submission"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/rankings"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/submit"
)

func main() {
	// Get dotenv file path
	dotenv_path_flag := flag.String("env", "", "")
	flag.Parse()
	dotenv_path := string(*dotenv_path_flag)

	// Load dotenv file
	err := godotenv.Load(dotenv_path)
	if err != nil {
		log.Print(err)
	}

	e := echo.New()

	// e.Pre(middleware.HTTPSRedirect())
	e.Pre(middleware.AddTrailingSlash())

	db, err := db.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Ping()

	for dpid := 1; dpid <= 28; dpid++ {
		if dpid > 16 && dpid < 23 {
			continue
		}
		row := db.QueryRow("SELECT MIN(diff) FROM debug_submissions WHERE dpid = ? AND result = 'AC'", dpid)
		var mindiff int
		if err := row.Scan(&mindiff); err != nil {
			return
		}
		db.Exec("UPDATE debug_problems SET mindiff = ? WHERE dpid = ?", mindiff, dpid)
	}

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

	listen_addr := os.Getenv("LISTEN_ADDR")
	if err := e.Start(listen_addr); err != nil {
		log.Fatal(err)
	}

	// e.StartTLS(":443", "/cert/cert.pem", "/cert/cert.key")
}
