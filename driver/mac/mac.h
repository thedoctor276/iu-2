#ifndef mac_h
#define mac_h

#import <Cocoa/Cocoa.h>
#import <WebKit/Webkit.h>

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
@interface AppDelegate : NSObject <NSApplicationDelegate>
@property NSMenu* dock;

- (instancetype) init;
- (void) onMenuClick:(id)sender;
@end

void* App_Init();
void App_Run();
void App_Quit();
void App_SetBadge(const char* b);

// ============================================================================
// Menu
// ============================================================================

typedef struct Menu__ {
    const char* name;
    const char* shortcut;
    const char* handlerName;
    const char* compoID;
    unsigned int indent;
    BOOL disabled;
    BOOL separator;
    BOOL nativeHandler;
} Menu__;

@interface MenuItem : NSMenuItem
@property NSString* name;
@property NSString* compoID;
@end

NSMenu* Menu_GetOrSet(NSMenu* base, NSString* name);
void Menu_Set(Menu__ nsmenu);
void Menu_SetDock(Menu__ nsmenu);
void Menu_SetContext(NSMenu* ctx, Menu__ menu);
void Menu_SetMenuItem(NSMenu* nsmenu , Menu__ menu, NSString* title);
void Menu_SetShortcut(NSMenuItem* item, NSString* shortcut);

// ============================================================================
// WebView
// ============================================================================

@interface WKWebView (withDrawsBackground)
- (void)_setDrawsTransparentBackground:(BOOL)drawsTransparentBackground;
@end

// ============================================================================
// Window
// ============================================================================

typedef struct WindowConfig__ {
    CGFloat x;
    CGFloat y;
    CGFloat width;
    CGFloat height;
    const char* title;
    const char* backgroundColor;
    unsigned int backgroundType;
    BOOL borderless;
    BOOL disableResize;
    BOOL disableClose;
    BOOL disableMinimize;
} WindowConfig__;

@interface WindowController : NSWindowController <NSWindowDelegate, WKNavigationDelegate, WKUIDelegate, WKScriptMessageHandler>
@property NSString* ID;
@property (weak) WKWebView* webView;

- (instancetype) initWithID:(NSString*)ID andConf:(WindowConfig__)conf;
- (void) setupWebView;
- (void) setupCustomTitleBar;
@end

@interface TitleBar : NSView
@end

void* Window_Create(const char* ID, WindowConfig__ conf);
void Window_Show(void* ptr);
void Window_Move(void* ptr, CGFloat x, CGFloat y);
void Window_Center(void* ptr);
void Window_Resize(void* ptr, CGFloat width, CGFloat height);
void Window_Close(void* ptr);
void Window_Render(void* ptr, const char* HTML, const char* baseURL);
void Window_CallJavascript(void* ptr, const char* c);
void Window_ShowContextMenu(void* ptr, const Menu__* menus, int count);
void Window_Alert(void* ptr, const char* msg);


// ============================================================================
// Color
// ============================================================================

@interface CIColor(MBCategory) 
+ (CIColor*)colorWithHex:(UInt32)col;
+ (CIColor*)colorWithHexString:(NSString*)str;
@end

// ============================================================================
// Util
// ============================================================================

const char* ResourcesPath();

#endif /* mac_h */