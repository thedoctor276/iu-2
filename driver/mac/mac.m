#include "_cgo_export.h"
#include "mac.h"
#import <WebKit/WebKit.h>

// This macro is used to defer the execution of a block of code in the main
// event loop.
#define defer(code) \
  dispatch_async(dispatch_get_main_queue(), ^{ code })

// This macro is used to execute a block of code in the main event loop while
// waiting for it to complete.
#define synchronize(code) \
  dispatch_sync(dispatch_get_main_queue(), ^{ code })

// ============================================================================
// App
// ============================================================================

@implementation AppDelegate
- (AppDelegate*) init {
    self.dock = [[NSMenu alloc] initWithTitle:@""];
    return self;
}

- (void)applicationDidFinishLaunching:(nonnull NSNotification *)aNotification {
    onLaunch();
}

- (BOOL) applicationShouldHandleReopen:(nonnull NSApplication *)sender hasVisibleWindows:(BOOL)flag {
    onReopen();
    return YES;
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
    onQuit();
    return NSTerminateNow;
}

- (NSMenu *)applicationDockMenu:(NSApplication *)sender {
    return self.dock;
}

- (void) onMenuClick:(id)sender {
    MenuItem* item = (MenuItem*)sender;
    onMenuClick((char*)[item.name UTF8String]);
}
@end

void* App_Init() {
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    
    NSApp.mainMenu = [[NSMenu alloc] initWithTitle:@""];
    Menu_GetOrSet(NSApp.mainMenu, @"");
    
    AppDelegate* dock = [[AppDelegate alloc] init];
    NSApp.delegate = dock;
    return (void*)CFBridgingRetain(dock);
}

void App_Run() {
    [NSApp run];
}

void App_Quit() {
    defer(
        [NSApp terminate:nil];
    );
}

// ============================================================================
// Menu
// ============================================================================

@implementation MenuItem
@end

NSMenu* Menu_GetOrSet(NSMenu* base, NSString* name) {
    NSMenuItem* item = [base itemWithTitle: name];
    
    if (item == nil) {
        item = [[NSMenuItem alloc] initWithTitle:name
                                          action:nil
                                   keyEquivalent:@""];
        NSMenu* menu = [[NSMenu alloc] initWithTitle:name];
        item.submenu = menu;
        [base addItem: item];
    }
    
    return item.submenu;
}

void Menu_Set(Menu__ menu) {
    NSString* name = [NSString stringWithUTF8String: menu.name];
    
    NSMenu* nsmenu = NSApp.mainMenu;    
    NSArray* path = [name componentsSeparatedByString:@"/"];
    
    for (int i = 0; i < path.count - 1; ++i) {        
        nsmenu = Menu_GetOrSet(nsmenu, path[i]);
    }
    
    if (nsmenu == NSApp.mainMenu) {
        nsmenu = [NSApp.mainMenu itemAtIndex: 0].submenu;
    }
    
    NSString* title = (NSString*)path.lastObject;
    Menu_SetMenuItem(nsmenu, menu, title);
}

void Menu_SetDock(Menu__ menu) {
    NSString* name = [NSString stringWithUTF8String: menu.name];
    
    NSMenu* nsmenu = ((AppDelegate*)NSApp.delegate).dock;
    NSArray* path = [name componentsSeparatedByString:@"/"];
    
    for (int i = 0; i < path.count - 1; ++i) {        
        nsmenu = Menu_GetOrSet(nsmenu, path[i]);
    }
    
    NSString* title = (NSString*)path.lastObject;
    Menu_SetMenuItem(nsmenu, menu, title);
}

void Menu_SetMenuItem(NSMenu* nsmenu , Menu__ menu, NSString* title) {
    if (menu.separator) {
        [nsmenu addItem: [NSMenuItem separatorItem]];
        return;
    }
    
    NSString* name = [NSString stringWithUTF8String: menu.name];
    NSString* shortcut = [NSString stringWithUTF8String: menu.shortcut];
    NSString* nativeAction = [NSString stringWithUTF8String: menu.nativeAction];
    
    NSMenuItem* item = [nsmenu itemWithTitle: title];
    
    if (item == nil) {
        MenuItem* mitem = [[MenuItem alloc] initWithTitle:title
                                                   action:nil
                                            keyEquivalent:@""];
        mitem.name = name;
        item = mitem;
        [nsmenu addItem: item];
    }
    
    item.indentationLevel = (NSInteger)menu.indent;
    
    if (shortcut.length != 0) {
        Menu_SetShortcut(item, shortcut);
    }
    
    if (nativeAction.length != 0) {
        item.action = NSSelectorFromString(nativeAction);
        return;
    }
    
    if (menu.enabled) {
        item.action = @selector(onMenuClick:);
    } else {
        item.action = nil;
    }
}

