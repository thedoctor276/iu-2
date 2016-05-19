#include "_cgo_export.h"
#import <UIKit/UIKit.h>
#include "ios.h"

// ============================================================================
// App
// ============================================================================

void App_Run() {
    @autoreleasepool {
		UIApplicationMain(0, nil, nil, nil);
	}
}

void App_Quit() {
    exit(0)
}