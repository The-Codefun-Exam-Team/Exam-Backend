package debugproblem

import (
	"github.com/labstack/echo/v4"

	"github.com/The-Codefun-Exam-Team/Exam-Backend/db"
)

type Group struct {
	group *echo.Group
	db    *db.DB
}

func New(db *db.DB, g *echo.Group) (*Group, error) {
	grp := &Group{
		group: g,
		db:    db,
	}

	g.GET("/", testHandler)
	g.GET("/:code", grp.ProblemGet)

	return grp, nil
}
