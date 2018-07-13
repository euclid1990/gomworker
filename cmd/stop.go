package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	util "github.com/euclid1990/gomworker/utilities"
)

func Stop(c *cli.Context) {
	fmt.Println("Stop")
	m := util.NewMaster()
	m.Stop()
}
