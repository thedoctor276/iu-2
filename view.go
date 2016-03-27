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
	return fmt.Sprintf("CallEventHandler(this.id, '%v', %v)", eventName, arg)
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

	if c.page == nil {
		iulog.Panicf(`component for %#v must be embedded in a page ~> use iu.NewPage(mainView View, config PageConfig) *Page`, v)
	}

	if c.page.Context() == nil {
		iulog.Panicf(`component.page for %#v must have a context ~> use [Context].Navigate(page *Page)`, v)
	}

	c.page.Context().InjectComponent(c)
}
