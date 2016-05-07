package mac

import (
	"fmt"
	"unsafe"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

var (
	windows = map[string]*Window{}
)

type Window struct {
	OnMinimize       func()
	OnDeminimize     func()
	OnFullScreen     func()
	OnExitFullScreen func()
	OnMove           func(x float64, y float64)
	OnResize         func(width float64, height float64)
	OnFocus          func()
	OnBlur           func()
	OnClose          func() bool

	ptr unsafe.Pointer
	*iu.DriverBase
}

func (w *Window) Show() {
	showWindow(w.ptr)
}

func (w *Window) Move(x float64, y float64) {
	moveWindow(w.ptr, x, y)
}

func (w *Window) Center() {
	centerWindow(w.ptr)
}

func (w *Window) Resize(width float64, height float64) {
	resizeWindow(w.ptr, width, height)
}

func (w *Window) RenderComponent(ID iu.ComponentToken, component string) {
	renderComponentInWindow(w.ptr, fmt.Sprint(ID), component)
}

func (w *Window) ShowContextMenu(ID iu.ComponentToken, m []iu.Menu) {
	showContextMenu(w.ptr, fmt.Sprint(ID), m)
}

func (w *Window) Alert(msg string) {
	showWindowAlert(w.ptr, msg)
}

func (w *Window) Close() {
	closeWindow(w.ptr)
}

func NewWindow(root iu.Component, c iu.DriverConfig) *Window {
	if !running {
		iulog.Panic("windows must be created once the app is launched ~> start creating windows in OnLaunch func")
	}

	w := &Window{
		ptr:        createWindow(string(c.ID), c.Window),
		DriverBase: iu.NewDriverBase(root, c),
	}

	iu.MountComponent(w.Root(), w)
	iu.RegisterDriver(w)
	iulog.Warn(w.Render())
	renderWindow(w.ptr, w.Render(), iu.ResourcesPath())
	return w
}
