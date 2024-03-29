package main

import (
	"github.com/oleewere/cslogs/cli"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

func main() {
	cli.StartApplication(Version, GitRevString)
}
