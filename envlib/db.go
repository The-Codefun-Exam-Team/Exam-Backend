package envlib

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewDB creates a database using the given DSN.
func NewDB(dsn string) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("mysql", dsn)
	return
}

// GetDSN creates a DSN from the given config struct.
func GetDSN(config *Config) string {
	cfg := mysql.Config{
		User:                 config.DBUsername,
		Passwd:               config.DBPassword,
		Net:                  "tcp",
		Addr:                 config.DBAddress, // Format: <host>:<port>
		DBName:               config.DBDatabaseName,
		AllowNativePasswords: true,
		ParseTime:            true, // Parse time from database to time.Time instead of []byte
	}

	dsn := cfg.FormatDSN()
	return dsn
}
