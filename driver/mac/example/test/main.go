package main

import (
	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
	"github.com/maxence-charriere/iu/driver/mac"
)

func main() {
	mac.SetMenu(mac.MenuQuit)
	mac.SetMenu(mac.MenuClose)
	mac.SetDockMenu(iu.Menu{
		Name:    "Foo",
		Enabled: true,
		Handler: func() { iulog.Warn("Foo handler") },
	})

	mac.OnLaunch = onLaunch
	mac.OnReopen = onReopen
	mac.OnQuit = func() { iulog.Warn("OnQuit") }

	mac.Run()
}

func onLaunch() {
	iulog.Warn("OnLaunch")

	win := mac.CreateWindow("Main", mac.WindowConfig{
		Width:  1240,
		Height: 720,
	})

	win.OnMinimize = func() { iulog.Warn("OnMinimize") }
	win.OnDeminimize = func() { iulog.Warn("OnDeminimize") }
	win.OnFullScreen = func() { iulog.Warn("OnFullScreen") }
	win.OnExitFullScreen = func() { iulog.Warn("OnExitFullScreen") }
	win.OnMove = func(x float64, y float64) { iulog.Warnf("OnMove (%v, %v)", x, y) }
	win.OnResize = func(width float64, height float64) { iulog.Warnf("OnResize (%v, %v)", width, height) }
	win.OnFocus = func() { iulog.Warn("OnFocus") }
	win.OnBlur = func() { iulog.Warn("OnBlur") }
	win.OnClose = func() bool {
		iulog.Warn("OnClose")
		return true
	}

	win.Show()

	win2 := mac.CreateWindow("Pref", mac.WindowConfig{
		Width:  1240,
		Height: 720,
	})

	win2.Show()
}

func onReopen() {
	var win *mac.Window
	var err error

	iulog.Warn("OnReopen")

	if win, err = mac.WindowByID("Main"); err != nil {
		win = mac.CreateWindow("Main", mac.WindowConfig{
			Width:  1240,
			Height: 720,
		})

		iulog.Warn("Window recreated")
	}

	win.Show()
}
