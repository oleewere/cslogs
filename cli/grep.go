package cli

import (
	"fmt"
	"os"

	"github.com/oleewere/cslogs/grep"
	"github.com/urfave/cli"
)

// GrepCommand for grepping in text files
func GrepCommand() cli.Command {
	return cli.Command{
		Name:  "grep",
		Usage: "Grep in text files",
		Action: func(c *cli.Context) error {
			waitErr, err := grep.Execute(grep.Options{})
			if waitErr != nil {
				fmt.Println(fmt.Sprintf("Grep error output: %v", waitErr))
				os.Exit(1)
			}
			return err
		},
		Flags: []cli.Flag{},
	}
}
