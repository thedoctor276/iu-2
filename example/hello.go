package main

import "github.com/maxence-charriere/iu"

type HelloComponent struct {
	iu.Div // embedding a div as a base

	Greeting *iu.Text  // to keep an easy access on the greeting
	Input    *iu.Input // to keep an easy access on the input
}

func (comp *HelloComponent) OnInputChange(v string) {
	comp.Input.Value = v

	comp.Greeting.Value = v
	comp.Greeting.Sync() // instruct the app to redraw only the Text component
}

func NewHelloComponent() *HelloComponent {
	// create the component
	var hello = &HelloComponent{
		Greeting: &iu.Text{Value: "World"},
		Input:    &iu.Input{},
	}

	// plug the event handler
	hello.Input.OnChange = hello.OnInputChange

	// design the Component
	hello.Div = iu.Div{
		Class: "hello",
		Content: []iu.Component{

			// the title
			&iu.H{
				Content: []iu.Component{
					&iu.Text{Value: "Hello,"},
					hello.Greeting,
				},
			},

			// the input
			hello.Input,
		},
	}

	return hello
}

func NewHelloPage() *iu.Page {
	return &iu.Page{
		Title: "Hello",
		Lang:  "en",
		CSS: []string{
			"hello.css",
		},

		Body: []iu.Component{
			NewHelloComponent(), // the hello component
		},
	}
}

func main() {
	var ctx = iu.EmptyContext{} // to be repaclaced by a Mac OSX window fo eg.
	var view = NewHelloPage()

	ctx.Navigate(view)
}
