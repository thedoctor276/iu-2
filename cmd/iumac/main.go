// Package to build a iu application for mac and package it in a .app.
package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	WorkingDir  string
	ExecName    string
	ForceFlag   bool
	VerboseFlag bool
)

func init() {
	var err error

	if WorkingDir, err = os.Getwd(); err != nil {
		log.Panic(err)
	}

	ExecName = filepath.Base(WorkingDir)

	flag.Usage = Usage
	flag.BoolVar(&ForceFlag, "f", false, "force app resources override")
	flag.BoolVar(&VerboseFlag, "v", false, "increase verbosity")
	flag.Parse()
}

func main() {
	var err error

	if err = Build(); err != nil {
		return
	}

	if err = MakeApp(); err != nil {
		log.Panic(err)
	}
}

func Build() (err error) {
	var args = []string{"build"}

	if VerboseFlag {
		args = append(args, "-v")
	}

	err = RunCmd("go", args...)
	return
}

func MakeApp() (err error) {
	conf, err := GetConfig()
	if err != nil {
		return err
	}

	appName := path.Join(WorkingDir, conf.Name+".app")
	if err = os.MkdirAll(appName, os.ModeDir|0755); err != nil {
		return err
	}

	contentsPath := path.Join(appName, "Contents")
	if err := os.MkdirAll(contentsPath, os.ModeDir|0755); err != nil {
		return err
	}

	macOSPath := path.Join(contentsPath, "MacOS")
	if err := os.MkdirAll(macOSPath, os.ModeDir|0755); err != nil {
		return err
	}

	resourcesPath := path.Join(contentsPath, "Resources")
	if err := os.MkdirAll(resourcesPath, os.ModeDir|0755); err != nil {
		return err
	}

	plist := MakePlist(contentsPath, conf)
	if err := plist.Save(contentsPath); err != nil {
		return err
	}

	if err := os.Rename(ExecName, path.Join(macOSPath, conf.Name)); err != nil {
		return err
	}

	resourcesSrc := path.Join(WorkingDir, "resources")
	if _, err := os.Stat(resourcesSrc); err != nil {
		if os.IsNotExist(err) {
			err = nil
		}

		return err
	}

	if err := MakeResources(contentsPath, resourcesSrc); err != nil {
		return err
	}

	if len(conf.Mac.Icon) != 0 {
		return GenerateIcon(resourcesPath, conf.Mac.Icon)
	}
	return nil
}

func MakeResources(contentsPath string, resourcesSrc string) (err error) {
	var args = []string{"-r", "--delete"}

	if !ForceFlag {
		args = append(args, "--update")

	}

	if VerboseFlag {
		args = append(args, "-v")
	}

	args = append(args, resourcesSrc, contentsPath)
	err = RunCmd("rsync", args...)
	return
}
