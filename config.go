package iu

const (
	// WindowBackgroundSolid represents a solid background.
	WindowBackgroundSolid WindowBackground = iota

	// WindowBackgroundLight represents a native light background.
	WindowBackgroundLight

	// WindowBackgroundUltraLight represents a native even lighter background.
	WindowBackgroundUltraLight

	// WindowBackgroundDark represents a native dark background.
	WindowBackgroundDark

	// WindowBackgroundUltraDark represents a even darker background.
	WindowBackgroundUltraDark
)

// DriverConfig is the configuration required by a driver.
type DriverConfig struct {
	ID     DriverToken
	Lang   string
	CSS    []string
	JS     []string
	OnLoad func()
	Window WindowConfig
}

// WindowConfig is the configuration that only apply on window type drivers.
type WindowConfig struct {
	X               float64
	Y               float64
	Width           float64
	Height          float64
	Title           string
	Background      WindowBackground
	Borderless      bool
	DisableResize   bool
	DisableClose    bool
	DisableMinimize bool
}

// WindowBackground set the driver background type.
type WindowBackground uint
