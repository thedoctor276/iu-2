package iu

import "testing"

func TestRegisterMenuHandler(t *testing.T) {
	e := MenuElement{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	RegisterMenuHandler(e)
	defer UnregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; !ok {
		t.Error("menu handler should have an entry with key", e.Name)
	}
}

func TestRegisterMenuHandlerWithoutHandler(t *testing.T) {
	e := MenuElement{
		Name: "File/Quit",
	}

	RegisterMenuHandler(e)
	defer UnregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; ok {
		t.Error("menu handler should not have an entry with key", e.Name)
	}
}

func TestUnregisterMenuHandlerWithoutHandler(t *testing.T) {
	e := MenuElement{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	RegisterMenuHandler(e)
	UnregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; ok {
		t.Error("menu handler should not have an entry with key", e.Name)
	}
}

func TestMenuHandler(t *testing.T) {
	e := MenuElement{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	RegisterMenuHandler(e)
	defer UnregisterMenuHandler(e.Name)

	if _, ok := MenuHandler(e.Name); !ok {
		t.Error("menu handler should have an entry with key", e.Name)
	}
}
