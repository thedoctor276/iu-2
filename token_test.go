package iu

import "testing"

func TestComponentTokenFromString(t *testing.T) {
	expected := ComponentToken(42)

	if id := ComponentTokenFromString("42"); id != expected {
		t.Errorf("id should be %v: %v", expected, id)
	}
}

func TestComponentTokenFromInvalidString(t *testing.T) {
	defer func() { recover() }()

	ComponentTokenFromString("4dasf2234-=2")
	t.Error("should have panic")
}

func TestNextComponentToken(t *testing.T) {
	currentComponentID = 0
	defer func() { currentComponentID = 0 }()

	if id := nextComponentToken(); id != ComponentToken(1) {
		t.Errorf("id should be 1: %v", id)
	}
}
