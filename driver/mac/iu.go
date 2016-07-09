// Package mac is the mac OSX driver for iu framework.
package mac

var (
	// OnLaunch is a handler which is called when the app is launched.
	// The call occurs after the call of Run().
	OnLaunch func()

	// OnReopen is a handler which is called when the app is already running
	// and the user try to launch it again.
	OnReopen func()

	// OnOpenFile is a handler which is called when the app is activated
	// by a file open request.
	OnOpenFile func(filename string)

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
