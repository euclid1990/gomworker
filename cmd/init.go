package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	util "github.com/euclid1990/gomworker/utilities"
)

func Init(c *cli.Context) {
	fmt.Println("Init")
	Db := util.NewDatabase()
	defer Db.Close()
	util.CreateWorkersTable(Db)
	util.GetWorkers(Db)
}
