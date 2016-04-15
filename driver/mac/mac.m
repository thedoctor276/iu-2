#include "_cgo_export.h"
#include "mac.h"
#import <QuartzCore/QuartzCore.h>

// ============================================================================
// App
// ============================================================================

@implementation AppDelegate
- (instancetype) init {
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
    
    if (item.compoID.length == 0) {
        onMenuClick((char*)[item.name UTF8String]);
        return;
    }
    
    onContextMenuClick((char*)[item.name UTF8String], (char*)[item.compoID UTF8String]);
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
    [NSApp terminate:nil];
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

void Menu_SetContext(NSMenu* ctx, Menu__ menu) {
    NSString* name = [NSString stringWithUTF8String: menu.name];
    NSArray* path = [name componentsSeparatedByString:@"/"];
    
    for (int i = 0; i < path.count - 1; ++i) {
        ctx = Menu_GetOrSet(ctx, path[i]);
    }
    
    NSString* title = (NSString*)path.lastObject;
    Menu_SetMenuItem(ctx, menu, title);
}

void Menu_SetMenuItem(NSMenu* nsmenu , Menu__ menu, NSString* title) {
    if (menu.separator) {
        [nsmenu addItem: [NSMenuItem separatorItem]];
        return;
    }
    
    NSString* name = [NSString stringWithUTF8String: menu.name];
    NSString* shortcut = [NSString stringWithUTF8String: menu.shortcut];
    NSString* handlerName = [NSString stringWithUTF8String: menu.handlerName];
    
    NSMenuItem* item = [nsmenu itemWithTitle: title];
    
    if (item == nil) {
        MenuItem* mitem = [[MenuItem alloc] initWithTitle:title
                                                   action:nil
                                            keyEquivalent:@""];
        mitem.name = name;
        
        if (menu.compoID != NULL){
            NSString* compoID = [NSString stringWithUTF8String: menu.compoID];
            mitem.compoID = compoID;
        }
        
        item = mitem;
        [nsmenu addItem: item];
    }
    
    item.indentationLevel = (NSInteger)menu.indent;
    
    if (shortcut.length != 0) {
        Menu_SetShortcut(item, shortcut);
    }
    
    if (handlerName.length != 0) {
        item.action = NSSelectorFromString(handlerName);
        return;
    }
    
    if (menu.disabled) {
        item.action = nil;
    } else {
        item.action = @selector(onMenuClick:);
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
- (instancetype) initWithID:(NSString*)ID andConf:(WindowConfig__)conf {
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
    
    NSVisualEffectView* visualEffectView = nil;
    
    if (conf.background != 0) {
        visualEffectView = [[NSVisualEffectView alloc] initWithFrame: contentRect];
        
        switch (conf.background) {
            case 2:
                visualEffectView.material = NSVisualEffectMaterialMediumLight;
                break;
                
            case 3:
                visualEffectView.material = NSVisualEffectMaterialDark;
                break;
                
            case 4:
                visualEffectView.material = NSVisualEffectMaterialUltraDark;
                break;
                
            default:
                visualEffectView.material = NSVisualEffectMaterialLight;
                break;
        }
        
        visualEffectView.blendingMode = NSVisualEffectBlendingModeBehindWindow;
        visualEffectView.state = NSVisualEffectStateActive;
        window.contentView = visualEffectView;
    }
    
    self = [self initWithWindow: window];
    self.ID = ID;
    self.windowFrameAutosaveName = ID;
    window.delegate = self;
    
    [self setupWebView];
    
    NSString* title = [NSString stringWithUTF8String: conf.title];
    
    if (title.length == 0) {
        [self setupCustomTitleBar];
    } else {
        window.title = title;
    }
    
    return self;
}

- (void) setupWebView {
    WKWebViewConfiguration* configuration = [[WKWebViewConfiguration alloc] init];
    WKUserContentController* userContentController = [[WKUserContentController alloc] init];
    [userContentController addScriptMessageHandler:self name:@"onCallEventHandler"];
    configuration.userContentController = userContentController;
    
    WKWebView* webView = [[WKWebView alloc] initWithFrame: NSMakeRect(0, 0, 0, 0)
                                            configuration:configuration];
    [webView _setDrawsTransparentBackground:YES];
    webView.translatesAutoresizingMaskIntoConstraints = NO;
    [self.window.contentView addSubview: webView];
    
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat: @"|[webView]|"
                                                                                    options: 0
                                                                                    metrics: nil
                                                                                      views: NSDictionaryOfVariableBindings(webView)]];
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat: @"V:|[webView]|"
                                                                                    options: 0
                                                                                    metrics: nil
                                                                                      views: NSDictionaryOfVariableBindings(webView)]];
    self.webView = webView;
    webView.navigationDelegate = self;
    webView.UIDelegate = self;
}

