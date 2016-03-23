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

type eventMessage struct {
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

func TryCallViewEvent(view View, name string, args ...interface{}) {
	tryCallEvent(view, name, args...)
}

func TryCallComponentEvent(component Component, name string, args ...interface{}) {
	tryCallEvent(component, name, args...)
}

func tryCallEvent(obj interface{}, name string, args ...interface{}) {
	var objValue = reflect.ValueOf(obj).Elem()
	var eventValue = objValue.FieldByName(name)
	var eventType = eventValue.Type()
	var argLen = len(args)
	var numIn int
	var in []reflect.Value

	if eventValue.Kind() != reflect.Func {
		iulog.Panicf("%v must be a func")
	}

	if eventValue.IsNil() {
		return
	}

	if numIn = eventType.NumIn(); numIn != argLen {
		iulog.Panicf("args and in have a different len: %v != %v", argLen, numIn)
	}

	in = make([]reflect.Value, numIn)

	for i := 0; i < numIn; i++ {
		var inType = eventType.In(i)
		var argValue = reflect.ValueOf(args[i])
		var argType = argValue.Type()

		if inType != argType {
			iulog.Panicf("inType: %v != argType: %v", inType, argType)
		}

		in[i] = argValue
	}

	eventValue.Call(in)
}

func tryCallComponentEventWithArg(component Component, name string, arg string) {
	var mouseEvent MouseEvent
	var wheelEvent WheelEvent
	var keyboardEvent KeyboardEvent
	var componentValue = reflect.ValueOf(component)
	var objValue = componentValue.Elem()
	var eventValue = objValue.FieldByName(name)
	var eventType reflect.Type
	var err error

	if !eventValue.IsValid() {
		eventValue = componentValue.MethodByName(name)
	}

	if eventValue.Kind() != reflect.Func {
		iulog.Panicf("%v must be a func")
	}

	if eventValue.IsNil() {
		return
	}

	eventType = eventValue.Type()

	if numIn := eventType.NumIn(); numIn == 0 {
		eventValue.Call(nil)
		return
	}

	switch argType := eventType.In(0); argType {
	case reflect.TypeOf(mouseEvent):
		if err = json.Unmarshal([]byte(arg), &mouseEvent); err != nil {
			iulog.Panic(err)
		}

		eventValue.Call([]reflect.Value{reflect.ValueOf(mouseEvent)})

	case reflect.TypeOf(wheelEvent):
		if err = json.Unmarshal([]byte(arg), &wheelEvent); err != nil {
			iulog.Panic(err)
		}

		eventValue.Call([]reflect.Value{reflect.ValueOf(wheelEvent)})

	case reflect.TypeOf(keyboardEvent):
		if err = json.Unmarshal([]byte(arg), &keyboardEvent); err != nil {
			iulog.Panic(err)
		}

		eventValue.Call([]reflect.Value{reflect.ValueOf(keyboardEvent)})

	default:
		eventValue.Call([]reflect.Value{reflect.ValueOf(arg)})
	}
}
