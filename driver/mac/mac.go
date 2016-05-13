// +build darwin

package mac

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa -framework WebKit -framework CoreImage
#include "mac.h"
*/
import "C"
import (
	"encoding/json"
	"strconv"
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
	cmenu := menuToCMenu(menu)
	defer freeCMenu(cmenu)

	C.Menu_Set(cmenu)
}

func setDockMenu(menu iu.Menu) {
	cmenu := menuToCMenu(menu)
	defer freeCMenu(cmenu)

	C.Menu_SetDock(cmenu)
}

func menuToCMenu(menu iu.Menu) C.Menu__ {
	cmenu := C.Menu__{
		indent:    C.uint(menu.Indent),
		disabled:  cbool(menu.Disabled),
		separator: cbool(menu.Separator),
	}

	cmenu.name = C.CString(menu.Name)
	cmenu.shortcut = C.CString(menu.Shortcut)
	cmenu.handlerName = C.CString(menu.HandlerName)

	return cmenu
}

func freeCMenu(cmenu C.Menu__) {
	defer C.free(unsafe.Pointer(cmenu.name))
	defer C.free(unsafe.Pointer(cmenu.shortcut))
	defer C.free(unsafe.Pointer(cmenu.handlerName))
	defer C.free(unsafe.Pointer(cmenu.compoID))
}

//export onMenuClick
func onMenuClick(name *C.char) {
	if h, ok := menuHandler(C.GoString(name)); ok {
		h()
	}
}

//export onContextMenuClick
func onContextMenuClick(name *C.char, compoID *C.char) {
	id := iu.ComponentTokenFromString(C.GoString(compoID))
	c := iu.ComponentByID(id)

	if cmc, ok := c.(iu.ContextMenuContainer); ok {
		iu.CallContextMenuHandler(cmc, C.GoString(name))
	}
}

// ============================================================================
// Window
// ============================================================================

func createWindow(ID string, conf iu.WindowConfig) unsafe.Pointer {
	cid := C.CString(ID)
	defer C.free(unsafe.Pointer(cid))

	ctitle := C.CString(conf.Title)
	defer C.free(unsafe.Pointer(ctitle))

	cbackgroundColor := C.CString(conf.BackgroundColor)
	defer C.free(unsafe.Pointer(cbackgroundColor))

	cconf := C.WindowConfig__{
		x:               C.CGFloat(conf.X),
		y:               C.CGFloat(conf.Y),
		width:           C.CGFloat(conf.Width),
		height:          C.CGFloat(conf.Height),
		title:           ctitle,
		backgroundColor: cbackgroundColor,
		backgroundType:  C.uint(conf.BackgroundType),
		borderless:      cbool(conf.Borderless),
		disableResize:   cbool(conf.DisableResize),
		disableClose:    cbool(conf.DisableClose),
		disableMinimize: cbool(conf.DisableMinimize),
	}

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

func renderWindow(ptr unsafe.Pointer, HTML string, baseURL string) {
	cHTML := C.CString(HTML)
	defer C.free(unsafe.Pointer(cHTML))

	cbaseURL := C.CString(baseURL)
	defer C.free(unsafe.Pointer(cbaseURL))

	C.Window_Render(ptr, cHTML, cbaseURL)
}

func renderComponentInWindow(ptr unsafe.Pointer, ID string, component string) {
	cID := C.CString(ID)
	defer C.free(unsafe.Pointer(cID))

	component = strconv.Quote(component)
	ccompo := C.CString(component)
	defer C.free(unsafe.Pointer(ccompo))

	C.Window_RenderComponent(ptr, cID, ccompo)
}

func showContextMenu(ptr unsafe.Pointer, compoID string, menus []iu.Menu) {
	var l int

	if l = len(menus); l == 0 {
		return
	}

	cmenus := make([]C.Menu__, l)

	for i := 0; i < l; i++ {
		m := menuToCMenu(menus[i])
		m.compoID = C.CString(compoID)
		cmenus[i] = m
		defer freeCMenu(cmenus[i])
	}

	C.Window_ShowContextMenu(ptr, unsafe.Pointer(&cmenus[0]), C.int(l))
}

func showWindowAlert(ptr unsafe.Pointer, msg string) {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.Window_Alert(ptr, cmsg)
}

//export onWindowMinimize
func onWindowMinimize(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnMinimize != nil {
		w.OnMinimize()
	}
}

//export onWindowDeminimize
func onWindowDeminimize(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnDeminimize != nil {
		w.OnDeminimize()
	}
}

//export onWindowFullScreen
func onWindowFullScreen(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnFullScreen != nil {
		w.OnFullScreen()
	}
}

//export onWindowExitFullScreen
func onWindowExitFullScreen(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnExitFullScreen != nil {
		w.OnExitFullScreen()
	}
}

//export onWindowMove
func onWindowMove(ID *C.char, x C.CGFloat, y C.CGFloat) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnMove != nil {
		w.OnMove(float64(x), float64(y))
	}
}

//export onWindowResize
func onWindowResize(ID *C.char, width C.CGFloat, height C.CGFloat) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnResize != nil {
		w.OnResize(float64(width), float64(height))
	}
}

//export onWindowFocus
func onWindowFocus(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnFocus != nil {
		w.OnFocus()
	}
}

//export onWindowBlur
func onWindowBlur(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.OnBlur != nil {
		w.OnBlur()
	}
}

//export onWindowClose
func onWindowClose(ID *C.char) C.BOOL {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	shouldClose := true

	if w.OnClose != nil {
		shouldClose = w.OnClose()
	}

	if shouldClose {
		iu.DismountComponent(w.Nav())
		iu.UnregisterDriver(w)
	}

	return cbool(shouldClose)
}

//export onWindowLoad
func onWindowLoad(ID *C.char) {
	id := iu.DriverToken(C.GoString(ID))
	d, _ := iu.DriverByID(id)
	w := d.(*Window)

	if w.Config().OnLoad != nil {
		w.Config().OnLoad()
	}
}

//export onCallEventHandler
func onCallEventHandler(name *C.char, msgJSON *C.char) {
	var msg iu.EventMessage
	var err error

	b := []byte(C.GoString(msgJSON))

	if err = json.Unmarshal(b, &msg); err != nil {
		iulog.Error(err)
		return
	}

	if len(msg.ID) == 0 {
		iulog.Panicf("no ID in %v", msg)
	}

	id := iu.ComponentTokenFromString(msg.ID)
	c := iu.ComponentByID(id)
	iu.CallComponentEvent(c, msg.Name, msg.Arg)
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

func appResourcesPath() string {
	return C.GoString(C.ResourcesPath())
}
