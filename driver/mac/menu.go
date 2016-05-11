package mac

import "github.com/maxence-charriere/iu"

var menuHandlers = map[string]func(){}

// Default mac app menus.
var (
	MenuAbout = iu.Menu{
		Name:          "About",
		HandlerName:   "orderFrontStandardAboutPanel:",
		NativeHandler: true,
	}

	MenuPreferences = iu.Menu{
		Name:     "Preferences",
		Shortcut: "meta+,",
		Disabled: true,
	}

	MenuHide = iu.Menu{
		Name:          "Hide",
		Shortcut:      "meta+h",
		HandlerName:   "hide:",
		NativeHandler: true,
	}

	MenuShowAll = iu.Menu{
		Name:          "Show All",
		HandlerName:   "unhideAllApplications:",
		NativeHandler: true,
	}

	MenuQuit = iu.Menu{
		Name:          "Quit",
		Shortcut:      "meta+q",
		HandlerName:   "terminate:",
		NativeHandler: true,
	}

	MenuUndo = iu.Menu{
		Name:          "Edit/Undo",
		Shortcut:      "meta+z",
		HandlerName:   "undo:",
		NativeHandler: true,
	}

	MenuRedo = iu.Menu{
		Name:          "Edit/Redo",
		Shortcut:      "shift+meta+z",
		HandlerName:   "redo:",
		NativeHandler: true,
	}

	MenuCut = iu.Menu{
		Name:          "Edit/Cut",
		Shortcut:      "meta+x",
		HandlerName:   "cut:",
		NativeHandler: true,
	}

	MenuCopy = iu.Menu{
		Name:          "Edit/Copy",
		Shortcut:      "meta+c",
		HandlerName:   "copy:",
		NativeHandler: true,
	}

	MenuPaste = iu.Menu{
		Name:          "Edit/Paste",
		Shortcut:      "meta+v",
		HandlerName:   "paste:",
		NativeHandler: true,
	}

	MenuSelectAll = iu.Menu{
		Name:          "Edit/Select All",
		Shortcut:      "meta+a",
		HandlerName:   "selectAll:",
		NativeHandler: true,
	}

	MenuClose = iu.Menu{
		Name:          "Window/Close",
		Shortcut:      "meta+w",
		HandlerName:   "performClose:",
		NativeHandler: true,
	}

	MenuMinimize = iu.Menu{
		Name:          "Window/Minimize",
		Shortcut:      "meta+m",
		HandlerName:   "performMiniaturize:",
		NativeHandler: true,
	}

	MenuBringAllToFront = iu.Menu{
		Name:          "Window/Bring All to Front",
		HandlerName:   "arrangeInFront:",
		NativeHandler: true,
	}

	CtxMenuCut = iu.Menu{
		Name:          "Cut",
		Shortcut:      "meta+x",
		HandlerName:   "cut:",
		NativeHandler: true,
	}

	CtxMenuCopy = iu.Menu{
		Name:          "Copy",
		Shortcut:      "meta+c",
		HandlerName:   "copy:",
		NativeHandler: true,
	}

	CtxMenuPaste = iu.Menu{
		Name:          "Paste",
		Shortcut:      "meta+v",
		HandlerName:   "paste:",
		NativeHandler: true,
	}
)

type menuHandlerMap map[string]func()

func registerMenuHandler(m iu.Menu) {
	if m.Handler == nil {
		return
	}

	menuHandlers[m.Name] = m.Handler
}

func unregisterMenuHandler(name string) {
	delete(menuHandlers, name)
}

func menuHandler(name string) (h func(), ok bool) {
	h, ok = menuHandlers[name]
	return
}
