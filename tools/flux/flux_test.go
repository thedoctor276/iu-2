package flux

import "testing"

func TestRegisterStore(t *testing.T) {
	foo := NewFooStore()

	RegisterStore(foo)
	defer UnregisterStore(foo)

	if id := foo.DispatchToken(); id == 0 {
		t.Error("id should not be 0")
	}

	if l := len(mainDispatcher.callbacks); l != 1 {
		t.Errorf("l should be 1: %v", l)
	}
}

func TestRegisterRegisteredStore(t *testing.T) {
	foo := NewFooStore()

	RegisterStore(foo)
	defer UnregisterStore(foo)

	RegisterStore(foo)
}

func TestUnregisterStore(t *testing.T) {
	foo := NewFooStore()

	RegisterStore(foo)
	UnregisterStore(foo)

	if l := len(mainDispatcher.callbacks); l != 0 {
		t.Errorf("l should be 0: %v", l)
	}
}

func TestUnregisterNonRegisteredStore(t *testing.T) {
	foo := NewFooStore()
	UnregisterStore(foo)
}

func TestWaitFor(t *testing.T) {
	foo := NewFooStore()
	bar := NewBarStore()
	bar.Foo = foo

	RegisterStore(foo)
	RegisterStore(bar)
	defer UnregisterStore(foo)
	defer UnregisterStore(bar)

	DispatchAction("test-wait-for")
}

func TestDispatch(t *testing.T) {
	Dispatch(Payload{})
}

func TestDispatchAction(t *testing.T) {
	DispatchAction("test-dispatch-action")
}

func TestDispatchActionWithData(t *testing.T) {
	DispatchActionWithData("test-dispatch-action-data", 42)
}
