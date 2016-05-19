// +build darwin
// +build arm arm64

package ios

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#include "ios.h"
*/
import "C"

// ============================================================================
// App
// ============================================================================

func runApp() {
	C.App_Run()
}

func quitApp() {
	C.App_Quit()
}
