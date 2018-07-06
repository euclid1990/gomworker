package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/euclid1990/gomworker/configs"
)

var Commands = []cli.Command{
	{
		Name:   configs.COMMAND_START,
		Usage:  "Start one or all workers.",
		Action: Start,
	},
	{
		Name:   configs.COMMAND_STOP,
		Usage:  "Stop one or all workers.",
		Action: Stop,
	},
	{
		Name:   configs.COMMAND_RESTART,
		Usage:  "Restart one or all workers.",
		Action: Restart,
	},
	{
		Name:   configs.COMMAND_STATUS,
		Usage:  "Show status for all worker.",
		Action: Status,
	},
}

func Start(c *cli.Context) {
	fmt.Println("Start")
}

func Stop(c *cli.Context) {
	fmt.Println("Stop")
}

func Restart(c *cli.Context) {
	fmt.Println("Restart")
}

func Status(c *cli.Context) {
	fmt.Println("Status")
}
