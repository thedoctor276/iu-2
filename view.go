package iu

import (
	"fmt"
	"reflect"

	"github.com/maxence-charriere/iu-log"
)

type View interface {
	Template() string
}

type viewMap map[string]interface{}

func (m viewMap) OnEvent(eventName string, arg string) string {
	if len(arg) == 0 {
		arg = "event"
	}

	return fmt.Sprintf("CallEventHandler('%v', '%v', %v)", m["ID"], eventName, arg)
}

func newViewMap(ID string) viewMap {
	return viewMap{
		"ID": ID,
	}
}

func ForRangeViews(top View, action func(view View) error) {
	if err := action(top); err != nil {
		iulog.Error(err)
		return
	}

	v := reflect.Indirect(reflect.ValueOf(top))
	t := v.Type()
	vInterface := reflect.TypeOf((*View)(nil)).Elem()

	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i).Type

		if ft.Implements(vInterface) {
			view := v.Field(i).Interface().(View)
			ForRangeViews(view, action)
		}
	}
}

func SyncView(v View) {
	c := compoM.Component(v)
	c.MustBeUsable()
	c.page.Context.InjectComponent(c)
}

func ShowContextMenu(v View) {
	val := reflect.Indirect(reflect.ValueOf(v))
	cmVal := val.FieldByName("ContextMenu")
	cmType := cmVal.Type()

	if cmVal.Kind() != reflect.Slice {
		iulog.Panicf("ContextMenu field in %#v is not a slice", v)
	}

	if cmType.Elem() != reflect.TypeOf((*Menu)(nil)).Elem() {
		iulog.Panicf("ContextMenu field in %#v is not a slice of iu.Menu", v)
	}

	m := cmVal.Interface().([]Menu)

	c := compoM.Component(v)
	c.MustBeUsable()
	c.page.Context.ShowContextMenu(m, c.ID())
}
