package configs

import (
	"os"
)

var (
	DB_DRIVER          string
	DB_CONNECTION      string
	DB_DATABASE        string
	DB_USERNAME        string
	DB_PASSWORD        string
	DB_REC_SOFTDELETES = "deleted_at is NULL"
	DB_TBL_WORKERS     = "workers"
	DB_SQL_PATH        = "migrations/gomworker.sql"
)

func Database() {
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_CONNECTION = os.Getenv("DB_CONNECTION")
	DB_DATABASE = os.Getenv("DB_DATABASE")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}
