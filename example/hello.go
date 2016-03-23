package main

import "github.com/maxence-charriere/iu"

const (
	HelloTpl = `
<div id={{.ID}}>
    <h1>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnInputChanged" "event"}}">
</div>
    `
)

type HelloComponent struct {
	Greeting string
	Input    string
	*iu.Composer
}

func (hello *HelloComponent) OnInputChanged(value string) {
	hello.Greeting = value
	hello.Sync()
}

func (hello *HelloComponent) Render() string {
	return hello.Composer.Render(hello)
}

func (hello *HelloComponent) Sync() {
	hello.Dirty()
	hello.Render()
}

func NewHelloComponent() *HelloComponent {
	return &HelloComponent{
		Composer: iu.NewComposer(HelloTpl),
	}
}

func main() {
	ctx := &iu.EmptyContext{}

	hello := iu.NewPage(NewHelloComponent(), iu.PageConfig{
		Title: "Hello page",
		CSS:   []string{"hello.css"},
	})

	ctx.Navigate(hello)
}
