package iu

import "testing"

func TestMountComponent(t *testing.T) {
	root := &Foo{}

	d := NewDriverTest(root, DriverConfig{})
	defer d.Close()

	c := &Foo{}
	MountComponent(c, d)
	defer DismountComponent(c)
}

func TestMountMountedComponent(t *testing.T) {
	defer func() { recover() }()

	root := &Foo{}

	d := NewDriverTest(root, DriverConfig{})
	defer d.Close()

	MountComponent(root, d)
	t.Error("should have panic")
}

func TestDismountMountedComponent(t *testing.T) {
	root := &Foo{}
	d := NewDriverTest(root, DriverConfig{})
	d.Close()
	DismountComponent(root)
}

func TestComponentByID(t *testing.T) {
	root := &Foo{}

	d := NewDriverTest(root, DriverConfig{})
	defer d.Close()

	ic := innerComponent(root)

	if r := ComponentByID(ic.ID); r != root {
		t.Errorf("r should be %p: %p", root, r)
	}
}

func TestNotMountedComponentByID(t *testing.T) {
	defer func() { recover() }()

	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	ComponentByID(ComponentToken(4242))
	t.Error("should have panic")
}

func TestInnerComponent(t *testing.T) {
	root := &Foo{}

	d := NewDriverTest(root, DriverConfig{})
	defer d.Close()

	if ic := innerComponent(root); ic.Component != root {
		t.Errorf("ic.Component should be %p: %p", root, ic.Component)
	}
}

func TestNonexistentInnerComponent(t *testing.T) {
	defer func() { recover() }()

	innerComponent(&Foo{})
	t.Error("should have panic")
}
