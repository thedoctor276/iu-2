package iu

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	DeltaPixel DeltaMode = iota
	DeltaLine
	DeltaPage
)

type EventMessage struct {
	ID   string
	Name string
	Arg  string
}

type MouseEvent struct {
	ClientX  float64
	ClientY  float64
	PageX    float64
	PageY    float64
	ScreenX  float64
	ScreenY  float64
	Button   int
	Detail   int
	AltKey   bool
	CtrlKey  bool
	MetaKey  bool
	ShiftKey bool
}

type WheelEvent struct {
	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode DeltaMode
}

type DeltaMode uint64

type KeyboardEvent struct {
	CharCode rune
	KeyCode  KeyCode
	Location KeyLocation
	AltKey   bool
	CtrlKey  bool
	MetaKey  bool
	ShiftKey bool
}

func CallViewEvent(view View, name string, arg string) (err error) {
	var mv reflect.Value

	v := reflect.ValueOf(view)

	if mv = v.MethodByName(name); !mv.IsValid() {
		v = reflect.Indirect(v)
		mv = v.FieldByName(name)
	}

	if !mv.IsValid() {
		err = fmt.Errorf("%#v doesn't have any method or field named", view, name)
		return
	}

	if mv.Kind() != reflect.Func {
		err = fmt.Errorf("field %v is not a func", name)
		return
	}

	if mv.IsNil() {
		return
	}

	mt := mv.Type()

	if mt.NumIn() == 0 {
		mv.Call(nil)
		return
	}

	argt := mt.In(0)
	argv := reflect.New(argt)
	argi := argv.Interface()

	if err = json.Unmarshal([]byte(arg), argi); err != nil {
		return
	}

	mv.Call([]reflect.Value{argv.Elem()})
	return
}
