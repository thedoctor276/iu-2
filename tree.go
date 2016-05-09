package iu

import "reflect"

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
	t := v.Type()

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)

		if !isComponentNodeGetter(t.Method(i), m) {
			continue
		}

		forRangeComponentValue(m.Call(nil)[0], action)
	}

	v = reflect.Indirect(v)
	t = v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if len(f.PkgPath) != 0 {
			continue
		}

		forRangeComponentValue(v.Field(i), action)
	}
}

func isComponentNodeGetter(m reflect.Method, v reflect.Value) bool {
	if len(m.PkgPath) != 0 {
		return false
	}

	t := v.Type()

	if t.NumIn() != 0 {
		return false
	}

	if t.NumOut() != 1 {
		return false
	}

	switch ot := t.Out(0); ot.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		return true

	default:
		if ct := reflect.TypeOf((*Component)(nil)).Elem(); !ot.Implements(ct) {
			return false
		}
	}

	return true
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
