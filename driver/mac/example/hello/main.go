package main

import (
	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu/driver/mac"
)

func main() {
	mac.SetMenu(mac.MenuQuit)
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
	win, err := mac.WindowByID("Main")

	if err != nil {
		win = newMainWindow()
	}

	win.Show()
}

func newMainWindow() *mac.Window {
	win := mac.CreateWindow("Main", mac.WindowConfig{
		Width:      1240,
		Height:     720,
		Background: mac.WindowBackgroundDark,
	})

	p := iu.NewPage(&Hello{}, iu.PageConfig{
		CSS: []string{"hello.css"},
	})

	win.Navigate(p)
	return win
}
