package iu

import "testing"

func TestViewMapOnEvent(t *testing.T) {
	expected := "CallEventHandler(this.id, 'OnClick', event)"
	m := viewMap{}

	if js := m.OnEvent("OnClick", "event"); js != expected {
		t.Errorf("js should be %v: %v", expected, js)
	}
}

func TestForRangeView(t *testing.T) {
	hello := Hello{
		Greeting: &World{},
	}

	ForRangeViews(hello, func(v View) error {
		return RegisterView(v)
	})

	defer ForRangeViews(hello, func(v View) error {
		return UnregisterView(v)
	})
}

func TestForRangeViewLoop(t *testing.T) {
	hello := &HelloLoop{
		Greeting: &World{},
	}

	hello.Parent = hello

	ForRangeViews(hello, func(v View) error {
		return RegisterView(v)
	})

	defer ForRangeViews(hello, func(v View) error {
		return UnregisterView(v)
	})
}

func TestSyncView(t *testing.T) {
	v := Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	SyncView(v)
}

func TestSyncViewWithNoLoadedPage(t *testing.T) {
	defer func() { recover() }()

	v := Hello{
		Greeting: &World{},
	}

	RegisterView(v)
	SyncView(v)
	t.Error("should have panic")
}

func TestSyncViewWithtoutContext(t *testing.T) {
	defer func() { recover() }()

	v := Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{})
	defer p.Close()

	SyncView(v)
	t.Error("should have panic")
}
