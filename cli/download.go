package cli

import (
	"github.com/urfave/cli"
)

// DownloadLogsCommand command for downloading logs from cloud storage
func DownloadLogsCommand() cli.Command {
	return cli.Command{
		Name:  "download",
		Usage: "Download logs recursively from cloud storage",
		Action: func(c *cli.Context) error {
			//s3.DownloadFiles("default", "oszabo", "cluster-logs/datahub/oli-distrox", "salt", "/Users/oliverszabo/Downloads/s3-download")
			return nil
		},
		Flags: []cli.Flag{},
	}
}
