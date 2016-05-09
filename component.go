package iu

import (
	"bytes"
	"fmt"
	"reflect"
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

// RenderComponent renders a component.
// Should be called when a component needs to be redrawn.
func RenderComponent(c Component) {
	ic := innerComponent(c)
	d := ic.Driver

	d.RenderComponent(ic.ID, ic.Render())
}

type component struct {
	Driver    Driver
	Component Component
	ID        ComponentToken
	Template  *template.Template
}

func (c *component) Render() string {
	var b bytes.Buffer

	m := propertyMap{}

	v := reflect.ValueOf(c.Component)
	t := v.Type()

	for i := 0; i < v.NumMethod(); i++ {
		method := t.Method(i)
		methodv := v.Method(i)

		if !isComponentTreeGetter(method, methodv) {
			continue
		}

		m[method.Name] = valueForRender(methodv.Call(nil)[0])
	}

	v = reflect.Indirect(v)
	t = v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		if len(f.PkgPath) != 0 {
			continue
		}

		m[f.Name] = valueForRender(v.Field(i))
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
