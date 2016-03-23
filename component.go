package iu

import "sync"

var (
	lastComponentID uint64
	componentMutex  sync.Mutex
)

type Component interface {
	ID() string

	View() View

	setView(v View)

	Render() string

	Sync()
}

func nextComponentId() uint64 {
	componentMutex.Lock()
	defer componentMutex.Unlock()

	lastComponentID++
	return lastComponentID
}
