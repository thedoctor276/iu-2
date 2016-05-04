package iu

import (
	"encoding/json"
	"reflect"

	"github.com/maxence-charriere/iu-log"
)

const (
	// DeltaPixel indicates that the delta values are specified in pixels.
	DeltaPixel DeltaMode = iota

	// DeltaLine indicates that the delta values are specified in lines.
	DeltaLine

	// DeltaPage indicates that the delta values are specified in pages.
	DeltaPage
)

// EventMessage represents a message sent by a driver to call a component function.
// This structure should be only used in a driver implementation.
type EventMessage struct {
	ID   string
	Name string
	Arg  string
}

// MouseEvent represents events that occur due to the user interacting
// with a pointing device (such as a mouse).
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

// WheelEvent represents events fired when a wheel button of a
// pointing device (usually a mouse) is rotated.
type WheelEvent struct {
	DeltaX    float64
	DeltaY    float64
	DeltaZ    float64
	DeltaMode DeltaMode
}

// DeltaMode is an indication of the units of measurement for a delta value.
type DeltaMode uint64

// KeyboardEvent describes a user interaction with the keyboard.
type KeyboardEvent struct {
	CharCode rune
	KeyCode  KeyCode
	Location KeyLocation
	AltKey   bool
	CtrlKey  bool
	MetaKey  bool
	ShiftKey bool
}

// CallComponentEvent calls a component event handler.
// Should be only used in a driver implementation.
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
