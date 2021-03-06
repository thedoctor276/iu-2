package iu

import "fmt"

// NavigateHandler is the representation of a component which can perform
// an action when navigated on.
type NavigateHandler interface {
	OnNavigate()
}

// LeaveHandler is the representation of a component which can perform
// an action when leaving.
type LeaveHandler interface {
	OnLeave()
}

// Navigation is a representation of a container which handles
// navigation between components.
type Navigation interface {
	Component

	CurrentComponent() Component

	Go(c Component)

	CanBack() bool

	Back() error

	CanNext() bool

	Next() error
}

func (n *navigation) OnDismount() {
	n.current = nil
	n.history = cleanhistory(n.history, 0)
}

type navigation struct {
	current Component
	history []Component
	index   int
}

func (n *navigation) Template() string {
	return `
<div id="iu-nav">
{{.CurrentComponent.Render}}
</div>
`
}

func (n *navigation) CurrentComponent() Component {
	if len(n.history) == 0 {
		return nil
	}

	return n.history[n.index]
}

func (n *navigation) Go(c Component) {
	if lh, ok := n.current.(LeaveHandler); ok {
		lh.OnLeave()
	}

	n.history = cleanhistory(n.history, n.index+1)
	n.history = append(n.history, c)
	n.index++
	n.current = c

	d := DriverByComponent(n)
	MountComponent(c, d)

	if nh, ok := c.(NavigateHandler); ok {
		nh.OnNavigate()
	}

	RenderComponent(n)
}

func (n *navigation) CanBack() bool {
	return n.index > 0
}

func (n *navigation) Back() (err error) {
	if !n.CanBack() {
		err = fmt.Errorf("no entry before %v: %v", n.CurrentComponent(), n.history)
		return
	}

	if lh, ok := n.current.(LeaveHandler); ok {
		lh.OnLeave()
	}

	n.index--
	n.current = n.CurrentComponent()

	if nh, ok := n.current.(NavigateHandler); ok {
		nh.OnNavigate()
	}

	RenderComponent(n)
	return
}

func (n *navigation) CanNext() bool {
	return n.index < len(n.history)-1
}

func (n *navigation) Next() (err error) {
	if !n.CanNext() {
		err = fmt.Errorf("no entry after %v: %v", n.CurrentComponent(), n.history)
		return
	}

	if lh, ok := n.current.(LeaveHandler); ok {
		lh.OnLeave()
	}

	n.index++
	n.current = n.CurrentComponent()

	if nh, ok := n.current.(NavigateHandler); ok {
		nh.OnNavigate()
	}

	RenderComponent(n)
	return
}

func newNavigation(c Component) *navigation {
	return &navigation{
		current: c,
		history: []Component{c},
	}
}

func cleanhistory(history []Component, idx int) []Component {
	for i := idx; i < len(history); i++ {
		c := history[i]
		DismountComponent(c)
	}

	if idx < len(history) {
		return history[:idx]
	}
	return history
}
