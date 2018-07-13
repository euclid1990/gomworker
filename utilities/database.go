package utilities

import (
	"database/sql"
	"fmt"
	cf "github.com/euclid1990/gomworker/configs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	squirrel "gopkg.in/Masterminds/squirrel.v1"
	"os"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase() *Database {
	var err error
	db := &Database{}
	fmt.Println(cf.DB_DRIVER)
	os.Getenv("DB_DRIVER")
	if cf.DB_DRIVER == "sqlite3" {
		db.Db, err = sql.Open(cf.DB_DRIVER, cf.DB_CONNECTION)
	} else {
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v",
			cf.DB_USERNAME, cf.DB_PASSWORD, cf.DB_CONNECTION, cf.DB_DATABASE)
		db.Db, err = sql.Open(cf.DB_DRIVER, dsn)
	}
	if err != nil {
		Logf(cf.LOG_CRITICAL, "Unable to connect to the database: %v", err)
	}
	return db
}

func (db *Database) CheckTableExists(tblName string) bool {
	defer func() {
		if r := recover(); r != nil {
			Logf(cf.LOG_CRITICAL, "Unable to check table [%v] exists: %v", tblName, r)
		}
	}()
	var cnt int
	qContent := `
	SELECT COUNT(*)
	FROM information_schema.tables
	WHERE table_schema = '{{.dbName}}'
	    AND table_name = '{{.tblName}}'
	LIMIT 1;
	`
	qData := map[string]interface{}{
		"dbName":  cf.DB_DATABASE,
		"tblName": tblName,
	}
	q, err := ParseStringTemplate(qContent, qData)
	if err != nil {
		panic(err)
	}
	err = db.Db.QueryRow(q).Scan(&cnt)
	if err != nil {
		panic(err)
	}
	return (cnt > 0)
}

func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Db.Exec(query, args...)
}

func (db *Database) Stmt() squirrel.StatementBuilderType {
	stmt := squirrel.StatementBuilder.RunWith(db.Db)
	return stmt
}

func (db *Database) Close() {
	db.Db.Close()
}
