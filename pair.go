package iu

import "reflect"

func PairViewComponent(view View, comp Component) {
	var interType = reflect.TypeOf((*Component)(nil)).Elem()

	view.RegisterComponent(comp)
	comp.setView(view)

	v := reflect.ValueOf(comp).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := f.Type()

		if t.Implements(interType) {
			c := f.Interface().(Component)
			PairViewComponent(view, c)
		}
	}
}

func UnpairViewComponent(view View, comp Component) {
	var interType = reflect.TypeOf((*Component)(nil)).Elem()

	view.UnregisterComponent(comp)
	comp.setView(view)

	v := reflect.ValueOf(comp).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := f.Type()

		if t.Implements(interType) {
			c := f.Interface().(Component)
			UnpairViewComponent(view, c)
		}
	}
}
