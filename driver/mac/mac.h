#ifndef mac_h
#define mac_h

#import <Cocoa/Cocoa.h>

// ============================================================================
// App
// ============================================================================
@interface AppDelegate : NSObject <NSApplicationDelegate>
@property NSMenu* dock;

- (AppDelegate*) init;
- (void) onMenuClick:(id)sender;
@end

void* App_Init();
void App_Run();
void App_Quit();

// ============================================================================
// Menu
// ============================================================================

typedef struct Menu__ {
    const char* name;
    const char* shortcut;
    const char* nativeAction;
    unsigned int indent;
    BOOL enabled;
    BOOL separator;
} Menu__;

@interface MenuItem : NSMenuItem
@property NSString* name;
@end

NSMenu* Menu_GetOrSet(NSMenu* base, NSString* name);
void Menu_Set(Menu__ nsmenu);
void Menu_SetDock(Menu__ nsmenu);
void Menu_SetMenuItem(NSMenu* nsmenu , Menu__ menu, NSString* title);
void Menu_SetShortcut(NSMenuItem* item, NSString* shortcut);

#endif /* mac_h */