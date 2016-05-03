package iu

import "path/filepath"

var resourcePath = func(elem ...string) string {
	elem = append([]string{"resources"}, elem...)
	return filepath.Join(elem...)
}

// ResourcePath constructs a path relative to the "resources" directory.
func ResourcePath(elem ...string) string {
	return resourcePath(elem...)
}

// SetResourcePath set the function which constructs a path relative to the "resources" directory.
// This call should be used only in a driver implementation.
func SetResourcePath(f func(...string) string) {
	resourcePath = f
}
