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
    [NSApp activateIgnoringOtherApps : YES];
    [NSApp run];
}

void App_Quit() {
    defer(
        [NSApp terminate:nil];
    );
}

void App_SetBadge(const char* b) {
    NSString* badge = [NSString stringWithUTF8String: b];
    
    defer(
        NSDockTile* dock = [NSApp dockTile];
        [dock setBadgeLabel:badge];
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
    
    if (menu.nativeHandler) {
        item.action = NSSelectorFromString(handlerName);
        return;
    }
    
    if (menu.disabled) {
        item.action = nil;
        return;
    } 
    
    item.action = @selector(onMenuClick:);
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
    
    if (conf.backgroundType == 0) {
        CIColor* c = [CIColor colorWithHexString:@"#f5f5f5"];        
        NSString* bgcolor = [NSString stringWithUTF8String:conf.backgroundColor];
        
        if (bgcolor.length != 0) {
            c = [CIColor colorWithHexString:bgcolor]; 
        }
        
        window.backgroundColor = [NSColor colorWithCIColor:c];
        
    } else {
        visualEffectView = [[NSVisualEffectView alloc] initWithFrame: contentRect];
        
        switch (conf.backgroundType) {
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
    
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat:@"|[webView]|"
                                                                                    options:0
                                                                                    metrics:nil
                                                                                      views:NSDictionaryOfVariableBindings(webView)]];
    [self.window.contentView addConstraints:[NSLayoutConstraint constraintsWithVisualFormat:@"V:|[webView]|"
                                                                                    options:0
                                                                                    metrics:nil
                                                                                      views:NSDictionaryOfVariableBindings(webView)]];
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
    
    self.window.delegate = nil;
    self.webView.navigationDelegate = nil;
    self.webView.UIDelegate = nil;
    
    WindowController* windowController = (__bridge_transfer WindowController*)((__bridge void*)self);
}

- (void)webView:(WKWebView *)webView didFinishNavigation:(WKNavigation *)navigation {
    onWindowLoad((char*)[self.ID UTF8String]);
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction
                                                     decisionHandler:(void (^)(WKNavigationActionPolicy))decisionHandler {
    if (navigationAction.navigationType == WKNavigationTypeReload || navigationAction.navigationType == WKNavigationTypeOther) {
        decisionHandler(WKNavigationActionPolicyAllow);
        return;
    }
    
    NSURL* url = navigationAction.request.URL;
    [[NSWorkspace sharedWorkspace] openURL:url];
    decisionHandler(WKNavigationActionPolicyCancel);
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
    defer(
        WindowController* windowController = (__bridge WindowController*)ptr;
        [windowController showWindow:windowController];
    );
}

void Window_Move(void* ptr, CGFloat x, CGFloat y) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    defer(
        CGPoint origin = windowController.window.frame.origin;
        origin.x = x;
        origin.y = y;
        [windowController.window setFrameOrigin:origin];
    );
}

void Window_Center(void* ptr) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    defer(
        [windowController.window center];
    );
}

void Window_Resize(void* ptr, CGFloat width, CGFloat height) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    defer(
        CGRect frame = windowController.window.frame;
        frame.size.width = width;
        frame.size.height = height;
        [windowController.window setFrame:frame
                                  display:YES];
    );
}

void Window_Close(void* ptr) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    
    defer(
        [windowController.window performClose:windowController];
    );
}

void Window_Render(void* ptr, const char* HTML, const char* baseURL) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    NSData* html = [NSData dataWithBytes:HTML length:strlen(HTML)];
    NSURL* base = [NSURL fileURLWithPath:[NSString stringWithUTF8String: baseURL]];
    
    defer(
        [windowController.webView loadData:html
                                  MIMEType:@"text/html"
                     characterEncodingName:@"UTF-8"
                                   baseURL:base];
    );
}

void Window_CallJavascript(void* ptr, const char* c) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    NSString* call = [NSString stringWithUTF8String:c];
    
    defer(
        [windowController.webView evaluateJavaScript: call
                                   completionHandler: nil];
    );
}

void Window_ShowContextMenu(void* ptr, const Menu__* menus, int count) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    NSMenu* ctxm = [[NSMenu alloc] initWithTitle:@"Context Menu"];
    
    for (int i = 0; i < count; ++i) {
        Menu_SetContext(ctxm, menus[i]);
    }
    
    defer(
        NSPoint p = [windowController.window mouseLocationOutsideOfEventStream];

        [ctxm popUpMenuPositioningItem:ctxm.itemArray[0]
                            atLocation:p
                                inView:windowController.webView];
    );
}

void Window_Alert(void* ptr, const char* msg) {
    WindowController* windowController = (__bridge WindowController*)ptr;
    NSString* message = [NSString stringWithUTF8String:msg];
    
    defer(
        NSAlert* alert = [[NSAlert alloc] init];
        [alert setMessageText:message];
        [alert beginSheetModalForWindow:windowController.window
                      completionHandler:nil];
    );
}

// ============================================================================
// Color
// ============================================================================

@implementation CIColor(MBCategory)
+ (CIColor*)colorWithHexString:(NSString*)str {
    const char *cstr = [str cStringUsingEncoding:NSASCIIStringEncoding];
    long x = strtol(cstr+1, NULL, 16);
    return [CIColor colorWithHex:x];
}

+ (CIColor*)colorWithHex:(UInt32)col {
    unsigned char b = col & 0xFF;;
    unsigned char g = (col >> 8) & 0xFF;
    unsigned char r = (col >> 16) & 0xFF;

    return [CIColor colorWithRed:(float)r/255.0f 
                           green:(float)g/255.0f 
                            blue:(float)b/255.0f 
                           alpha:1];
}
@end

// ============================================================================
// Util
// ============================================================================

const char* ResourcesPath() {
    NSBundle* mainBundle;
    mainBundle = [NSBundle mainBundle];
    return mainBundle.resourcePath.UTF8String;
}