package main

import (
	"github.com/codegangsta/cli"
	"github.com/euclid1990/gomworker/cmd"
	cf "github.com/euclid1990/gomworker/configs"
	util "github.com/euclid1990/gomworker/utilities"
	"os"
)

func init() {
	// Read env vars from .env file
	util.LoadEnv("")
	// Init variable from
	cf.Database()
}

func main() {
	// Create cli application
	app := cli.NewApp()
	app.Name = "Gomworker"
	app.Version = "1.0.0"
	app.Usage = "Support Laravel running multiple Queue workers."
	app.Commands = cmd.Commands
	app.Run(os.Args)
}
