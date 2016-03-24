package iu

import "testing"

// Hello
type Hello struct {
	Greeting *Component
	Input    string
	Name     string
}

func (hello *Hello) Template() string {
	return `
<div id={{.ID}}>
    {{.Name}}
    {{.Greeting.Render}}
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnInputChanged" "event"}}">
</div>
    `
}

func NewHello() *Hello {
	return &Hello{
		Greeting: NewComponent(&World{}),
		Name:     "Maxence",
	}
}

// World
type World struct {
	Greeting string
}

func (world *World) Template() string {
	return `<h1 id={{.ID}}>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>`
}

// WrongWorld
type WrongWorld struct {
	WrongGreeting string
}

func (world *WrongWorld) Template() string {
	return `<h1 id={{.ID}}>Hello, <span onclick="{{.OnEvent}}>{{.Greeting}}</span></h1>`
}

// Tests
func TestComponentRender(t *testing.T) {
	hello := NewHello()
	comp := NewComponent(hello)
	t.Log(comp.Render())
}

func TestComponentRenderInvalidTemplate(t *testing.T) {
	defer func() { recover() }()

	world := &WrongWorld{}
	comp := NewComponent(world)
	t.Log(comp.Render())
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

	if h := comp.composer; h != hello {
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
