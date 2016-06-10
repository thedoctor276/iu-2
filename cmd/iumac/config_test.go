package main

import (
	"os"
	"testing"

	"github.com/maxence-charriere/iu/tools/config"
)

func TestGetConfig(t *testing.T) {
	var conf config.App
	var err error

	CreateConfig()
	defer os.Remove(config.ConfigFilename)

	if conf, err = GetConfig(); err != nil {
		t.Fatal(err)
	}

	if conf.Name != ExecName {
		t.Fatalf("conf.Name should be %v: %v", ExecName, conf.Name)
	}
}

func TestGetNonexistentConfig(t *testing.T) {
	var conf config.App
	var err error

	defer os.Remove(config.ConfigFilename)

	if conf, err = GetConfig(); err != nil {
		t.Fatal(err)
	}

	if conf.Name != ExecName {
		t.Fatalf("conf.Name should be %v: %v", ExecName, conf.Name)
	}
}

func TestCreateConfig(t *testing.T) {
	if _, err := CreateConfig(); err != nil {
		t.Fatal(err)
	}

	os.Remove(config.ConfigFilename)
}
