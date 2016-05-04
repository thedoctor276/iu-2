package iu

import "path/filepath"

var resourcesPath = func(elem ...string) string {
	elem = append([]string{"resources"}, elem...)
	return filepath.Join(elem...)
}

// ResourcesPath constructs a path relative to the "resources" directory.
func ResourcesPath(elem ...string) string {
	return resourcesPath(elem...)
}

// SetResourcesPath set the function which constructs a path relative to the "resources" directory.
// This call should be used only in a driver implementation.
func SetResourcesPath(f func(...string) string) {
	resourcesPath = f
}
