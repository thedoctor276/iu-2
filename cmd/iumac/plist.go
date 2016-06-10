package main

import (
	"os"
	"path"
	"text/template"

	"github.com/maxence-charriere/iu/tools/config"
)

var (
	PlistFilename = "Info.Plist"

	PlistTpl = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>   
    <key>CFBundleDevelopmentRegion</key>
	<string>{{.CFBundleDevelopmentRegion}}</string>
    <key>CFBundleVersion</key>
	<string>{{.CFBundleVersion}}</string>
    <key>CFBundleName</key>
	<string>{{.CFBundleName}}</string>
    <key>CFBundleDisplayName</key>
	<string>{{.CFBundleDisplayName}}</string>
    <key>CFBundleExecutable</key>
	<string>{{.CFBundleExecutable}}</string>
	<key>CFBundleIdentifier</key>
	<string>{{.CFBundleIdentifier}}</string>
    <key>CFBundlePackageType</key>
	<string>{{.CFBundlePackageType}}</string>
    <key>CFBundleSignature</key>
	<string>{{.CFBundleSignature}}</string>
    <key>CFBundleInfoDictionaryVersion</key>
	<string>{{.CFBundleInfoDictionaryVersion}}</string>
    <key>CFBundleSupportedPlatforms</key>
    <array>
	<string>{{.CFBundleSupportedPlatforms}}</string>
	</array>
    <key>DTCompiler</key>
	<string>{{.DTCompiler}}</string>
    <key>DTSDKName</key>
	<string>{{.DTSDKName}}</string>
    <key>DTXcode</key>
	<string>{{.DTXcode}}</string>
    <key>DTXcodeBuild</key>
	<string>{{.DTXcodeBuild}}</string>
    <key>DTPlatformBuild</key>
	<string>{{.DTPlatformBuild}}</string>
    <key>DTPlatformVersion</key>
	<string>{{.DTPlatformVersion}}</string>
    <key>DTPlatformName</key>
	<string>{{.DTPlatformName}}</string>
    <key>DTSDKBuild</key>
	<string>{{.DTSDKBuild}}</string>
    <key>LSMinimumSystemVersion</key>
	<string>{{.LSMinimumSystemVersion}}</string>
    <key>CFBundleIconFile</key>
	<string>{{.CFBundleIconFile}}</string>
    <key>LSUIElement</key>
	<string>{{.LSUIElement}}</string>
    <key>NSPrincipalClass</key>
	<string>NSApplication</string>
	<key>NSAppTransportSecurity</key>
	<dict>
        <key>NSAllowsArbitraryLoads</key>
        <true/>
    </dict>
	<key>com.apple.security.app-sandbox</key>
    <true/>
    <key>com.apple.security.inherit</key>
    <true/>
</dict>
</plist>
`
)

type Plist struct {
	CFBundleDevelopmentRegion     string
	CFBundleVersion               string
	CFBundleName                  string
	CFBundleDisplayName           string
	CFBundleExecutable            string
	CFBundleIdentifier            string
	CFBundlePackageType           string
	CFBundleSignature             string
	CFBundleInfoDictionaryVersion string
	CFBundleSupportedPlatforms    string
	DTCompiler                    string
	DTSDKName                     string
	DTXcode                       string
	DTXcodeBuild                  string
	DTPlatformBuild               string
	DTPlatformVersion             string
	DTPlatformName                string
	DTSDKBuild                    string
	LSMinimumSystemVersion        string
	CFBundleIconFile              string
	LSUIElement                   string
}

func (plist Plist) Save(contentsPath string) error {
	filename := path.Join(contentsPath, PlistFilename)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	tpl, err := template.New("Plist").Parse(PlistTpl)
	if err != nil {
		return err
	}

	return tpl.Execute(file, plist)
}

func MakePlist(contentsPath string, conf config.App) Plist {
	return Plist{
		CFBundleDevelopmentRegion:     conf.Mac.DevelopmentRegion,
		CFBundleVersion:               conf.Version,
		CFBundleName:                  conf.Name,
		CFBundleDisplayName:           conf.Name,
		CFBundleExecutable:            ExecName,
		CFBundleIdentifier:            conf.Mac.Identifier,
		CFBundlePackageType:           "APPL",
		CFBundleSignature:             "????",
		CFBundleInfoDictionaryVersion: "6.0",
		CFBundleSupportedPlatforms:    "MacOSX",
		DTCompiler:                    "com.apple.compilers.llvm.clang.1_0",
		DTSDKName:                     "macosx10.11",
		DTXcode:                       "0721",
		DTXcodeBuild:                  "7C1002",
		DTPlatformBuild:               "7C1002",
		DTPlatformVersion:             "GM",
		DTPlatformName:                "macosx10.11",
		DTSDKBuild:                    "15C43",
		LSMinimumSystemVersion:        conf.Mac.DeploymentTarget,
		CFBundleIconFile:              conf.Mac.Icon,
		LSUIElement:                   "true",
	}
}
