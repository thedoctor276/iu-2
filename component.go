package iu

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

var (
	lastComponentID uint64
	componentMutex  sync.Mutex
)

type Component struct {
	page     *Page
	view     View
	id       string
	template *template.Template
}

func (comp *Component) ID() string {
	return comp.id
}

func (comp *Component) Render() string {
	var buffer bytes.Buffer
	var err error

	viewValue := reflect.Indirect(reflect.ValueOf(comp.view))
	viewType := viewValue.Type()
	viewInterfaceType := reflect.TypeOf((*View)(nil)).Elem()

	m := viewMap{}

	for i := 0; i < viewType.NumField(); i++ {
		fieldName := viewType.Field(i).Name
		fieldType := viewType.Field(i).Type

		if fieldType.Implements(viewInterfaceType) {
			view := viewValue.Field(i).Interface().(View)
			m[fieldName] = compoM.Component(view)
		} else {
			m[fieldName] = viewValue.Field(i).Interface()
		}
	}

	m["ID"] = comp.id

	if comp.template == nil {
		if comp.template, err = template.New("").Parse(comp.view.Template()); err != nil {
			iulog.Panic(err)
		}
	}

	if err = comp.template.Execute(&buffer, m); err != nil {
		iulog.Panic(err)
	}

	return buffer.String()
}

func NewComponent(view View) *Component {
	component := &Component{
		view: view,
		id:   fmt.Sprintf("iu-%v", nextComponentId()),
	}

	return component
}

func nextComponentId() uint64 {
	componentMutex.Lock()
	defer componentMutex.Unlock()

	lastComponentID++
	return lastComponentID
}
