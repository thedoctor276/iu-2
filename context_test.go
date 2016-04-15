package iu

import "github.com/maxence-charriere/iu-log"

type EmptyContext struct {
	page *Page
}

func (ctx *EmptyContext) Name() string {
	return "EmptyContext"
}

func (ctx *EmptyContext) CurrentPage() *Page {
	return ctx.page
}

func (ctx *EmptyContext) InjectComponent(component *Component) {
	iulog.Printf(`Inject {ID: %v} in context`, component.ID())
}

func (ctx *EmptyContext) Navigate(page *Page) {
	page.Context = ctx
	iulog.Printf(`Navigate -> %v`, page.Render())
}

func (ctx *EmptyContext) ShowContextMenu(menus []Menu, compoID string) {
	iulog.Printf(`ShowContextMenu -> %v`, menus)
}
