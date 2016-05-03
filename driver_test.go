package iu

import (
	"testing"

	"github.com/maxence-charriere/iu-log"
)

type DriverTest struct {
	*DriverBase
}

func (d *DriverTest) RenderComponent(c Component) {
	iulog.Printf("rendering %#v", c)
}

func (d *DriverTest) ShowContextMenu(ID ComponentToken, m []Menu) {
	iulog.Printf("showing %#v for %v", m, ID)
}

func (d *DriverTest) Alert(msg string) {
	iulog.Printf("alert %v", msg)
}

func (d *DriverTest) Close() {
	iulog.Printf("driver %p is closing", d)
	DismountComponents(d.root)
}

func NewDriverTest(root Component, c DriverConfig) *DriverTest {
	d := &DriverTest{
		DriverBase: NewDriverBase(root, c),
	}

	MountComponents(root, d)
	return d
}

func TestNewDriverBase(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	NewDriverBase(root, conf)
}

func TestDriverBaseRender(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	d := NewDriverTest(root, conf)
	defer d.Close()

	t.Log(d.render())
}

func TestDriverBaseRenderWithConf(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{
		Title: "Test",
		Lang:  "fr",
		CSS:   []string{"test.css"},
		JS:    []string{"test.js"},
	}
	d := NewDriverTest(root, conf)
	defer d.Close()

	t.Log(d.render())
}

func TestDriverByComponent(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	d := NewDriverTest(root, conf)
	defer d.Close()

	if driver := DriverByComponent(root); driver != d {
		t.Errorf("driver should be %p: %p", d, driver)
	}
}
