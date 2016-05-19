package mac

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

var (
	windows = map[string]*Window{}
)

// Window represents a cocoa window driver.
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

// Show shows the window.
func (w *Window) Show() {
	showWindow(w.ptr)
}

// Move moves the window.
func (w *Window) Move(x float64, y float64) {
	moveWindow(w.ptr, x, y)
}

// Center moves the window to the center of the screen.
func (w *Window) Center() {
	centerWindow(w.ptr)
}

// Resize resizes the window.
func (w *Window) Resize(width float64, height float64) {
	resizeWindow(w.ptr, width, height)
}

// RenderComponent renders a component.
func (w *Window) RenderComponent(ID iu.ComponentToken, component string) {
	component = strconv.Quote(component)
	call := fmt.Sprintf(`RenderComponent("%v", %v)`, ID, component)

	w.CallJavascript(call)
}

// ShowContextMenu shows a context menu.
func (w *Window) ShowContextMenu(ID iu.ComponentToken, m []iu.Menu) {
	showContextMenu(w.ptr, fmt.Sprint(ID), m)
}

// CallJavascript calls a javascript function.
func (w *Window) CallJavascript(call string) {
	callJavascriptInWindow(w.ptr, call)
}

// Alert prompts a message.
func (w *Window) Alert(msg string) {
	showWindowAlert(w.ptr, msg)
}

// Close closes the window.
func (w *Window) Close() {
	closeWindow(w.ptr)
}

// NewWindow creates a new window instance.
// Windows should be created only with this function.
func NewWindow(main iu.Component, c iu.DriverConfig) *Window {
	if !running {
		iulog.Panic("windows must be created once the app is launched ~> start creating windows in OnLaunch func")
	}

	w := &Window{
		ptr:        createWindow(string(c.ID), c.Window),
		DriverBase: iu.NewDriverBase(main, c),
	}

	iu.MountComponent(w.Nav(), w)
	iu.RegisterDriver(w)
	renderWindow(w.ptr, w.Render(), iu.ResourcesPath())
	return w
}
