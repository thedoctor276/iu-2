// Package ios is the IOS driver for iu framework.
package ios

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
// This call on ios is not a clean termination, it should not be used.
func Quit() {
	quitApp()
}
