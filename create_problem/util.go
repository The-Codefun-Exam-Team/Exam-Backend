package create

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var checkDuplicatedQuery = `
SELECT

code AS dpcode

FROM debug_problems
WHERE rid = ?
LIMIT 1
`

var getMaxCode = `
SELECT

MAX(code)

FROM debug_problems
WHERE code REGEXP ?
`

func checkDuplicated(db *sqlx.DB, id int) (string, error) {
	var dpcode string
	err := db.Get(&dpcode, checkDuplicatedQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", err
		}
	}
	return dpcode, nil
}

func getSuggestedCode(db *sqlx.DB) (string, error) {
	// Currently hard-coding the regex patterns
	// TODO: Allow accepting pattern from requests

	prefix_pattern := "D"
	number_pattern := "[0-9]+"
	suffix_pattern := ""

	code_pattern := prefix_pattern + number_pattern + suffix_pattern

	var maxcode string
	err := db.Get(&maxcode, getMaxCode, code_pattern)
	if err != nil {
		return "", err
	}

	number_regex := regexp.MustCompile(number_pattern)
	matches := number_regex.FindAllString(maxcode, -1)

	// Assuming there is only 1 matches
	// This part will be changed to support more complex patterns

	crr, err := strconv.Atoi(matches[0])
	if err != nil {
		return "", err
	}

	return prefix_pattern + fmt.Sprintf("%03d", crr+1) + suffix_pattern, nil
}
