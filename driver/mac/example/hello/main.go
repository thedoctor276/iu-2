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

	hello := &Hello{
		ContextMenu: []iu.Menu{
			iu.Menu{
				Name:     "Custom button",
				Shortcut: "meta+k",
			},
			iu.Menu{Separator: true},
			mac.CtxMenuCut,
			mac.CtxMenuCopy,
			mac.CtxMenuPaste,
		},
	}

	p := iu.NewPage(hello, iu.PageConfig{
		CSS: []string{"hello.css"},
	})

	win.Navigate(p)
	return win
}
