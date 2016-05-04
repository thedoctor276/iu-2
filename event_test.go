package iu

import (
	"encoding/json"
	"testing"

	"github.com/maxence-charriere/iu-log"
)

func TestCallComponentEvent(t *testing.T) {
	c := &Foo{}
	CallComponentEvent(c, "OnDismount", "")
}

func TestCallComponentEventAsField(t *testing.T) {
	c := &Foo{
		OnTest: func() {
			iulog.Print("OnTest")
		},
	}

	CallComponentEvent(c, "OnTest", "")
}

func TestCallComponentEventAsNilField(t *testing.T) {
	c := &Foo{}
	CallComponentEvent(c, "OnTest", "")
}

func TestCallComponentUnknownEvent(t *testing.T) {
	defer func() { recover() }()

	c := &Foo{}
	CallComponentEvent(c, "SillyTest", "")
	t.Error("should have panic")
}

func TestCallComponentNonEventt(t *testing.T) {
	defer func() { recover() }()

	c := &Foo{}
	CallComponentEvent(c, "Prefix", "")
	t.Error("should have panic")
}

func TestCallComponentEventWithEventArg(t *testing.T) {
	c := &Foo{}

	b, err := json.Marshal(MouseEvent{ClientX: 42})
	if err != nil {
		t.Fatal(err)
	}

	CallComponentEvent(c, "OnClick", string(b))
}

func TestCallComponentEventWithInvalidEventArg(t *testing.T) {
	defer func() { recover() }()

	c := &Foo{}
	CallComponentEvent(c, "OnClick", "J'aime les filles !")
	t.Error("should have panic")
}
