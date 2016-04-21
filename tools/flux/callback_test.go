package flux

import "testing"

func TestCallbackInit(t *testing.T) {
	c := newCallback(func(p Payload) {})
	c.Pending = true
	c.Handled = true
	c.Init()

	if c.Pending {
		t.Errorf("c.Pending should be %v: %v", false, c.Pending)
	}

	if c.Handled {
		t.Errorf("c.Handled should be %v: %v", false, c.Handled)
	}
}

func TestCallbackCall(t *testing.T) {
	c := newCallback(func(p Payload) {})
	p := Payload{}
	c.Call(p)

	if !c.Pending {
		t.Errorf("c.Pending should be %v: %v", true, c.Pending)
	}

	if !c.Handled {
		t.Errorf("c.Handled should be %v: %v", true, c.Handled)
	}
}

func TestNewCallback(t *testing.T) {
	newCallback(func(p Payload) {})
}
