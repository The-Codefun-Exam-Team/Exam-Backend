package main

import (
	"log"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
	"github.com/The-Codefun-Exam-Team/Exam-Backend/general"
)

func main(){
	e := echo.New()

	cf_db, err := db.New("unknown:password@tcp(localhost:3306)/codefun")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := cf_db.Query("SELECT * FROM subs_code LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var rid int
		var code string
		var er string

		if err := rows.Scan(&rid, &code, &er); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("rid: %v, code: %v, er: %v", rid, code, er)
	}

	e.GET("/ping", general.Ping)
	e.Start(":1700")
}