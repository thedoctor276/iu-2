package mac

import "github.com/maxence-charriere/iu"

var (
	OnLaunch func()
	OnReopen func()
	OnQuit   func()

	running bool
)

func Run() {
	running = true
	defer func() { running = false }()

	runApp()
}

func Quit() {
	quitApp()
}

func SetMenu(menu iu.Menu) {
	registerMenuHandler(menu)
	setMenu(menu)
}

func SetDockMenu(menu iu.Menu) {
	registerMenuHandler(menu)
	setDockMenu(menu)
}
