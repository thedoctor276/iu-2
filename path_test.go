package iu

import (
	"path/filepath"
	"testing"
)

func TestResourcesPath(t *testing.T) {
	expected := "resources/test"

	if p := ResourcesPath("test"); p != expected {
		t.Errorf("p should be %v: %v", expected, p)
	}
}

func TestSetResourcesPath(t *testing.T) {
	f := func(elem ...string) string {
		elem = append([]string{"rss"}, elem...)
		return filepath.Join(elem...)
	}

	SetResourcesPath(f)
	expected := "rss"

	if p := ResourcesPath(); p != expected {
		t.Errorf("p should be %v: %v", expected, p)
	}
}
