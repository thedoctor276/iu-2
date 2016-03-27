package iu

import (
	"fmt"
	"testing"
)

// Hello
type Hello struct {
	Greeting *World
	Input    string
}

func (hello Hello) Template() string {
	return `
<div id={{.ID}}>
    {{.Greeting.Render}}
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnInputChanged" "event"}}">
</div>
    `
}

// World
type World struct {
	Greeting  string
	OnNothing func()
}

func (World *World) OnChange(e KeyboardEvent) {
	fmt.Println("OnChange", e)
}

func (World *World) OnClick(e MouseEvent) {
	fmt.Println("OnClick", e)
}

func (World *World) OnWheel(e WheelEvent) {
	fmt.Println("OnWheel", e)
}

func (World *World) OnChecked(e bool) {
	fmt.Println("OnChecked", e)
}

func (World *World) OnString(e string) {
	fmt.Println("OnString", e)
}

func (World *World) OnNumber(e float64) {
	fmt.Println("OnNumber", e)
}

func (world *World) Template() string {
	return `<h1 id={{.ID}}>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>`
}

// WrongWorld
type WrongWorld struct {
	WrongGreeting string
}

func (world WrongWorld) Template() string {
	return `<h1 id={{.ID}}>Hello, <span onclick="{{.OnEvent}}>{{.Greeting}}</span></h1>`
}

// HelloLoop
type HelloLoop struct {
	Greeting *World
	Parent   *HelloLoop
	Input    string
}

func (hello *HelloLoop) Template() string {
	return `
<div id={{.ID}}>
    {{.Greeting.Render}}
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnInputChanged" "event"}}">
</div>
    `
}

// Tests
func TestComponentRender(t *testing.T) {
	hello := Hello{
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

func TestNewComponent(t *testing.T) {
	lastComponentID = 0
	defer func() { lastComponentID = 0 }()

	hello := Hello{}
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
