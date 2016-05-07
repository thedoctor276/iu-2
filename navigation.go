package iu

import "fmt"

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

type navigation struct {
	Current Component
	history []Component
	index   int
}

func (n *navigation) Template() string {
	return `
<div id="iu-nav">
{{.Current.Render}}
</div>
`
}

func (n *navigation) OnDismount() {
	n.Current = nil
	n.history = cleanhistory(n.history, 0)
}

func (n *navigation) CurrentComponent() Component {
	return n.history[n.index]
}

func (n *navigation) Go(c Component) {
	n.history = cleanhistory(n.history, n.index+1)
	n.history = append(n.history, c)
	n.index++
	n.Current = c

	d := DriverByComponent(n)
	MountComponent(c, d)
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

	n.index--
	n.Current = n.CurrentComponent()
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

	n.index++
	n.Current = n.CurrentComponent()
	RenderComponent(n)
	return
}

func newNavigation(c Component) *navigation {
	return &navigation{
		Current: c,
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
