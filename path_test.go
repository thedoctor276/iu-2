package iu

import (
	"path/filepath"
	"testing"
)

func TestResourcePath(t *testing.T) {
	expected := "resources/test"

	if p := ResourcePath("test"); p != expected {
		t.Errorf("p should be %v: %v", expected, p)
	}
}

func TestSetResourcePath(t *testing.T) {
	f := func(elem ...string) string {
		elem = append([]string{"rss"}, elem...)
		return filepath.Join(elem...)
	}

	SetResourcePath(f)
	expected := "rss"

	if p := ResourcePath(); p != expected {
		t.Errorf("p should be %v: %v", expected, p)
	}
}
