package cli

import "github.com/urfave/cli"

// GrepCommand for grepping in text files
func GrepCommand() cli.Command {
	return cli.Command{
		Name:  "grep",
		Usage: "Grep in text files",
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{},
	}
}
