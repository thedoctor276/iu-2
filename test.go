package iu

import (
	"fmt"

	"github.com/maxence-charriere/iu-log"
)

// DriverTest is a driver to be used only for test purpose.
type DriverTest struct {
	*DriverBase
}

// RenderComponent emulates a RenderComponent() call.
func (d *DriverTest) RenderComponent(ID ComponentToken, component string) string {
	c := fmt.Sprintf("rendering %v: %v", ID, component)
	iulog.Print(c)
	return c
}

// ShowContextMenu emulates a ShowContextMenu() call.
func (d *DriverTest) ShowContextMenu(ID ComponentToken, m []Menu) {
	iulog.Printf("showing %#v for %v", m, ID)
}

// CallJavascript emulates a CallJavascript() call.
func (d *DriverTest) CallJavascript(call string) {
	iulog.Printf("call js: %v", call)
}

// Alert emulates an Alert() call.
func (d *DriverTest) Alert(msg string) {
	iulog.Printf("alert %v", msg)
}

// Close closes the driver.
func (d *DriverTest) Close() {
	iulog.Printf("driver %p is closing", d)
	DismountComponent(d.main)
	UnregisterDriver(d)
}

// NewDriverTest create a new test driver.
// Should be used only for test purpose.
func NewDriverTest(root Component, c DriverConfig) *DriverTest {
	d := &DriverTest{
		DriverBase: NewDriverBase(root, c),
	}

	MountComponent(d.Nav(), d)
	RegisterDriver(d)
	return d
}