void Menu_SetShortcut(NSMenuItem* item, NSString* shortcut) {
    NSArray* keys = [shortcut componentsSeparatedByString:@"+"];
    
    item.keyEquivalentModifierMask = 0;
    
    for (NSString* k in keys) {
        if ([k isEqual: @"meta"]) {
            item.keyEquivalentModifierMask |= NSCommandKeyMask;
        } else if ([k isEqual: @"ctrl"]) {
            item.keyEquivalentModifierMask |= NSControlKeyMask;
        } else if ([k isEqual: @"alt"]) {
            item.keyEquivalentModifierMask |= NSAlternateKeyMask;
        } else if ([k isEqual: @"shift"]) {
            item.keyEquivalentModifierMask |= NSShiftKeyMask;
        } else if ([k isEqual: @"fn"]) {
            item.keyEquivalentModifierMask |= NSFunctionKeyMask;
        } else if ([k isEqual: @""]) {
            item.keyEquivalent = @"+";
        } else {
            item.keyEquivalent = k;
        }
    }
}

// ============================================================================
// Window
// ============================================================================

@implementation WindowController
- (WindowController*) initWithID:(NSString*)ID andConf:(WindowConfig__)conf {
    NSRect contentRect = NSMakeRect(conf.x, conf.y, conf.width, conf.height);
    
    NSUInteger styleMask = NSTitledWindowMask
        | NSFullSizeContentViewWindowMask
        | NSClosableWindowMask
        | NSMiniaturizableWindowMask
        | NSResizableWindowMask;
    
    if (conf.borderless) {
        styleMask = styleMask & ~NSTitledWindowMask;
    }
    
    if (conf.disableResize) {
        styleMask = styleMask & ~NSResizableWindowMask;
    }
    
    if (conf.disableClose) {
        styleMask = styleMask & ~NSClosableWindowMask;
    }
    
    if (conf.disableMinimize) {
        styleMask = styleMask & ~NSMiniaturizableWindowMask;
    }

    NSWindow* window = [[NSWindow alloc] initWithContentRect:contentRect
                                                   styleMask:styleMask
                                                     backing:NSBackingStoreBuffered
                                                       defer:NO];
    
    NSString* title = [NSString stringWithUTF8String: conf.title];
    
    if (title.length == 0) {
        window.titlebarAppearsTransparent = YES;
    } else {
        window.title = title;
    }
    
    WindowController* windowController = [[WindowController alloc] initWithWindow: window];
    windowController.ID = ID;
    windowController.windowFrameAutosaveName = ID;
    window.delegate = windowController;
    return windowController;
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
    onWindowMinimize((char*)[self.ID UTF8String]);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
    onWindowDeminimize((char*)[self.ID UTF8String]);
}

- (void)windowDidEnterFullScreen:(NSNotification *)notification {
    onWindowFullScreen((char*)[self.ID UTF8String]);
}

- (void)windowDidExitFullScreen:(NSNotification *)notification {
    onWindowExitFullScreen((char*)[self.ID UTF8String]);
}

- (void)windowDidMove:(NSNotification *)notification {
    NSWindow* window = (NSWindow*)notification.object;
    onWindowMove((char*)[self.ID UTF8String], window.frame.origin.x, window.frame.origin.y);
}

- (void)windowDidResize:(NSNotification *)notification {
    NSWindow* window = (NSWindow*)notification.object;
    onWindowResize((char*)[self.ID UTF8String], window.frame.size.width, window.frame.size.height);
}

- (void)windowDidBecomeKey:(NSNotification *)notification {
    onWindowFocus((char*)[self.ID UTF8String]);
}

- (void)windowDidResignKey:(NSNotification *)notification {
    onWindowBlur((char*)[self.ID UTF8String]);
}

- (BOOL)windowShouldClose:(id)sender {
    return onWindowClose((char*)[self.ID UTF8String]);
}

- (void)windowWillClose:(NSNotification *)notification {
    WindowController* windowController = (__bridge_transfer WindowController*)((__bridge void*)self);
}
@end

void* Window_Create(const char* ID, WindowConfig__ conf) {
    NSString* IDString = [NSString stringWithUTF8String: ID];
    WindowController* controller = [[WindowController alloc] initWithID:IDString
                                                                andConf:conf];

    return (__bridge_retained void*)controller;
}

void Window_Show(void* ptr) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    [windowController showWindow:windowController];
}

void Window_Move(void* ptr, CGFloat x, CGFloat y) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    CGPoint origin = windowController.window.frame.origin;
    origin.x = x;
    origin.y = y;
    [windowController.window setFrameOrigin:origin];
}

void Window_Center(void* ptr) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    [windowController.window center];
}

void Window_Resize(void* ptr, CGFloat width, CGFloat height) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    CGRect frame = windowController.window.frame;
    frame.size.width = width;
    frame.size.height = height;
    [windowController.window setFrame:frame
                              display:YES];
}

void Window_Close(void* ptr) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    [windowController.window performClose:windowController];
}