package main

import (
	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu/driver/mac"
)

func main() {
	mac.SetMenu(mac.MenuQuit)
	mac.SetMenu(mac.MenuCut)
	mac.SetMenu(mac.MenuCopy)
	mac.SetMenu(mac.MenuPaste)
	mac.SetMenu(mac.MenuSelectAll)
	mac.SetMenu(mac.MenuClose)

	mac.OnLaunch = onLaunch
	mac.OnReopen = onReopen

	mac.Run()
}

func onLaunch() {
	win := newMainWindow()
	win.Show()
}

func onReopen() {
	d, ok := iu.DriverByID("Main")
	if !ok {
		d = newMainWindow()
	}

	d.(*mac.Window).Show()
}

func newMainWindow() *mac.Window {
	hello := &Hello{}

	return mac.NewWindow(hello, iu.DriverConfig{
		ID:  "Main",
		CSS: []string{"hello.css"},
		Window: iu.WindowConfig{
			Width:      1240,
			Height:     720,
			Background: iu.WindowBackgroundDark,
		},
	})
}
