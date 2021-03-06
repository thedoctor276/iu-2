package iu

import "github.com/maxence-charriere/iu-log"

var (
	innerComponents = map[Component]*component{}
	components      = map[ComponentToken]Component{}
)

// MountHandler is the representation of a component which can perform
// an action when it is mounted.
type MountHandler interface {
	// OnMount is the method to be called when the component is mounted.
	OnMount()
}

// DismountHandler is the representation of a component which can
// perform an action when it is dismounted.
type DismountHandler interface {
	// OnDismount is the method to be called when the component is dismounted.
	OnDismount()
}

// MountComponent mounts a component and all its sub components on a driver.
//
// Mounting a component make it ready for event handling.
// This should be used only when creating a component dynamically.
// eg. in a repeater or a list.
//
// Don't forget to call DismountComponents(c Component) when a manually mounted
// component is not required anymore.
// It will prevent memory leak.
func MountComponent(root Component, d Driver) {
	ForRangeComponent(root, func(c Component) {
		mountComponent(c, d)
	})
}

func mountComponent(c Component, d Driver) {
	ic, ok := innerComponents[c]
	if ok {
		iulog.Panicf("component %T is already mounted", c)
		return
	}

	ic = newComponent(c, d)
	innerComponents[c] = ic
	components[ic.ID] = c

	if mh, ok := c.(MountHandler); ok {
		mh.OnMount()
	}
}

// DismountComponent dismounts a component and all its sub components.
// Should be call only on components mounted with MountComponent(c Component).
func DismountComponent(root Component) {
	ForRangeComponent(root, func(c Component) {
		dismountComponent(c)
	})
}

func dismountComponent(c Component) {
	ic, ok := innerComponents[c]
	if !ok {
		iulog.Warnf("can't dismount component %T: component not mounted", c)
		return
	}

	delete(innerComponents, c)
	delete(components, ic.ID)

	if uh, ok := c.(DismountHandler); ok {
		uh.OnDismount()
	}
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
		iulog.Panicf("component %T isn't mounted", c)
	}

	return ic
}
