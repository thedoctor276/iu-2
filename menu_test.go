package iu

import "testing"

func TestCallContextMenuHandler(t *testing.T) {
	n := "custom menu"
	v := &World{
		ContextMenu: []Menu{
			Menu{
				Name:    n,
				Handler: func() {},
			},
		},
	}

	RegisterView(v)
	defer UnregisterView(v)

	CallContextMenuHandler(v, n)
}

func TestCallContextMenuWithoutHandler(t *testing.T) {
	n := "custom menu"
	v := &World{
		ContextMenu: []Menu{
			Menu{
				Name: n,
			},
		},
	}

	RegisterView(v)
	defer UnregisterView(v)

	CallContextMenuHandler(v, n)
}

func TestCallNonexistentContextMenuHandler(t *testing.T) {
	n := "custom menu"
	v := &World{
		ContextMenu: []Menu{
			Menu{
				Name: "random",
			},
		},
	}

	RegisterView(v)
	defer UnregisterView(v)

	CallContextMenuHandler(v, n)
}

func TestCallNonexistentContextMenu(t *testing.T) {
	defer func() { recover() }()

	n := "custom menu"
	v := &Hello{}

	RegisterView(v)
	defer UnregisterView(v)

	CallContextMenuHandler(v, n)
	t.Error("should have panic")
}

func TestCallContextMenuInvalid(t *testing.T) {
	defer func() { recover() }()

	n := "custom menu"
	v := &HelloWrongContextMenu{}

	RegisterView(v)
	defer UnregisterView(v)

	CallContextMenuHandler(v, n)
	t.Error("should have panic")
}
