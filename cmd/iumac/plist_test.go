package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/maxence-charriere/iu/tools/config"
)

func TestPlistSave(t *testing.T) {
	var contentPath = "."
	var conf config.App
	var err error

	if conf, _ = CreateConfig(); err != nil {
		t.Fatal(err)
	}

	defer os.Remove(config.ConfigFilename)

	plist := MakePlist(contentPath, conf)

	if err = plist.Save(contentPath); err != nil {
		t.Fatal(err)
	}

	os.Remove(PlistFilename)

}

func TestMakePlist(t *testing.T) {
	contentPath := "."

	conf, err := CreateConfig()
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(config.ConfigFilename)

	plist := MakePlist(contentPath, conf)

	if bundleName := ExecName; plist.CFBundleName != bundleName {
		t.Errorf("plist.CFBundleName should be %v: %v", bundleName, plist.CFBundleName)
	}

	if bundleIdentifier := fmt.Sprintf("%v.%v", os.Getenv("USER"), ExecName); plist.CFBundleIdentifier != bundleIdentifier {
		t.Errorf("plist.CFBundleIdentifier should be %v: %v", bundleIdentifier, plist.CFBundleIdentifier)
	}

	if bundleVersion := "1"; plist.CFBundleVersion != bundleVersion {
		t.Errorf("plist.CFBundleVersion should be %v: %v", bundleVersion, plist.CFBundleVersion)
	}

	if bundlePkgType := "APPL"; plist.CFBundlePackageType != bundlePkgType {
		t.Errorf("plist.CFBundlePackageType should be %v: %v", bundlePkgType, plist.CFBundlePackageType)
	}

	if bundleSign := "????"; plist.CFBundleSignature != bundleSign {
		t.Errorf("plist.CFBundleSignature should be %v: %v", bundleSign, plist.CFBundleSignature)
	}

	if bundleExec := ExecName; plist.CFBundleExecutable != bundleExec {
		t.Errorf("plist.CFBundleExecutable should be %v: %v", bundleExec, plist.CFBundleExecutable)
	}
}
