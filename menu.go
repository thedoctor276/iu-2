package iu

import (
	"reflect"

	"github.com/maxence-charriere/iu-log"
)

type Menu struct {
	Name          string
	Shortcut      string
	HandlerName   string
	Indent        uint
	Disabled      bool
	Separator     bool
	NativeHandler bool
	Handler       func()
}

func CallContextMenuHandler(view View, name string) {
	v := reflect.Indirect(reflect.ValueOf(view))
	ctxmv := v.FieldByName("ContextMenu")

	if !ctxmv.IsValid() {
		iulog.Panicf("view %v doesn't have a context menu", v.Type())
	}

	if mt := reflect.TypeOf((*[]Menu)(nil)).Elem(); ctxmv.Type() != mt {
		iulog.Panicf("ContextMenu in view %v is not a []iu.Menu", v.Type())
	}

	ctxm := ctxmv.Interface().([]Menu)

	for _, m := range ctxm {
		if m.Name != name {
			continue
		}

		if m.Handler == nil {
			iulog.Warnf(`menu named "%v" in %v doesn't have a handler`, name, v.Type())
			return
		}

		m.Handler()
		return
	}
}
