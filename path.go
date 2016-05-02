package iu

var resourcePath func(...string) string

// ResourcePath constructs a path relative to the "resources" directory.
func ResourcePath(elem ...string) string {
    resourcePath(...elem)
}

// SetResourcePath set the function which constructs a path relative to the "resources" directory.
// This call should be used only in a driver implementation.
func SetResourcePath(f func(...string) string){
    resourcePath = f
}
