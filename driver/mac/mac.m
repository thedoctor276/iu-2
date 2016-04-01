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
    NSLog(@"dock: %@", nsmenu); 
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

