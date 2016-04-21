package flux

import (
	"testing"

	"github.com/maxence-charriere/iu-log"
)

type FooStore struct {
	*StoreBase
}

func (f *FooStore) OnDispatch(p Payload) {
	iulog.Print("FooStore handles OnDispatch(p Payload)")
}

func NewFooStore() *FooStore {
	return &FooStore{
		StoreBase: NewStoreBase(),
	}
}

type BarStore struct {
	Foo *FooStore
	*StoreBase
}

func (b *BarStore) OnDispatch(p Payload) {
	WaitFor(b.Foo)
	iulog.Print("BarStore handles OnDispatch(p Payload)")
}

func NewBarStore() *BarStore {
	return &BarStore{
		StoreBase: NewStoreBase(),
	}
}

func TestStoreBaseDispatchToken(t *testing.T) {
	s := NewFooStore()
	s.setDispatchToken(42)

	if id := s.DispatchToken(); id != 42 {
		t.Error("is should be 42:", id)
	}
}

func TestStoreBaseAddListener(t *testing.T) {
	s := NewFooStore()
	l := func(e Event) {}
	s.AddListener(l)

	if l := len(s.listeners); l != 1 {
		t.Errorf("l should be 1: %v", l)
	}
}

func TestStoreBaseRemoveListener(t *testing.T) {
	s := NewFooStore()
	l := func(e Event) {}
	id := s.AddListener(l)
	s.RemoveListener(id)

	if l := len(s.listeners); l != 0 {
		t.Errorf("l should be 0: %v", l)
	}
}

func TestStoreBaseRemoveNonexistentListener(t *testing.T) {
	s := NewFooStore()
	s.RemoveListener(42)
}

func TestStoreBaseEmit(t *testing.T) {
	emited := false
	s := NewFooStore()

	l := func(e Event) {
		emited = true
	}

	s.AddListener(l)
	s.Emit("Test")

	if !emited {
		t.Errorf("emited should be %v: %v", true, emited)
	}
}
