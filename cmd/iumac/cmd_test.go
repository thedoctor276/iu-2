package main

import "testing"

func TestRunCmd(t *testing.T) {
	var err error

	if err = RunCmd("ls", "-la"); err != nil {
		t.Fatal(err)
	}
}
