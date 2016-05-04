package main

import (
	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
	"github.com/maxence-charriere/iu/driver/mac"
)

type Hello struct {
	Greeting string
	Input    string
}

func (h *Hello) Template() string {
	return `
<div class="content">
    <div class="hellobox">
        <h1>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>
        <input type="text" 
               autofocus 
               value="{{if .Greeting}}{{.Greeting}}{{end}}" 
               placeholder="What is your name?" 
               onchange="{{.RaiseEvent "OnChange" "value"}}"
               oncontextmenu="{{.RaiseEvent "OnContextMenu"}}">
    </div>
</div>
    `
}

func (h *Hello) ContextMenu() []iu.Menu {
	return []iu.Menu{
		iu.Menu{
			Name:     "Custom button",
			Shortcut: "meta+k",
		},
		iu.Menu{Separator: true},
		mac.CtxMenuCut,
		mac.CtxMenuCopy,
		mac.CtxMenuPaste,
	}
}

func (h *Hello) OnChange(name string) {
	h.Greeting = name
	iu.RenderComponent(h)
}

func (h *Hello) OnContextMenu() {
	iu.ShowContextMenu(h)
}

func (h *Hello) CustomCtx() {
	iulog.Warn("Custom context menu Callback")
}
