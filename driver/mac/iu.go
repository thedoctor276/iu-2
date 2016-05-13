// Package mac is the mac OSX driver for iu framework.
package mac

import "github.com/maxence-charriere/iu"

var (
	// OnLaunch is a handler which is called when the app is launched.
	// The call occurs after the call of Run().
	OnLaunch func()

	// OnReopen is a handler which is called when the app is already running
	// and the user try to launch it again.
	OnReopen func()

	// OnQuit is a handler which is call when the app is about to quit.
	OnQuit func()

	running bool
)

// Run runs the app.
func Run() {
	running = true
	defer func() { running = false }()

	runApp()
}

// Quit quits the app.
func Quit() {
	quitApp()
}

// SetMenu set a menu item in the app menu.
func SetMenu(menu iu.Menu) {
	registerMenuHandler(menu)
	setMenu(menu)
}

// SetDockMenu set a menu item in the app dock menu.
func SetDockMenu(menu iu.Menu) {
	registerMenuHandler(menu)
	setDockMenu(menu)
}
