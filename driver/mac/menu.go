package mac

import "github.com/maxence-charriere/iu"

var (
	menuHandlers = map[string]func(){}

	MenuAbout = iu.Menu{
		Name:         "About",
		NativeAction: "orderFrontStandardAboutPanel:",
	}

	MenuPreferences = iu.Menu{
		Name:     "Preferences",
		Shortcut: "meta+,",
		Enabled:  false,
	}

	MenuHide = iu.Menu{
		Name:         "Hide",
		Shortcut:     "meta+h",
		NativeAction: "hide:",
	}

	MenuShowAll = iu.Menu{
		Name:         "Show All",
		NativeAction: "unhideAllApplications:",
	}

	MenuQuit = iu.Menu{
		Name:         "Quit",
		Shortcut:     "meta+q",
		NativeAction: "terminate:",
	}

	MenuUndo = iu.Menu{
		Name:         "Edit/Undo",
		Shortcut:     "meta+z",
		NativeAction: "undo:",
	}

	MenuRedo = iu.Menu{
		Name:         "Edit/Redo",
		Shortcut:     "shift+meta+z",
		NativeAction: "redo:",
	}

	MenuCut = iu.Menu{
		Name:         "Edit/Cut",
		Shortcut:     "meta+x",
		NativeAction: "cut:",
	}

	MenuCopy = iu.Menu{
		Name:         "Edit/Copy",
		Shortcut:     "meta+c",
		NativeAction: "copy:",
	}

	MenuPaste = iu.Menu{
		Name:         "Edit/Paste",
		Shortcut:     "meta+v",
		NativeAction: "paste:",
	}

	MenuSelectAll = iu.Menu{
		Name:         "Edit/Select All",
		Shortcut:     "meta+a",
		NativeAction: "selectAll:",
	}

	MenuClose = iu.Menu{
		Name:         "Window/Close",
		Shortcut:     "meta+w",
		NativeAction: "performClose:",
	}

	MenuMinimize = iu.Menu{
		Name:         "Window/Minimize",
		Shortcut:     "meta+m",
		NativeAction: "performMiniaturize:",
	}

	MenuBringAllToFront = iu.Menu{
		Name:         "Window/Bring All to Front",
		NativeAction: "arrangeInFront:",
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
