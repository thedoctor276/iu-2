package mac

import (
	"fmt"
	"unsafe"

	"github.com/maxence-charriere/iu-log"
)

var (
	windows = map[string]*Window{}
)

type WindowConfig struct {
	X               float64
	Y               float64
	Width           float64
	Height          float64
	Title           string
	Borderless      bool
	DisableResize   bool
	DisableClose    bool
	DisableMinimize bool
}

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
}

func (win *Window) Show() {
	showWindow(win.ptr)
}

func (win *Window) Move(x float64, y float64) {
	moveWindow(win.ptr, x, y)
}

func (win *Window) Center() {
	centerWindow(win.ptr)
}

func (win *Window) Resize(width float64, height float64) {
	resizeWindow(win.ptr, width, height)
}

func (win *Window) Close() {
	closeWindow(win.ptr)
}

func CreateWindow(ID string, conf WindowConfig) *Window {
	if !running {
		iulog.Panic("windows must be created once the app is launched ~> start creating windows in mac.OnLaunch func")
	}

	if _, ok := windows[ID]; ok {
		iulog.Panicf("window with ID %v is already created", ID)
	}

	win := &Window{
		ptr: createWindow(ID, conf),
	}

	windows[ID] = win
	return win
}

func WindowByID(ID string) (window *Window, err error) {
	var ok bool

	if window, ok = windows[ID]; !ok {
		err = fmt.Errorf("window with ID %v is not created", ID)
	}

	return
}
