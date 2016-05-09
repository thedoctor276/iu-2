package iu

import "testing"

func TestForRangeComponent(t *testing.T) {
	c := &Bar{
		Foo: &Foo{},
	}

	ForRangeComponent(c, func(c Component) {
		t.Logf("%#v", c)
	})
}
