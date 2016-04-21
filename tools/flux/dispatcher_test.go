package flux

import "testing"

func TestDispatcherRegister(t *testing.T) {
	disp := newDispatcher()
	disp.Register(func(p Payload) {})

	if l := len(disp.callbacks); l != 1 {
		t.Errorf("l should be 1: %v", l)
	}
}

func TestDispatcherUnregister(t *testing.T) {
	disp := newDispatcher()
	ID := disp.Register(func(p Payload) {})
	disp.Unregister(ID)

	if l := len(disp.callbacks); l != 0 {
		t.Errorf("l should be 0: %v", l)
	}
}

func TestDispatcherUnregisterNonRegistered(t *testing.T) {
	disp := newDispatcher()
	ID := DispatchToken(42)
	disp.Unregister(ID)
}

func TestDispatcherWaitFor(t *testing.T) {
	disp := newDispatcher()
	fooID := DispatchToken(777)

	foo := func(p Payload) {
		t.Log("foo called")
	}

	bar := func(p Payload) {
		t.Log("bar called")
		disp.WaitFor(fooID)
	}

	boom := func(p Payload) {
		t.Log("boom called")
		disp.WaitFor(fooID)
	}

	disp.Register(bar)
	fooID = disp.Register(foo)
	disp.Register(boom)
	disp.Dispatch(Payload{})
}

func TestDispatcherCircularWaitFor(t *testing.T) {
	defer func() { recover() }()

	disp := newDispatcher()
	barID := DispatchToken(999)

	bar := func(p Payload) {
		t.Log("bar called", barID)
		disp.WaitFor(barID)
	}

	barID = disp.Register(bar)
	disp.Dispatch(Payload{})
	t.Error("should panic")
}

func TestDispatcherWaitForNonRegistered(t *testing.T) {
	defer func() { recover() }()

	disp := newDispatcher()

	bar := func(p Payload) {
		disp.WaitFor(4242)
	}

	disp.Register(bar)
	disp.Dispatch(Payload{})
	t.Error("should panic")
}

func TestDispatcherWaitForOutsideDispatch(t *testing.T) {
	defer func() { recover() }()

	disp := newDispatcher()
	disp.WaitFor(84)
	t.Error("should panic")
}

func TestDispatcherDispatch(t *testing.T) {
	p := Payload{Action: "foo"}
	disp := newDispatcher()

	disp.Register(func(p Payload) {
		t.Log(p)
	})

	disp.Dispatch(p)
}

func TestDispatcherDispatchInDispatch(t *testing.T) {
	defer func() { recover() }()

	disp := newDispatcher()

	disp.Register(func(p Payload) {
		disp.Dispatch(p)
	})

	disp.Dispatch(Payload{})
	t.Error("should panic")
}
