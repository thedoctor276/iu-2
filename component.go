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
	composer Composer
	id       string
	template *template.Template
}

func (comp *Component) ID() string {
	return comp.id
}

func (comp *Component) Render() string {
	var buffer bytes.Buffer
	var err error

	componentValue := reflect.ValueOf(comp)
	componentType := componentValue.Type()
	composerValue := reflect.Indirect(reflect.ValueOf(comp.composer))
	composerType := composerValue.Type()

	m := composerMap{}

	for i := 0; i < composerType.NumField(); i++ {
		fieldName := composerType.Field(i).Name
		fieldType := composerType.Field(i).Type

		if fieldType == componentType {
			m[fieldName] = composerValue.Field(i).Interface()
		} else {
			m[fieldName] = composerValue.Field(i).Interface()
		}
	}

	m["ID"] = comp.id

	if comp.template == nil {
		if comp.template, err = template.New("").Parse(comp.composer.Template()); err != nil {
			iulog.Panic(err)
		}
	}

	if err = comp.template.Execute(&buffer, m); err != nil {
		iulog.Panic(err)
	}

	return buffer.String()
}

func NewComponent(composer Composer) *Component {
	return &Component{
		composer: composer,
		id:       fmt.Sprintf("iu-%v", nextComponentId()),
	}
}

func nextComponentId() uint64 {
	componentMutex.Lock()
	defer componentMutex.Unlock()

	lastComponentID++
	return lastComponentID
}

// func ForRangeComponents(top *Component, action func(comp *Component)) {
// 	action(top)

// 	componentType := componentValue.Type()
//     composerValue := reflect.ValueOf(comp.composer)

// 	for i := 0; i < componentValue.NumField(); i++ {
// 		f := componentValue.Field(i)
// 		t := f.Type()

// 		if t == componentType {
// 			c := f.Interface().(*Component)
// 			ForRangeComponents(c, action)
// 		}
// 	}
// }
