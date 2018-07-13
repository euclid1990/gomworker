package utilities

import (
	"fmt"
	cf "github.com/euclid1990/gomworker/configs"
	"github.com/go-sql-driver/mysql"
	squirrel "gopkg.in/Masterminds/squirrel.v1"
	"strings"
)

type Worker struct {
	Id         int
	Status     string
	Queue      string
	Once       int
	Delay      int
	Force      int
	Memory     int
	Sleep      int
	Timeout    int
	Tries      int
	Started_at mysql.NullTime
	Created_at mysql.NullTime
	Updated_at mysql.NullTime
	Deleted_at mysql.NullTime
}

func CreateWorkersTable(Db *Database) bool {
	if Db.CheckTableExists(cf.DB_TBL_WORKERS) {
		return false
	}
	q, _ := GetFileContent(cf.DB_SQL_PATH)
	subQ := strings.Split(q, ";")

	for _, s := range subQ {
		if NotBlankLine(s) {
			_, err := Db.Exec(s)
			if err != nil {
				panic(err)
			}
		}
	}
	return true
}

/**
 * Treat input to function as variadic
 * GetWorkers(Db, []int{1, 2, 3}...)
 **/
func GetWorkers(Db *Database, ids ...int) []*Worker {
	var workers []*Worker
	p := Db.Stmt().Select("*").From(cf.DB_TBL_WORKERS).Where(cf.DB_REC_SOFTDELETES)

	if len(ids) > 0 {
		p = p.Where(squirrel.Eq{"id": ids})
	}
	rows, _ := p.Query()
	defer rows.Close()

	for rows.Next() {
		worker := &Worker{}
		if err := rows.Scan(
			&worker.Id,
			&worker.Status,
			&worker.Queue,
			&worker.Once,
			&worker.Delay,
			&worker.Force,
			&worker.Memory,
			&worker.Sleep,
			&worker.Timeout,
			&worker.Tries,
			&worker.Started_at,
			&worker.Created_at,
			&worker.Updated_at,
			&worker.Deleted_at,
		); err != nil {
			fmt.Println(err)
		}
		fmt.Println(*worker)
		workers = append(workers, worker)
	}
	fmt.Println(workers)
	return workers
}
