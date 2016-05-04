# iu
[![GoDoc](https://godoc.org/github.com/maxence-charriere/iu?status.svg)](https://godoc.org/github.com/maxence-charriere/iu) [![Go Report Card](https://goreportcard.com/badge/github.com/maxence-charriere/iu)](https://goreportcard.com/report/github.com/maxence-charriere/iu) [![Build Status](https://travis-ci.org/maxence-charriere/iu.svg?branch=master)](https://travis-ci.org/maxence-charriere/iu)

Package to create user interfaces with GO, HTML and CSS.

## Install
```
go get -u github.com/maxence-charriere/iu
```

## Available platforms
- OSX
![OK](https://upload.wikimedia.org/wikipedia/commons/thumb/8/80/Symbol_OK.svg/16px-Symbol_OK.svg.png) - package builder come soon
- IOS
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)
- Android
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)
- Windows
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)

## Getting started
[Documentation](https://godoc.org/github.com/maxence-charriere/iu)

![Hello](https://www.dropbox.com/s/kagdq53o2j7ttr0/Screen%20Shot%202016-03-28%20at%2018.11.51.png?raw=1)

### Create a [component](https://github.com/maxence-charriere/iu/blob/master/component.go)
```go
// hello.go
type Hello struct {
	Greeting    string
	Input       string
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

```

### Configure the app
```go
// main.go

func main() {
	mac.SetMenu(mac.MenuQuit)
	mac.SetMenu(mac.MenuCut)
	mac.SetMenu(mac.MenuCopy)
	mac.SetMenu(mac.MenuPaste)
	mac.SetMenu(mac.MenuSelectAll)
	mac.SetMenu(mac.MenuClose)

	mac.OnLaunch = onLaunch
	mac.OnReopen = onReopen

	mac.Run()
}

func onLaunch() {
	win := newMainWindow()
	win.Show()
}

func onReopen() {
	d, ok := iu.DriverByID("Main")
	if !ok {
		d = newMainWindow()
	}

	d.(*mac.Window).Show()
}

func newMainWindow() *mac.Window {
	hello := &Hello{}

	return mac.NewWindow(hello, iu.DriverConfig{
		ID:  "Main",
		CSS: []string{"hello.css"},
		Window: iu.WindowConfig{
			Width:      1240,
			Height:     720,
			Background: iu.WindowBackgroundDark,
		},
	})
}

```

### Stylize the result
```css
h1 {
    font-weight: 300;
}

input {
    border: 0pt;
    border-left: 2pt;
    border-color: lightgrey;
    background-color: transparent;
    border-style: solid;
    box-shadow: none;
    padding: 5pt;
    outline: none;
    
    font-size: 11pt;
    color: white;    
}

input:focus {
    border-color: dodgerblue;
}

.content {
    margin: 0pt;
    width: 100%;
    height: 100%;
    
    background-size: cover;
    background-repeat: no-repeat;
    background-position: 50% 50%;
    background-image: url("bg.jpg");
    
    color: white;
}

.hellobox {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}
```
