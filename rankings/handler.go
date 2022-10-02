package rankings

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (g *Group) RankingsGet(c echo.Context) error {
	var rows *sql.Rows
	var err error

	str_groupid := c.FormValue("group")
	groupid, err := strconv.Atoi(str_groupid)
	if err != nil {
		return err
	}

	str_pageid := c.FormValue("pageid")
	pageid, err := strconv.Atoi(str_pageid)
	if err != nil {
		return err
	}

	str_limit := c.FormValue("limit")
	limit, err := strconv.Atoi(str_limit)
	if err != nil {
		return err
	}

	if groupid > 0 {
		rows, err = g.db.Query(`WITH score_table AS (SELECT tid, dpid, MAX(score) AS score FROM debug_submissions GROUP BY tid, dpid)
		SELECT teams.group, teams.tid, teams.teamname, teams.name, SUM(score_table.score) AS score FROM score_table
		INNER JOIN teams ON teams.tid = score_table.tid
		WHERE teams.group = ?
		GROUP BY score_table.tid ORDER BY score DESC
		LIMIT ? OFFSET ?`, groupid, limit, (pageid-1)*limit)
	} else {
		rows, err = g.db.Query(`WITH score_table AS (SELECT tid, dpid, MAX(score) AS score FROM debug_submissions GROUP BY tid, dpid)
		SELECT teams.group, teams.tid, teams.teamname, teams.name, SUM(score_table.score) AS score FROM score_table 
		INNER JOIN teams ON teams.tid = score_table.tid
		GROUP BY score_table.tid ORDER BY score DESC
		LIMIT ? OFFSET ?`, limit, (pageid-1)*limit)
	}

	_ = rows

	return c.NoContent(http.StatusOK)
}
