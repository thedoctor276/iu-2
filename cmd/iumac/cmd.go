package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
)

func RunCmd(command string, args ...string) (err error) {
	var cmd *exec.Cmd
	var stdout io.ReadCloser
	var stderr io.ReadCloser

	cmd = exec.Command(command, args...)

	if stdout, err = cmd.StdoutPipe(); err != nil {
		return
	}

	if stderr, err = cmd.StderrPipe(); err != nil {
		return
	}

	go PrintStdout(stdout)
	go PrintStderr(stderr)

	if err = cmd.Start(); err != nil {
		return
	}

	err = cmd.Wait()
	return
}

func PrintStdout(stdout io.Reader) {
	var reader = bufio.NewReader(stdout)
	var data []byte
	var err error

	for {
		if data, err = reader.ReadBytes('\n'); err != nil {
			if err == io.EOF {
				return
			}

			log.Fatal(err)
		}

		os.Stdout.Write(data)
	}
}

func PrintStderr(stdout io.Reader) {
	var reader = bufio.NewReader(stdout)
	var data []byte
	var err error

	for {
		if data, err = reader.ReadBytes('\n'); err != nil {
			if err == io.EOF {
				return
			}

			log.Fatal(err)
		}

		os.Stderr.Write(data)
	}
}
