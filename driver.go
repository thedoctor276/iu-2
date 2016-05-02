package iu

import (
	"bytes"
	"html/template"

	"github.com/maxence-charriere/iu-log"
)

const (
	// WindowBackgroundSolid represents a solid background.
	WindowBackgroundSolid WindowBackground = iota

	// WindowBackgroundLight represents a native light background.
	WindowBackgroundLight

	// WindowBackgroundUltraLight represents a native even lighter background.
	WindowBackgroundUltraLight

	// WindowBackgroundDark represents a native dark background.
	WindowBackgroundDark

	// WindowBackgroundUltraDark represents a even darker background.
	WindowBackgroundUltraDark

	pageTpl = `
<!doctype html>
<html lang="{{.Lang}}">
<head>
    <title>{{if .Title}}{{.Title}}{{else}}Golang loves UI :){{end}}</title>
    <meta charset="utf-8" /> 
    
    <style media="all" type="text/css">
        html {
            height: 100%;
            width: 100%;
            margin: 0pt;
            background-color: transparent;
        }
        
        body {
            height: 100%;
            width: 100%;
            margin: 0pt;
            font-family: "Helvetica Neue", "Segoe UI";
            font-size: 11pt;
            overflow-x: hidden;
            overflow-y: hidden;
            background-color: transparent;
        }
    </style>
    
{{range .CSS}}
    <link rel="stylesheet" href="{{.}}" />{{end}}
</head>
<body oncontextmenu="event.preventDefault()">
{{.Root.Render}}

<script>
{{.FrameworkJS}}
</script>
{{range .Config.JS}}
<script src="{{.}}"></script>{{end}}
</body>
</html>	
`
)

// Driver is a representation of what handles the rendering of the user interface.
type Driver interface {
	RenderComponent(c Component)

	ShowContextMenu(ID ComponentToken, m []Menu)

	Alert(msg string)

	Close()
}

// DriverConfig is the configuration required by a driver.
type DriverConfig struct {
	Title  string
	Lang   string
	CSS    []string
	JS     []string
	Window WindowConfig
}

// WindowConfig is the configuration that only apply on window type drivers.
type WindowConfig struct {
	X               float64
	Y               float64
	Width           float64
	Height          float64
	Title           string
	Background      WindowBackground
	Borderless      bool
	DisableResize   bool
	DisableClose    bool
	DisableMinimize bool
}

// WindowBackground set the driver background type.
type WindowBackground uint

// DriverBase is the base implemetation of a driver.
// It should be embedded in any driver.
type DriverBase struct {
	OnLoaded func()

	config          DriverConfig
	innerComponents map[Component]*component
	components      map[ComponentToken]Component
	root            Component
	template        *template.Template
}

func (d *DriverBase) render() string {
	var b bytes.Buffer

	m := map[string]interface{}{
		"Title":       d.config.Title,
		"Lang":        d.config.Lang,
		"CSS":         d.config.CSS,
		"JS":          d.config.JS,
		"Root":        innerComponent(d.root),
		"FrameworkJS": frameworkJS,
	}

	if err := d.template.Execute(&b, m); err != nil {
		iulog.Panic(err)
	}

	return b.String()
}

// NewDriverBase create an instance of a DriverBase and mounts all its components.
func NewDriverBase(root Component, c DriverConfig) *DriverBase {
	tpl, err := template.New("").Parse(pageTpl)
	if err != nil {
		iulog.Panic(err)
	}

	return &DriverBase{
		config:          c,
		innerComponents: map[Component]*component{},
		components:      map[ComponentToken]Component{},
		root:            root,
		template:        tpl,
	}
}

// DriverByComponent returns the driver where a component is mounted.
func DriverByComponent(c Component) Driver {
	ic := innerComponent(c)
	return ic.Driver
}
