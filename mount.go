package iu

import "github.com/maxence-charriere/iu-log"

var (
	innerComponents map[Component]*component
	components      map[ComponentToken]Component
)

// MountComponent makes a component ready for event handling.
// This should be used only when creating a component dynamically.
// eg. in a repeater or a list.
//
// Don't forget to call UnmountComponent(c Component) when a manually mounted
// component is not required anymore.
// It will prevent memory leak.
func MountComponent(c Component, d Driver) {
	ic, ok := innerComponents[c]
	if ok {
		iulog.Panicf("component %#v is already mounted", c)
		return
	}

	ic = newComponent(c, d)
	innerComponents[c] = ic
	components[ic.ID] = c

	if mh, ok := c.(MountHandler); ok {
		mh.OnMount()
	}
}

// MountComponents mounts a component and all its sub components on a driver.
func MountComponents(root Component, d Driver) {
	ForRangeComponent(root, func(c Component) {
		MountComponent(c, d)
	})
}

// UnmountComponent unmounts a component.
// Should be call only on components mounted with MountComponent(c Component).
func UnmountComponent(c Component) {
	ic, ok := innerComponents[c]
	if !ok {
		iulog.Warnf("can't dismount component %#v: component not mounted", c)
		return
	}

	delete(innerComponents, c)
	delete(components, ic.ID)

	if uh, ok := c.(UnmountHandler); ok {
		uh.OnUnmount()
	}
}

// UnmountComponents unmounts a component and all its sub components.
func UnmountComponents(root Component) {
	ForRangeComponent(root, func(c Component) {
		UnmountComponent(c)
	})
}

// ComponentByID returns a component by it's ID.
// Should be called only in a driver implementation, once the component is mounted.
func ComponentByID(ID ComponentToken) Component {
	c, ok := components[ID]
	if !ok {
		iulog.Panicf("no component mounted with ID %v", ID)
	}

	return c
}

func innerComponent(c Component) *component {
	ic, ok := innerComponents[c]
	if !ok {
		iulog.Panicf("component %#v isn't mounted", c)
	}

	return ic
}
