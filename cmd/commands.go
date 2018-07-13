package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	cf "github.com/euclid1990/gomworker/configs"
	util "github.com/euclid1990/gomworker/utilities"
)

var (
	Db       *util.Database
	Commands = []cli.Command{
		{
			Name:   cf.COMMAND_INIT,
			Usage:  "Init system database.",
			Action: Init,
		},
		{
			Name:   cf.COMMAND_START,
			Usage:  "Start one or all workers.",
			Action: Start,
		},
		{
			Name:   cf.COMMAND_STOP,
			Usage:  "Stop one or all workers.",
			Action: Stop,
		},
		{
			Name:   cf.COMMAND_RESTART,
			Usage:  "Restart one or all workers.",
			Action: Restart,
		},
		{
			Name:   cf.COMMAND_STATUS,
			Usage:  "Show status for all worker.",
			Action: Status,
		},
	}
)

func Restart(c *cli.Context) {
	fmt.Println("Restart")
}

func Status(c *cli.Context) {
	fmt.Println("Status")
}
