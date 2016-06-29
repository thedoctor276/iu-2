package iu

import (
	"bytes"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

const (
	pageTpl = `
<!doctype html>
<html lang="{{if .Lang}}{{.Lang}}{{else}}en{{end}}">
<head>
    <title>{{.ID}}</title>
    <meta charset="utf-8"> 
    
    <style media="all" type="text/css">
        html {
            height: 100%;
            width: 100%;
            margin: 0;
        }
        
        body {
            height: 100%;
            width: 100%;
            margin: 0;
            font-family: "Helvetica Neue", "Segoe UI";
        }
		
		#iu-nav {
			height: 100%;
            width: 100%;
		}
    </style>
    
{{range .CSS}}
    <link type="text/css" rel="stylesheet" href="{{.}}" />{{end}}
</head>
<body oncontextmenu="event.preventDefault()">
{{.Main.Render}}

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

	Nav() Navigation

	RenderComponent(ID ComponentToken, component string) string

	ShowContextMenu(ID ComponentToken, m []Menu)

	CallJavascript(call string)

	Alert(msg string)

	Close()
}

// DriverBase is the base implemetation of a driver.
// It should be embedded in any driver.
type DriverBase struct {
	OnLoaded func()

	config          DriverConfig
	innerComponents map[Component]*component
	components      map[ComponentToken]Component
	main            Navigation
	template        *template.Template
}

// Config returns the driver configuration.
func (d *DriverBase) Config() DriverConfig {
	return d.config
}

// Nav returns the Navigation component.
func (d *DriverBase) Nav() Navigation {
	return d.main
}

// Render renders the whole component tree.
func (d *DriverBase) Render() string {
	var b bytes.Buffer

	m := map[string]interface{}{
		"ID":          d.Config().ID,
		"Lang":        d.Config().Lang,
		"CSS":         d.Config().CSS,
		"JS":          d.Config().JS,
		"Main":        innerComponent(d.main),
		"FrameworkJS": frameworkJS,
	}

	d.template.Execute(&b, m)
	return b.String()
}

// NewDriverBase create an instance of a DriverBase and mounts all its components.
func NewDriverBase(main Component, c DriverConfig) *DriverBase {
	tpl := template.Must(template.New("").Parse(pageTpl))

	if len(c.ID) == 0 {
		c.ID = "Main driver"
	}

	d := &DriverBase{
		config:          c,
		innerComponents: map[Component]*component{},
		components:      map[ComponentToken]Component{},
		main:            newNavigation(main),
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
