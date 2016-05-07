package iu

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"text/template"

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

// ComponentToken is an unique identifier for a component.
type ComponentToken uint

// ComponentTokenFromString converts a string to a ComponentToken.
func ComponentTokenFromString(s string) ComponentToken {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		iulog.Panic(err)
	}

	return ComponentToken(id)
}

// ForRangeComponent performs an action on all components in the tree starting by root.
//
// Nodes types:
// - Component
// - array
// - slice
// - map
func ForRangeComponent(root Component, action func(c Component)) {
	action(root)

	v := reflect.ValueOf(root)
	v = reflect.Indirect(v)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)

		if !f.CanSet() {
			continue
		}

		forRangeComponentValue(f, action)
	}
}

func forRangeComponentValue(v reflect.Value, action func(c Component)) {
	t := v.Type()

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			forRangeComponentValue(v.Index(i), action)
		}

	case reflect.Map:
		for _, k := range v.MapKeys() {
			forRangeComponentValue(v.MapIndex(k), action)
		}

	default:
		if c, ok := v.Interface().(Component); ok {
			ForRangeComponent(c, action)
		}
	}
}

// RenderComponent renders a component.
// Should be called when a component needs to be redrawn.
func RenderComponent(c Component) {
	ic := innerComponent(c)
	d := ic.Driver

	d.RenderComponent(ic.ID, ic.Render())
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
	v = reflect.Indirect(v)
	t := v.Type()
	m := propertyMap{}

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)

		if !fv.CanSet() {
			continue
		}

		m[f.Name] = valueForRender(fv)
	}

	m["ID"] = c.ID

	if err := c.Template.Execute(&b, m); err != nil {
		iulog.Panic(err)
	}

	return b.String()
}

func valueForRender(v reflect.Value) interface{} {
	switch f := v.Interface().(type) {
	case Component:
		return innerComponent(f)

	case string:
		f = template.HTMLEscapeString(f)
		return HTMLEntities(f)

	default:
		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			s := make([]interface{}, v.Len())

			for i := 0; i < v.Len(); i++ {
				s[i] = valueForRender(v.Index(i))
			}

			return s

		case reflect.Map:
			t := v.Type()
			mt := reflect.MapOf(t.Key(), reflect.TypeOf((*interface{})(nil)).Elem())
			mv := reflect.MakeMap(mt)

			for _, k := range v.MapKeys() {
				val := valueForRender(v.MapIndex(k))
				mv.SetMapIndex(k, reflect.ValueOf(val))
			}

			return mv.Interface()

		default:
			return f
		}
	}
}

func newComponent(c Component, d Driver) *component {
	v := reflect.ValueOf(c)
	v = reflect.Indirect(v)

	if v.NumField() == 0 {
		iulog.Panicf("a component should have at least 1 field: %#v", c)
	}

	r := c.Template()
	if !strings.Contains(r, `data-iu-id="{{.ID}}"`) {
		r = strings.Replace(r, ">", ` data-iu-id="{{.ID}}">`, 1)
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
