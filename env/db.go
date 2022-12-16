package env

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(dsn string) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("mysql", dsn)
	return
}

func getDSN(config *Config) string {
	cfg := mysql.Config{
		User: config.DBUsername,
		Passwd: config.DBPassword,
		Net: "tcp",
		Addr: config.DBAddress,
		DBName: config.DBDatabaseName,
		AllowNativePasswords: true,
		ParseTime: true,
	}

	dsn := cfg.FormatDSN()
	return dsn
}