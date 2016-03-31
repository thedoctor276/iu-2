package iu

var menuHandlers = map[string]func(){}

type MenuElement struct {
	Name         string
	Shortcut     string
	NativeAction string
	Indent       uint
	Enabled      bool
	Separator    bool
	Handler      func()
}

type menuHandlerMap map[string]func()

func RegisterMenuHandler(e MenuElement) {
	if e.Handler == nil {
		return
	}

	menuHandlers[e.Name] = e.Handler
}

func UnregisterMenuHandler(name string) {
	delete(menuHandlers, name)
}

func MenuHandler(name string) (h func(), ok bool) {
	h, ok = menuHandlers[name]
	return
}
