package iu

import "fmt"

var (
	// SetBadgeHandler is the function pointer to be used when SetBadge is called.
	// Should be used only in driver implementations.
	SetBadgeHandler func(string)
)

func init() {
	SetBadgeHandler = func(badge string) {
		fmt.Printf(`handling SetBadge("%v")\n`, badge)
	}
}

// SetBadge set the app badge.
func SetBadge(badge string) {
	SetBadgeHandler(badge)
}
