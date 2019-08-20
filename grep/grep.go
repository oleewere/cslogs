package grep

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"sync"
)

// Options holds grep options
type Options struct {
	Pattern string
	Folder  string
	Colored bool
	Context int
	Format  string
}

// CreateGrepCommand run local grep command
func CreateGrepCommand(options Options) *exec.Cmd {
	command := "zgrep"
	if options.Format == "none" {
		command = "grep"
	}
	commandArgs := make([]string, 0)
	commandArgs = append(commandArgs, "-r")
	if options.Context > 0 {
		commandArgs = append(commandArgs, "-C")
		commandArgs = append(commandArgs, string(options.Context))
	}
	//options.Folder = "/Users/oliverszabo/Downloads/s3-download"
	//options.Pattern = "salt"
	commandArgs = append(commandArgs, options.Pattern)
	commandArgs = append(commandArgs, options.Folder)

	var grepCommand *exec.Cmd

	grepCommand = exec.Command(command, commandArgs...)
	return grepCommand
}

// Execute process grep command based on pre defined options
func Execute(options Options) (error, error) {
	cmd := CreateGrepCommand(options)
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		return err, nil
	}
	if errStdout != nil {
		return nil, errStdout
	}
	if errStderr != nil {
		return nil, errStderr
	}
	return nil, nil
}
