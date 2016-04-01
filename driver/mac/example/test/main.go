package main

import (
	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
	"github.com/maxence-charriere/iu-osx"
)

func main() {
	mac.SetMenu(mac.MenuQuit)
	mac.SetMenu(mac.MenuClose)
	mac.SetDockMenu(iu.Menu{
		Name:    "Foo",
		Enabled: true,
		Handler: func() { iulog.Warn("Foo handler") },
	})

	mac.OnLaunch = func() { iulog.Warn("OnLaunch") }
	mac.OnReopen = func() { iulog.Warn("OnReopen") }
	mac.OnQuit = func() { iulog.Warn("OnQuit") }

	mac.Run()
}
