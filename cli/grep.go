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
			options := grep.Options{Colored: c.Bool("c"), Context: c.Int("C"), Folder: c.String("d"), File: c.String("f"), Pattern: c.String("pattern")}
			waitErr, err := grep.Execute(options)
			if waitErr != nil {
				fmt.Println(fmt.Sprintf("Grep error output (probably match not found): %v", waitErr))
				os.Exit(1)
			}
			return err
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "pattern, p", Usage: "Pattern to match in files", Required: true},
			cli.StringFlag{Name: "directory, d", Usage: "Search in directory recursively"},
			cli.StringFlag{Name: "file, f", Usage: "Search in files"},
			cli.StringFlag{Name: "format, t", Usage: "Format of the searched files (none or gz)", Value: "gz"},
			cli.BoolFlag{Name: "colored, c", Usage: "Color the matches"},
			cli.IntFlag{Name: "context, C", Usage: "Show context lines of the matches (number)"},
		},
	}
}
