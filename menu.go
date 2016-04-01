package iu

type Menu struct {
	Name         string
	Shortcut     string
	NativeAction string
	Indent       uint
	Enabled      bool
	Separator    bool
	Handler      func()
}
