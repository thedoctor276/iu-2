package iu

import "testing"

func TestCompoManagerRegister(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Register(v)

	if _, ok := manager.components[v]; !ok {
		t.Error("manager should have 1 element")
	}
}

func TestCompoManagerRegisterExistentView(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Register(v)
	manager.Register(v)
}

func TestCompoManagerUnregister(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Register(v)
	manager.Unregister(v)

	if _, ok := manager.components[v]; ok {
		t.Error("manager should be empty")
	}
}

func TestCompoManagerUnregisterNonexistent(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Unregister(v)
}

func TestCompoManagerComponent(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Register(v)
	manager.Component(v)
}

func TestCompoManagerNoRegisteredComponent(t *testing.T) {
	defer func() { recover() }()

	v := Hello{}

	manager := newCompoManager()
	manager.Component(v)
	t.Error("should have panic")
}

func TestCompoManagerView(t *testing.T) {
	v := Hello{}

	manager := newCompoManager()
	manager.Register(v)
	c := manager.Component(v)

	manager.View(c.ID())
}

func TestCompoManagerNoRegisteredView(t *testing.T) {
	defer func() { recover() }()

	manager := newCompoManager()
	manager.View("iu-1")
	t.Error("should have panic")
}

func TestRegisterView(t *testing.T) {
	v := Hello{}
	RegisterView(v)
}

func TestUnregisterView(t *testing.T) {
	v := Hello{}
	RegisterView(v)
	UnregisterView(v)
}
