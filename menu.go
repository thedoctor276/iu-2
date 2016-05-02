package iu

import "github.com/maxence-charriere/iu-log"

// ContextMenu is the representation of a component that can display a context menu.
type ContextMenu interface {
	// ContextMenu returns a slice of menu.
	ContextMenu() []Menu

	Component
}

// Menu represents a menu item abstraction.
type Menu struct {
	Name          string
	Shortcut      string
	HandlerName   string
	Indent        uint
	Disabled      bool
	Separator     bool
	NativeHandler bool
	Handler       func()
}

// ShowContextMenu shows the context menu of a component.
func ShowContextMenu(c ContextMenu) {
	ic := innerComponent(c)
	ic.Driver.ShowContextMenu(ic.ID, c.ContextMenu())
}

// CallContextMenuHandler calls menu handler of a component that implements the ContextMenu interface.
// This call should be used only in a driver implementation.
func CallContextMenuHandler(c ContextMenu, name string) {
	for _, m := range c.ContextMenu() {
		if m.Name != name {
			continue
		}

		if m.Handler == nil {
			iulog.Warnf(`menu named "%v" in %#v doesn't have a handler`, name, c)
			return
		}

		m.Handler()
		return
	}
}
