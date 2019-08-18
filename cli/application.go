package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// StartApplication holder of cli commands, can start cli application
func StartApplication(version string, gitRevString string) {
	app := cli.NewApp()
	app.Name = "cslogs"
	app.Usage = "Tool for download and grep in logs on cloud storage"
	app.EnableBashCompletion = true
	app.UsageText = "cslogs command [command options] [arguments...]"
	if len(version) > 0 {
		app.Version = version
	} else {
		app.Version = "0.1.0"
	}
	if len(gitRevString) > 0 {
		app.Version = app.Version + fmt.Sprintf(" (git short hash: %v)", gitRevString)
	}
	app.Email = "oleewere@gmail.com"
	app.Author = "Oliver Mihaly Szabo"
	app.Copyright = "Copyright 2019 Oliver Mihaly Szabo"
	app.Commands = []cli.Command{}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
