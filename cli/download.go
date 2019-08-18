package cli

import "github.com/urfave/cli"

// DownloadLogsCommand command for downloading logs from cloud storage
func DownloadLogsCommand() cli.Command {
	return cli.Command{
		Name:  "download",
		Usage: "Download logs recursively from cloud storage",
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{},
	}
}
