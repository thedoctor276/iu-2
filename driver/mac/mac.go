// +build darwin

package mac

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa -framework WebKit
#include "mac.h"
*/
import "C"
import (
	"unsafe"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

var (
	dockPtr unsafe.Pointer
)

// ============================================================================
// App
// ============================================================================

func init() {
	dockPtr = C.App_Init()
}

func runApp() {
	C.App_Run()
}

func quitApp() {
	C.App_Quit()
}

//export onLaunch
func onLaunch() {
	if OnLaunch != nil {
		OnLaunch()
	}
}

//export onReopen
func onReopen() {
	if OnReopen != nil {
		OnReopen()
	}
}

//export onQuit
func onQuit() {
	if OnQuit != nil {
		OnQuit()
	}
}

// ============================================================================
// Menu
// ============================================================================

func setMenu(menu iu.Menu) {
	cmenu := C.Menu__{
		indent:    C.uint(menu.Indent),
		enabled:   cbool(menu.Enabled),
		separator: cbool(menu.Separator),
	}

	cmenu.name = C.CString(menu.Name)
	defer C.free(unsafe.Pointer(cmenu.name))

	cmenu.shortcut = C.CString(menu.Shortcut)
	defer C.free(unsafe.Pointer(cmenu.shortcut))

	cmenu.nativeAction = C.CString(menu.NativeAction)
	defer C.free(unsafe.Pointer(cmenu.nativeAction))

	C.Menu_Set(cmenu)
}

func setDockMenu(menu iu.Menu) {
	cmenu := C.Menu__{
		indent:    C.uint(menu.Indent),
		enabled:   cbool(menu.Enabled),
		separator: cbool(menu.Separator),
	}

	cmenu.name = C.CString(menu.Name)
	defer C.free(unsafe.Pointer(cmenu.name))

	cmenu.shortcut = C.CString(menu.Shortcut)
	defer C.free(unsafe.Pointer(cmenu.shortcut))

	cmenu.nativeAction = C.CString(menu.NativeAction)
	defer C.free(unsafe.Pointer(cmenu.nativeAction))

	C.Menu_SetDock(cmenu)
}

//export onMenuClick
func onMenuClick(name *C.char) {
	if h, ok := menuHandler(C.GoString(name)); ok {
		h()
	}
}

// ============================================================================
// Window
// ============================================================================

func createWindow(ID string, conf WindowConfig) unsafe.Pointer {
	cid := C.CString(ID)
	defer C.free(unsafe.Pointer(cid))

	cconf := C.WindowConfig__{
		x:               C.CGFloat(conf.X),
		y:               C.CGFloat(conf.Y),
		width:           C.CGFloat(conf.Width),
		height:          C.CGFloat(conf.Height),
		title:           C.CString(conf.Title),
		borderless:      cbool(conf.Borderless),
		disableResize:   cbool(conf.DisableResize),
		disableClose:    cbool(conf.DisableClose),
		disableMinimize: cbool(conf.DisableMinimize),
	}
	defer C.free(unsafe.Pointer(cconf.title))

	return unsafe.Pointer(C.Window_Create(cid, cconf))
}

func showWindow(ptr unsafe.Pointer) {
	C.Window_Show(ptr)
}

func moveWindow(ptr unsafe.Pointer, x float64, y float64) {
	C.Window_Move(ptr, C.CGFloat(x), C.CGFloat(y))
}

func centerWindow(ptr unsafe.Pointer) {
	C.Window_Center(ptr)
}

func resizeWindow(ptr unsafe.Pointer, width float64, height float64) {
	C.Window_Resize(ptr, C.CGFloat(width), C.CGFloat(height))
}

func closeWindow(ptr unsafe.Pointer) {
	C.Window_Close(ptr)
}

//export onWindowMinimize
func onWindowMinimize(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnMinimize != nil {
		win.OnMinimize()
	}
}

//export onWindowDeminimize
func onWindowDeminimize(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnDeminimize != nil {
		win.OnDeminimize()
	}
}

//export onWindowFullScreen
func onWindowFullScreen(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnFullScreen != nil {
		win.OnFullScreen()
	}
}

//export onWindowExitFullScreen
func onWindowExitFullScreen(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnExitFullScreen != nil {
		win.OnExitFullScreen()
	}
}

//export onWindowMove
func onWindowMove(ID *C.char, x C.CGFloat, y C.CGFloat) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnMove != nil {
		win.OnMove(float64(x), float64(y))
	}
}

//export onWindowResize
func onWindowResize(ID *C.char, width C.CGFloat, height C.CGFloat) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnResize != nil {
		win.OnResize(float64(width), float64(height))
	}
}

//export onWindowFocus
func onWindowFocus(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnFocus != nil {
		win.OnFocus()
	}
}

//export onWindowBlur
func onWindowBlur(ID *C.char) {
	win, err := WindowByID(C.GoString(ID))

	if err != nil {
		iulog.Panic(err)
	}

	if win.OnBlur != nil {
		win.OnBlur()
	}
}

//export onWindowClose
func onWindowClose(ID *C.char) C.BOOL {
	id := C.GoString(ID)
	win, err := WindowByID(id)

	if err != nil {
		iulog.Panic(err)
	}

	shouldClose := true

	if win.OnClose != nil {
		shouldClose = win.OnClose()
	}

	if shouldClose {
		delete(windows, id)
	}

	return cbool(shouldClose)
}

// ============================================================================
// Util
// ============================================================================

func cbool(b bool) C.BOOL {
	if b {
		return 1
	}

	return 0
}
