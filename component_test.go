package iu

import "testing"

func TestComponentRender(t *testing.T) {
	hello := &Hello{
		Greeting: &World{},
	}

	RegisterView(hello)
	RegisterView(hello.Greeting)
	defer UnregisterView(hello)
	defer UnregisterView(hello.Greeting)

	comp := compoM.Component(hello)
	t.Log(comp.Render())
}

func TestComponentRenderInvalidTemplate(t *testing.T) {
	defer func() { recover() }()

	world := WrongWorld{}

	RegisterView(world)
	defer UnregisterView(world)

	comp := NewComponent(world)
	t.Log(comp.Render())
	t.Error("should have panic")
}

func TestComponentMustBeUsable(t *testing.T) {
	v := &Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	c := compoM.Component(v)
	c.MustBeUsable()
}

func TestComponentIsNotUsablePage(t *testing.T) {
	defer func() { recover() }()

	v := &Hello{
		Greeting: &World{},
	}

	RegisterView(v)
	c := compoM.Component(v)
	c.MustBeUsable()
	t.Error("should have panic")
}

func TestComponentIsNotUsableContext(t *testing.T) {
	defer func() { recover() }()

	v := &Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{})
	defer p.Close()

	c := compoM.Component(v)
	c.MustBeUsable()
	t.Error("should have panic")
}

func TestNewComponent(t *testing.T) {
	lastComponentID = 0
	defer func() { lastComponentID = 0 }()

	hello := &Hello{}
	comp := NewComponent(hello)

	if id := comp.ID(); id != "iu-1" {
		t.Error("id should be iu-1:", id)
	}

	if h := comp.view; h != hello {
		t.Errorf("h should be %v: %v", hello, h)
	}
}

func TestNextComponentId(t *testing.T) {
	lastComponentID = 0
	defer func() { lastComponentID = 0 }()

	nextComponentId()

	if lastComponentID != 1 {
		t.Errorf("lastComponentID should be 1: %v", lastComponentID)
	}
}
