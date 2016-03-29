package iu

import (
	"encoding/json"
	"testing"
)

func TestCallViewEvent(t *testing.T) {
	v := &World{
		OnNothing: func() {
			t.Log("OnNothing")
		},
	}

	if err := CallViewEvent(v, "OnNothing", ""); err != nil {
		t.Error(err)
	}
}

func TestCallViewEventNotSet(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnNothing", ""); err != nil {
		t.Error(err)
	}
}

func TestCallViewNonexistentEvent(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnDrinkingBeer", ""); err == nil {
		t.Error("should error")
	}
}

func TestCallViewNoEvent(t *testing.T) {
	v := Hello{}

	if err := CallViewEvent(v, "Input", ""); err == nil {
		t.Error("should error")
	}
}

func TestCallViewMouseEvent(t *testing.T) {
	var kes []byte
	var err error

	ke := MouseEvent{
		MetaKey: true,
		ClientX: 42,
	}

	if kes, err = json.Marshal(ke); err != nil {
		t.Fatal(err)
	}

	v := &World{}

	if err = CallViewEvent(v, "OnClick", string(kes)); err != nil {
		t.Error(err)
	}
}

func TestCallViewWheelEvent(t *testing.T) {
	var kes []byte
	var err error

	ke := WheelEvent{
		DeltaX: 42,
	}

	if kes, err = json.Marshal(ke); err != nil {
		t.Fatal(err)
	}

	v := &World{}

	if err = CallViewEvent(v, "OnWheel", string(kes)); err != nil {
		t.Error(err)
	}
}

func TestCallViewKeyboardEvent(t *testing.T) {
	var kes []byte
	var err error

	ke := KeyboardEvent{
		CtrlKey: true,
		KeyCode: Key4,
	}

	if kes, err = json.Marshal(ke); err != nil {
		t.Fatal(err)
	}

	v := &World{}

	if err = CallViewEvent(v, "OnChange", string(kes)); err != nil {
		t.Error(err)
	}
}

func TestCallViewEventWithBool(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnChecked", "true"); err != nil {
		t.Error(err)
	}
}

func TestCallViewEventWithString(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnString", `"J'aime les filles"`); err != nil {
		t.Error(err)
	}
}

func TestCallViewEventWithNumber(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnNumber", "42.42"); err != nil {
		t.Error(err)
	}
}

func TestCallViewEventBadFormat(t *testing.T) {
	v := &World{}

	if err := CallViewEvent(v, "OnChange", "[stupid no json]"); err == nil {
		t.Error("should error")
	}
}
