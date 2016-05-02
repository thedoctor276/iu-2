package iu

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"sync"

	"github.com/maxence-charriere/iu-log"
)

var (
	currentComponentID ComponentToken
	componentMtx       sync.Mutex
)

// Component is the representation of a component.
type Component interface {
	// Template returns a string containing the HTML representation of the component.
	// The string must have a template format compatible with go package text/template.
	Template() string
}

// MountHandler is the representation of a component which can perform
// an action when it is mounted.
type MountHandler interface {
	// OnMount is the method to be called when the component is mounted.
	OnMount()
}

// UnmountHandler is the representation of a component which can
// perform an action when it is unmounted.
type UnmountHandler interface {
	// OnUnmount is the method to be called when the component is unmounted.
	OnUnmount()
}

// ComponentToken is an unique identifier for a component.
type ComponentToken uint

// ForRangeComponent performs an action on all components in the tree starting by root.
func ForRangeComponent(root Component, action func(c Component)) {
	action(root)

	v := reflect.ValueOf(root)
	ct := v.Type()
	v = reflect.Indirect(v)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Type.Implements(ct) {
			c := v.Field(i).Interface().(Component)
			ForRangeComponent(c, action)
		}
	}
}

// RenderComponent renders a component.
// Should be called when a component needs to be redrawn.
func RenderComponent(c Component) {
	d := DriverByComponent(c)
	d.RenderComponent(c)
}

func nextComponentToken() ComponentToken {
	componentMtx.Lock()
	defer componentMtx.Unlock()

	currentComponentID++
	return currentComponentID
}

type component struct {
	Driver    Driver
	Component Component
	ID        ComponentToken
	Template  *template.Template
}

func (c *component) Render() string {
	var b bytes.Buffer

	v := reflect.ValueOf(c.Component)
	ct := v.Type()
	v = reflect.Indirect(v)
	t := v.Type()
	m := propertyMap{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Type.Implements(ct) {
			c := v.Field(i).Interface().(Component)
			m[f.Name] = innerComponent(c)
		} else if f.Type == reflect.TypeOf((*string)(nil)).Elem() {
			s := v.Field(i).Interface().(string)
			s = template.HTMLEscapeString(s)
			m[f.Name] = HTMLEntities(s)
		} else {
			m[f.Name] = v.Field(i).Interface()
		}
	}

	m["ID"] = c.ID

	if err := c.Template.Execute(&b, m); err != nil {
		iulog.Panic(err)
	}

	return b.String()
}

func newComponent(c Component, d Driver) *component {
	r := c.Template()

	if !strings.Contains(r, `data-iu-id="{{.ID}}"`) {
		r = strings.Replace(r, " ", ` data-iu-id="{{.ID}}" `, 1)
	}

	tpl, err := template.New("").Parse(r)
	if err != nil {
		iulog.Panic(err)
	}

	return &component{
		Driver:    d,
		Component: c,
		ID:        nextComponentToken(),
		Template:  tpl,
	}
}

type propertyMap map[string]interface{}

func (m propertyMap) RaiseEvent(name string, arg ...string) string {
	a := "event"

	if len(arg) != 0 {
		a = arg[0]
	}

	if len(arg) > 1 {
		iulog.Warn("RaiseEvent(name string, arg ...string)  currently support only 1 arg, the others will be ignored")
	}

	return fmt.Sprintf("CallEventHandler('%v', '%v', %v)", m["ID"], name, a)
}
