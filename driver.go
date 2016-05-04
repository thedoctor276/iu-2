package iu

import (
	"bytes"
	"text/template"

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
<html lang="{{if .Lang}}{{.Lang}}{{else}}en{{end}}">
<head>
    <title>{{.ID}}</title>
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
{{range .JS}}
<script src="{{.}}"></script>{{end}}
</body>
</html>	
`
)

var (
	drivers = map[DriverToken]Driver{}
)

// Driver is a representation of what handles the rendering of the user interface.
type Driver interface {
	Config() DriverConfig

	RenderComponent(ID ComponentToken, component string)

	ShowContextMenu(ID ComponentToken, m []Menu)

	Alert(msg string)

	Close()
}

// DriverConfig is the configuration required by a driver.
type DriverConfig struct {
	ID     DriverToken
	Lang   string
	CSS    []string
	JS     []string
	OnLoad func()
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

// DriverToken is an identifier for a driver.
type DriverToken string

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

// Config returns the driver configuration.
func (d *DriverBase) Config() DriverConfig {
	return d.config
}

// Root returns the root component.
func (d *DriverBase) Root() Component {
	return d.root
}

// Render renders the whole component tree.
func (d *DriverBase) Render() string {
	var b bytes.Buffer

	m := map[string]interface{}{
		"ID":          d.Config().ID,
		"Lang":        d.Config().Lang,
		"CSS":         d.Config().CSS,
		"JS":          d.Config().JS,
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
	tpl := template.Must(template.New("").Parse(pageTpl))

	if len(c.ID) == 0 {
		c.ID = "Main driver"
	}

	d := &DriverBase{
		config:          c,
		innerComponents: map[Component]*component{},
		components:      map[ComponentToken]Component{},
		root:            root,
		template:        tpl,
	}

	return d
}

// RegisterDriver register a driver by its configuration ID.
// It makes the driver ready fo event handling.
// Should be only used in a driver implementation.
func RegisterDriver(d Driver) {
	c := d.Config()

	if d, ok := drivers[c.ID]; ok {
		iulog.Panicf("a driver with id = %v is already registered: %#v", c.ID, d)
	}

	drivers[c.ID] = d
}

// UnregisterDriver unregisters a driver.
func UnregisterDriver(d Driver) {
	c := d.Config()
	delete(drivers, c.ID)
}

// DriverByID returns a registered driver.
func DriverByID(ID DriverToken) (d Driver, ok bool) {
	d, ok = drivers[ID]
	return
}

// DriverByComponent returns the driver where a component is mounted.
func DriverByComponent(c Component) Driver {
	ic := innerComponent(c)
	return ic.Driver
}
