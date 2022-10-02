package models

import (
	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

func ReadSubsCode(db *db.DB, rid int) (string, error) {
	var r int
	var code, er string
	row := db.QueryRow("SELECT * FROM subs_code WHERE rid = ?", rid)

	if err := row.Scan(&r, &code, &er); err != nil {
		return code, err
	}

	return code, nil
}