package iu

import (
	"reflect"
	"testing"
)

func TestForRangeComponent(t *testing.T) {
	c := &Bar{
		Foo: &Foo{},
	}

	ForRangeComponent(c, func(c Component) {
		t.Logf("%#v", c)
	})
}

func testIsCompomentNodeGetter(name string, c Component) bool {
	v := reflect.ValueOf(c)
	typ := v.Type()
	m, _ := typ.MethodByName(name)
	mv := v.MethodByName(name)

	return isComponentNodeGetter(m, mv)
}

func TestIsCompomentNodeGetter(t *testing.T) {
	name := "CurrentComponent"
	c := &navigation{}

	if !testIsCompomentNodeGetter(name, c) {
		t.Errorf("%v should be a component node getter", name)
	}
}

func TestIsCompomentNodeGetterUnexported(t *testing.T) {
	name := "unexportedComponent"
	c := &Bar{}

	if testIsCompomentNodeGetter(name, c) {
		t.Errorf("%v should not be a component node getter", name)
	}
}

func TestIsCompomentNodeGetterWithArgs(t *testing.T) {
	name := "Go"
	c := &navigation{}

	if testIsCompomentNodeGetter(name, c) {
		t.Errorf("%v should not be a component node getter", name)
	}
}

func TestIsCompomentNodeGetterNotSingleRet(t *testing.T) {
	name := "CanBack"
	c := &navigation{}

	if testIsCompomentNodeGetter(name, c) {
		t.Errorf("%v should not be a component node getter", name)
	}
}

func TestIsCompomentNodeGetterNotNode(t *testing.T) {
	name := "Template"
	c := &navigation{}

	if testIsCompomentNodeGetter(name, c) {
		t.Errorf("%v should not be a component node getter", name)
	}
}
