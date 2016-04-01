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
// Util
// ============================================================================

func cbool(b bool) C.BOOL {
	if b {
		return 1
	}

	return 0
}
