package iu

import "testing"

func TestShowContextMenu(t *testing.T) {
	c := &Foo{}

	d := NewDriverTest(c, DriverConfig{})
	defer d.Close()

	ShowContextMenu(c)
}

func TestCallContextMenuHandler(t *testing.T) {
	c := &Foo{}

	d := NewDriverTest(c, DriverConfig{})
	defer d.Close()

	CallContextMenuHandler(c, "Quit")
}

func TestCallContextMenuWithoutHandler(t *testing.T) {
	c := &Foo{}

	d := NewDriverTest(c, DriverConfig{})
	defer d.Close()

	CallContextMenuHandler(c, "Minimize")
}
