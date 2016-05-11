package iu

import (
	"testing"

	"github.com/maxence-charriere/iu-log"
)

// ============================================================================
// Foo
// ============================================================================

type Foo struct {
	Number int
	Prefix string
	OnTest func()
}

func (f *Foo) Template() string {
	return `<span>{{if .Prefix}}{{.Prefix}}-{{end}}Foo({{.Number}})</span>`
}

func (f *Foo) OnMount() {
	iulog.Printf("%p ~> I'm mounted", f)
}

func (f *Foo) OnDismount() {
	iulog.Printf("%p ~> I'm dismounted", f)
}

func (f *Foo) ContextMenu() []Menu {
	return []Menu{
		Menu{
			Name: "Minimize",
		},

		Menu{
			Name: "Quit",
			Handler: func() {
				iulog.Printf("%p quit", f)
			},
		},
	}
}

func (f *Foo) OnClick(e MouseEvent) {
	iulog.Printf("%p ~> OnClick %v", f, e)
}

// ============================================================================
// Bar
// ============================================================================

type Bar struct {
	Foo  *Foo
	Foos []*Foo
	Fook map[string]*Foo
}

func (b *Bar) Template() string {
	return `
<div>
    {{.Foo.Render}}
    <span>Bar</span>
	{{range .Foos}}
	{{.Render}}
	{{end}}
</div>  
`
}

func (b *Bar) unexportedComponent() *Foo {
	return nil
}

type EmptyFoo struct {
}

func (f *EmptyFoo) Template() string {
	return `<span>I'm an empty foo</span>`
}

// ============================================================================
// InvalidFoo
// ============================================================================

type InvalidFoo struct {
	Number int
}

func (f *InvalidFoo) Template() string {
	return `
<div>{{.SexyFoo% }}</div>
`
}

// ============================================================================
// InvalidBar
// ============================================================================

type InvalidBar struct {
	Slice []int
}

func (f *InvalidBar) Template() string {
	return `
<div>{{index .Slice 42}}</div>
`
}

// ============================================================================
// Tests
// ============================================================================

func TestRenderComponent(t *testing.T) {
	f := &Foo{}
	b := &Bar{
		Foo: f,
	}

	d := NewDriverTest(b, DriverConfig{})
	defer d.Close()

	RenderComponent(f)
}

func TestComponentRender(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	c := &Foo{
		Prefix: "Super",
		Number: 42,
	}

	MountComponent(c, d)
	defer DismountComponent(c)

	ic := innerComponent(c)
	t.Log(ic.Render())
}

func TestComponentRenderTree(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	f := &Foo{}
	b := &Bar{
		Foo: f,
	}

	MountComponent(b, d)
	defer DismountComponent(b)

	ic := innerComponent(b)
	t.Log(ic.Render())
}

func TestComponentRenderTreeWithSlice(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	f := &Foo{}
	foos := []*Foo{
		&Foo{},
		&Foo{},
		&Foo{},
	}

	b := &Bar{
		Foo:  f,
		Foos: foos,
	}

	MountComponent(b, d)
	defer DismountComponent(b)

	ic := innerComponent(b)
	t.Log(ic.Render())
}

func TestComponentRenderTreeWithMap(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	f := &Foo{}
	foos := map[string]*Foo{
		"abra":     &Foo{},
		"kadabra":  &Foo{},
		"alakazam": &Foo{},
	}

	b := &Bar{
		Foo:  f,
		Fook: foos,
	}

	MountComponent(b, d)
	defer DismountComponent(b)

	ic := innerComponent(b)
	t.Log(ic.Render())
}

func TestComponentRenderInvalidField(t *testing.T) {
	defer func() { recover() }()

	c := &InvalidBar{}

	d := NewDriverTest(c, DriverConfig{})
	defer d.Close()

	RenderComponent(c)
	t.Error("should have panic")
}

func TestNewComponent(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	c := &Foo{}
	newComponent(c, d)
}

func TestNewComponentEmpty(t *testing.T) {
	defer func() { recover() }()

	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	c := &EmptyFoo{}
	newComponent(c, d)
	t.Error("should have panic")
}

func TestNewComponentInvalidTemplate(t *testing.T) {
	defer func() { recover() }()

	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	c := &InvalidFoo{}
	newComponent(c, d)
	t.Error("should have panic")
}

func TestPropertyMapRaiseEvent(t *testing.T) {
	m := propertyMap{
		"ID": ComponentToken(42),
	}

	js := m.RaiseEvent("Onclick", "event")
	t.Log(js)
}

func TestPropertyMapRaiseEventWithMultipleArgs(t *testing.T) {
	m := propertyMap{
		"ID": ComponentToken(42),
	}

	js := m.RaiseEvent("Onclick", "event", "name")
	t.Log(js)
}
