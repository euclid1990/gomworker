package main

import (
	"github.com/codegangsta/cli"
	"github.com/euclid1990/gomworker/cmd"
	"github.com/euclid1990/gomworker/utilities"
	"os"
)

func main() {
	// Read env vars from .env file
	utilities.LoadEnv("")
	// Create cli application
	app := cli.NewApp()
	app.Name = "Gomworker"
	app.Version = "1.0.0"
	app.Usage = "Support Laravel running multiple Queue workers."
	app.Commands = cmd.Commands
	app.Run(os.Args)
}
