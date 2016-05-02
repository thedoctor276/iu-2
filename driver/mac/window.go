package mac

import (
	"fmt"
	"unsafe"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

const (
	WindowBackgroundSolid WindowBackground = iota
	WindowBackgroundLight
	WindowBackgroundUltraLight
	WindowBackgroundDark
	WindowBackgroundUltraDark
)

var (
	windows = map[string]*Window{}
)

type WindowBackground uint

type WindowConfig struct {
	X               float64
	Y               float64
	Width           float64
	Height          float64
	Title           string
	Background      WindowBackground
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
	OnNavigate       func()

	currentPage *iu.Page
	ptr         unsafe.Pointer
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

func (win *Window) CurrentPage() *iu.Page {
	return win.currentPage
}

func (win *Window) Navigate(page *iu.Page) {
	if page != win.currentPage && win.currentPage != nil {
		win.currentPage.Close()
	}

	win.currentPage = page
	page.Context = win
	win.Show()
	navigateInWindow(win.ptr, page.Render(), iu.Path())
}

func (win *Window) InjectComponent(component *iu.Component) {
	injectComponentInWindow(win.ptr, component.ID(), component.Render())
}

func (win *Window) ShowContextMenu(menus []iu.Menu, compoID string) {
	showContextMenu(win.ptr, menus, compoID)
}

func (win *Window) Alert(msg string) {
	showWindowAlert(win.ptr, msg)
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
