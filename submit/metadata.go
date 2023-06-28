package submit

import (
	"github.com/jmoiron/sqlx"
)

var getMetadataQuery = `
SELECT

dpid, 
problems.code,
runs.language,
subs_code.code AS original_code

FROM debug_problems
INNER JOIN runs ON runs.rid = debug_problems.rid
INNER JOIN subs_code ON subs_code.rid = debug_problems.rid
INNER JOIN problems ON problems.pid = debug_problems.pid

WHERE debug_problems.code = ?
`

type Metadata struct {
	Dpid         int    `db:"dpid"`
	PCode        string `db:"code"`
	Language     string `db:"language"`
	OriginalCode string `db:"original_code"`
}

func GetMetadata(db *sqlx.DB, code string) (*Metadata, error) {
	var metadata Metadata
	err := db.Get(&metadata, getMetadataQuery, code)
	return &metadata, err
}
