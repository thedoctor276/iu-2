package config

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	var err error

	config := App{
		Name:    "Test",
		Version: "1",
	}

	if err = Save(config); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = os.Remove(ConfigFilename); err != nil {
			t.Fatal(err)
		}
	}()

	if !Exists() {
		t.Fatalf("%v should exist", ConfigFilename)
	}
}

func TestNotExists(t *testing.T) {
	if Exists() {
		t.Fatalf("%v should not exist", ConfigFilename)
	}
}

func TestRead(t *testing.T) {
	var readConfig App
	var err error

	config := App{
		Name:    "Test",
		Version: "1",
	}

	if err = Save(config); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = os.Remove(ConfigFilename); err != nil {
			t.Fatal(err)
		}
	}()

	if readConfig, err = Read(); err != nil {
		t.Fatal(err)
	}

	if readConfig.Name != config.Name {
		t.Errorf("readConfig.Name should be %v: %v", config.Name, readConfig.Name)
	}

	if readConfig.Version != config.Version {
		t.Errorf("readConfig.Version should be %v: %v", config.Version, readConfig.Version)
	}

}

func TestReadFailed(t *testing.T) {
	if _, err := Read(); err == nil {
		t.Error("should error")
	}
}

func TestSave(t *testing.T) {
	config := App{
		Name:    "Test",
		Version: "1",
	}

	if err := Save(config); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(ConfigFilename); err != nil {
		t.Fatal(err)
	}
}

func TestSaveOverride(t *testing.T) {
	config := App{
		Name:    "Test",
		Version: "1",
	}

	if err := Save(config); err != nil {
		t.Fatal(err)
	}

	config.Name = "Test 2"

	if err := Save(config); err != nil {
		t.Fatal(err)
	}

	if err := os.Remove(ConfigFilename); err != nil {
		t.Fatal(err)
	}
}
