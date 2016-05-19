package iu

import "testing"

func TestDriverTest(t *testing.T) {
	c := &Foo{}
	d := NewDriverTest(c, DriverConfig{})
	d.CallJavascript("Alert('pouette')")
	d.Alert("test")
}
