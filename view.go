package iu

type View interface {
	Context() Context

	Component(id string) Component

	RegisterComponent(component Component)

	UnregisterComponent(component Component)

	Init(context Context)

	Render() string
}
