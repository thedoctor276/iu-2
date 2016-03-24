package iu

import "testing"

func TestComposerMapOnEvent(t *testing.T) {
	expected := "CallEventHandler(this.id, 'OnClick', event)"
	m := composerMap{}

	if js := m.OnEvent("OnClick", "event"); js != expected {
		t.Errorf("js should be %v: %v", expected, js)
	}
}
