package iu

import "testing"

func TestNavigationGo(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	c := &Foo{Prefix: "Foo 1"}

	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	n.Go(c)

	if n.CurrentComponent() != c {
		t.Errorf("current component should be %#v: %#v", c, n.CurrentComponent())
	}

	if l := len(n.history); l != 2 {
		t.Errorf("l should be 2: %v", l)
	}
}

func TestNavigationBack(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	c := &Foo{Prefix: "Foo 1"}

	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	n.Go(c)
	n.Back()

	if n.CurrentComponent() != root {
		t.Errorf("current component should be %#v: %#v", root, n.CurrentComponent())
	}

	if l := len(n.history); l != 2 {
		t.Errorf("l should be 2: %v", l)
	}
}

func TestNavigationBackFail(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	if err := n.Back(); err == nil {
		t.Error("should not be nil")
	}
}

func TestNavigationNext(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	c := &Foo{Prefix: "Foo 1"}

	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	n.Go(c)
	n.Back()
	n.Next()

	if n.CurrentComponent() != c {
		t.Errorf("current component should be %#v: %#v", c, n.CurrentComponent())
	}

	if l := len(n.history); l != 2 {
		t.Errorf("l should be 2: %v", l)
	}
}

func TestNavigationNextFail(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	if err := n.Next(); err == nil {
		t.Error("should not be nil")
	}
}

func TestNavigationGoBackGo(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	c := &Foo{Prefix: "Foo 1"}
	c2 := &Foo{Prefix: "Foo 2"}

	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	n.Go(c)
	n.Back()
	n.Go(c2)

	if n.CurrentComponent() != c2 {
		t.Errorf("current component should be %#v: %#v", c2, n.CurrentComponent())
	}

	if l := len(n.history); l != 2 {
		t.Errorf("l should be 2: %v", l)
	}
}

func TestNavigationGoGoBackGo(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	root := &Foo{Prefix: "root"}
	c := &Foo{Prefix: "Foo 1"}
	c2 := &Foo{Prefix: "Foo 2"}
	c3 := &Foo{Prefix: "Foo 3"}

	n := newNavigation(root)
	MountComponent(n, d)
	defer DismountComponent(n)

	n.Go(c)
	n.Go(c2)
	n.Back()
	n.Go(c3)

	if n.CurrentComponent() != c3 {
		t.Errorf("current component should be %#v: %#v", c3, n.CurrentComponent())
	}

	if l := len(n.history); l != 3 {
		t.Errorf("l should be 3: %v", l)
	}
}

func TestNewNavigation(t *testing.T) {
	c := &Foo{}
	n := newNavigation(c)

	if n.CurrentComponent() != c {
		t.Errorf("current component should be %#v: %#v", c, n.CurrentComponent())
	}

	if n.index != 0 {
		t.Errorf("index should be 0: %v", n.index)
	}
}
