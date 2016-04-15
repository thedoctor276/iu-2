package iu

import (
	"fmt"
	"testing"
)

// ============================================================================
// Hello
// ============================================================================

type Hello struct {
	Greeting *World
	Input    string
}

func (hello *Hello) Template() string {
	return `
<div id={{.ID}}>
    {{.Greeting.Render}}
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnChange" "event"}}">
</div>
    `
}

// ============================================================================
// World
// ============================================================================

type World struct {
	Greeting    string
	ContextMenu []Menu
	OnNothing   func()
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

// ============================================================================
// WrongWorld
// ============================================================================

type WrongWorld struct {
	WrongGreeting string
}

func (world WrongWorld) Template() string {
	return `<h1 id={{.ID}}>Hello, <span onclick="{{.OnEvent}}>{{.Greeting}}</span></h1>`
}

// ============================================================================
// HelloLoop
// ============================================================================

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

// ============================================================================
// HelloWrongContextMenu
// ============================================================================

type HelloWrongContextMenu struct {
	Input       string
	ContextMenu int
}

func (hello *HelloWrongContextMenu) Template() string {
	return `
<div id={{.ID}}>
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnChange" "event"}}">
</div>
    `
}

// ============================================================================
// HelloWrongContextMenuBis
// ============================================================================

type HelloWrongContextMenuBis struct {
	Input       string
	ContextMenu []int
}

func (hello *HelloWrongContextMenuBis) Template() string {
	return `
<div id={{.ID}}>
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnChange" "event"}}">
</div>
    `
}

// ============================================================================
// Tests
// ============================================================================

func TestViewMapOnEvent(t *testing.T) {
	expected := "CallEventHandler('iu-42', 'OnClick', event)"
	m := newViewMap("iu-42")

	if js := m.OnEvent("OnClick", "event"); js != expected {
		t.Errorf("js should be %v: %v", expected, js)
	}
}

func TestForRangeView(t *testing.T) {
	hello := &Hello{
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
	v := &Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	SyncView(v)
}

func TestShowContextMenu(t *testing.T) {
	v := &World{}

	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	ShowContextMenu(v)
}

func TestShowContextMenuNoSlice(t *testing.T) {
	defer func() { recover() }()

	v := &HelloWrongContextMenu{}
	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	ShowContextMenu(v)

	t.Error("should have panic")
}

func TestShowContextMenuNoSliceMenu(t *testing.T) {
	defer func() { recover() }()

	v := &HelloWrongContextMenuBis{}
	p := NewPage(v, PageConfig{})
	ctx := &EmptyContext{}

	defer p.Close()

	ctx.Navigate(p)
	ShowContextMenu(v)

	t.Error("should have panic")
}
