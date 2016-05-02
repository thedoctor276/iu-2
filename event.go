package iu

import (
	"encoding/json"
	"reflect"

	"github.com/maxence-charriere/iu-log"
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

func CallComponentEvent(c Component, name string, arg string) {
	v := reflect.ValueOf(c)
	mv := v.MethodByName(name)

	if !mv.IsValid() {
		v := reflect.Indirect(v)
		mv = v.FieldByName(name)
	}

	if !mv.IsValid() {
		iulog.Panicf("%#v doesn't have any method or field named %v", c, name)
	}

	if mv.Kind() != reflect.Func {
		iulog.Panicf("field %v is not a func", name)
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

	if err := json.Unmarshal([]byte(arg), argi); err != nil {
		iulog.Panic(err)
	}

	mv.Call([]reflect.Value{argv.Elem()})
}
