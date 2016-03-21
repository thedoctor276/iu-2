package iu

import "github.com/maxence-charriere/iu-log"

type Context interface {
	Name() string

	CurrentView() View

	InjectComponent(component Component)
}

type EmptyContext struct {
	view View
}

func (ctx *EmptyContext) Name() string {
	return "EmptyContext"
}

func (ctx *EmptyContext) CurrentView() View {
	return ctx.view
}

func (ctx *EmptyContext) InjectComponent(component Component) {
	iulog.Printf(`Inject %v{ID: %v} in context`, component.Tag(), component.ID())
}
