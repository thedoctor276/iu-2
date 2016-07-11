package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	ConfigFilename = "config.json"
)

// App represents the app configuration.
type App struct {
	Name    string
	Version string
	Mac     Mac
}

// Mac represents the app configuration specific to mac.
type Mac struct {
	DevelopmentRegion string
	Identifier        string
	DeploymentTarget  string
	Icon              string
	SandboxMode       bool
	AppRole           string
	UTI               []string
}

// Exists informs about if the config file exists or not.
func Exists() bool {
	if _, err := os.Stat(ConfigFilename); os.IsNotExist(err) {
		return false
	}

	return true
}

// Read reads the configuration file.
func Read() (config App, err error) {
	var data []byte

	if data, err = ioutil.ReadFile(ConfigFilename); err != nil {
		return
	}

	err = json.Unmarshal(data, &config)
	return
}

// Save saves the configuration into a config file.
func Save(config App) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ConfigFilename, data, 0644)
}
