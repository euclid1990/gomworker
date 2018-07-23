package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	util "github.com/euclid1990/gomworker/utilities"
)

func Start(c *cli.Context) {
	fmt.Println("Start")
	util.RemoveFiles(".logs/*.log")
	s := util.NewServer()
	s.Start()
}