- (void) setupCustomTitleBar {
    self.window.titlebarAppearsTransparent = YES;
    
    TitleBar* titleBar = [[TitleBar alloc] init];
    [self.window.contentView addSubview:titleBar];
    titleBar.translatesAutoresizingMaskIntoConstraints = false;
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat: @"|[titleBar]|"
                                                                                    options: 0
                                                                                    metrics: nil
                                                                                      views: NSDictionaryOfVariableBindings(titleBar)]];
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat: @"V:|[titleBar(==22)]"
                                                                                    options: 0
                                                                                    metrics: nil
                                                                                      views: NSDictionaryOfVariableBindings(titleBar)]];
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
    [self.webView.configuration.userContentController removeScriptMessageHandlerForName:@"onCallEventHandler"];
    
    WindowController* windowController = (__bridge_transfer WindowController*)((__bridge void*)self);
    windowController.window.delegate = nil;
}

- (void) webView:(WKWebView*)webView didFinishNavigation:(WKNavigation*)navigation {
    onWindowNavigate((char*)[self.ID UTF8String]);
}

- (void)webView:(WKWebView *)webView runJavaScriptAlertPanelWithMessage:(NSString*)message initiatedByFrame:(WKFrameInfo*)frame completionHandler:(void (^)(void))completionHandler {
    NSAlert* alert = [[NSAlert alloc] init];
    [alert setMessageText:message];
    [alert beginSheetModalForWindow:self.window
                  completionHandler:nil];
    
    completionHandler();
}

- (void) userContentController:(nonnull WKUserContentController *)userContentController didReceiveScriptMessage:(nonnull WKScriptMessage *)message {
    if ([message.name  isEqual: @"onCallEventHandler"]) {
        NSString* msg = (NSString*)message.body;
        onCallEventHandler((char*)[self.ID UTF8String], (char*)[msg UTF8String]);
    }
}
@end

@implementation TitleBar
- (void)mouseDragged:(nonnull NSEvent*)theEvent {
    [[self window] performWindowDragWithEvent:theEvent];
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

void Window_Navigate(void* ptr, const char* HTML, const char* baseURL) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    NSString* html = [NSString stringWithUTF8String: HTML];
    NSURL* base = [NSURL fileURLWithPath:[NSString stringWithUTF8String: baseURL]];
    
    [windowController.webView loadHTMLString:html
                                     baseURL:base];
}

void Window_InjectComponent(void* ptr, const char* ID, const char* component) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    NSString* call = [NSString stringWithFormat:@"InjectComponent(\"%s\", %s)", ID, component];
    [windowController.webView evaluateJavaScript: call
                               completionHandler: nil];
}

void Window_ShowContextMenu(void* ptr, const Menu__* menus, int count) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    NSMenu* ctxm = [[NSMenu alloc] initWithTitle:@"Context Menu"];
    
    for (int i = 0; i < count; ++i) {
        Menu_SetContext(ctxm, menus[i]);
    }
    
    NSPoint p = [windowController.window mouseLocationOutsideOfEventStream];

    [ctxm popUpMenuPositioningItem:ctxm.itemArray[0]
                        atLocation:p
                            inView:windowController.webView];
}

// ============================================================================
// Util
// ============================================================================

const char* ResourcePath() {
    NSBundle* mainBundle;
    mainBundle = [NSBundle mainBundle];
    return mainBundle.resourcePath.UTF8String;
}