package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	util "github.com/euclid1990/gomworker/utilities"
)

func Stop(c *cli.Context) {
	fmt.Println("Stop")
	s := util.NewServer()
	s.Stop()
}
