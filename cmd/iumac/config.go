package main

import (
	"fmt"
	"os"

	"github.com/maxence-charriere/iu/tools/config"
)

func GetConfig() (conf config.App, err error) {
	if !config.Exists() {
		conf, err = CreateConfig()
		return
	}

	conf, err = config.Read()
	return
}

func CreateConfig() (conf config.App, err error) {
	conf = config.App{
		Name:    ExecName,
		Version: "1",
		Mac: config.Mac{
			DevelopmentRegion: "en",
			Identifier:        fmt.Sprintf("%v.%v", os.Getenv("USER"), ExecName),
			DeploymentTarget:  "10.11",
		},
	}

	err = config.Save(conf)
	return
}
