package mac

import (
	"testing"

	"github.com/maxence-charriere/iu"
)

func TestRegisterMenuHandler(t *testing.T) {
	e := iu.Menu{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	registerMenuHandler(e)
	defer unregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; !ok {
		t.Error("menu handler should have an entry with key", e.Name)
	}
}

func TestRegisterMenuHandlerWithoutHandler(t *testing.T) {
	e := iu.Menu{
		Name: "File/Quit",
	}

	registerMenuHandler(e)
	defer unregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; ok {
		t.Error("menu handler should not have an entry with key", e.Name)
	}
}

func TestUnregisterMenuHandlerWithoutHandler(t *testing.T) {
	e := iu.Menu{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	registerMenuHandler(e)
	unregisterMenuHandler(e.Name)

	if _, ok := menuHandlers[e.Name]; ok {
		t.Error("menu handler should not have an entry with key", e.Name)
	}
}

func TestMenuHandler(t *testing.T) {
	e := iu.Menu{
		Name:    "File/Quit",
		Handler: func() { t.Log("I'm a menu handler") },
	}

	registerMenuHandler(e)
	defer unregisterMenuHandler(e.Name)

	if _, ok := menuHandler(e.Name); !ok {
		t.Error("menu handler should have an entry with key", e.Name)
	}
}
