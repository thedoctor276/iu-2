package iu

const (
	// WindowBackgroundSolid represents a solid background.
	WindowBackgroundSolid WindowBackgroundType = iota

	// WindowBackgroundLight represents a native light background.
	// Can have a significant impact on scrolling performances.
	WindowBackgroundLight

	// WindowBackgroundUltraLight represents a native even lighter background.
	// Can have a significant impact on scrolling performances.
	WindowBackgroundUltraLight

	// WindowBackgroundDark represents a native dark background.
	// Can have a significant impact on scrolling performances.
	WindowBackgroundDark

	// WindowBackgroundUltraDark represents a even darker background.
	// Can have a significant impact on scrolling performances.
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
	BackgroundColor string
	BackgroundType  WindowBackgroundType
	Borderless      bool
	DisableResize   bool
	DisableClose    bool
	DisableMinimize bool
}

// WindowBackgroundType represents the driver background type.
type WindowBackgroundType uint
