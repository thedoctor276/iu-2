package iu

import "github.com/maxence-charriere/iu-log"

type Context interface {
	Name() string

	CurrentView() View

	InjectComponent(component Component)

	Navigate(view View)
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
	iulog.Printf(`Inject {ID: %v} in context`, component.ID())
}

func (ctx *EmptyContext) Navigate(view View) {
	view.Init(ctx)
	iulog.Printf(`Navigate -> %v`, view.Render())
}
