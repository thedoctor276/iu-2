package iu

import (
	"reflect"
	"sync"
)

var (
	lastComponentID uint64
	componentMutex  sync.Mutex
)

type Component interface {
	ID() string

	View() View

	setView(v View)

	Render() string

	Sync()
}

func nextComponentId() uint64 {
	componentMutex.Lock()
	defer componentMutex.Unlock()

	lastComponentID++
	return lastComponentID
}

func ForRangeComponents(top Component, action func(comp Component)) {
	var interType = reflect.TypeOf((*Component)(nil)).Elem()

	action(top)

	v := reflect.ValueOf(comp).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := f.Type()

		if t.Implements(interType) {
			c := f.Interface().(Component)
			ForRangeComponents(c, action)
		}
	}
}
