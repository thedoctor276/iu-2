# iu
Package to create user interfaces for iu apps.
## Install
```
go get -u github.com/maxence-charriere/iu
```
## Concept
**iu** is a tool to create apps using go programming language. It is composed of two layers:
* **UI** where the user interface is designed (this package).

* **App** which is the container for the UI (currently a Mac OSX app, will come later on IOs, Android and Windows)

## Getting started
The UI layer is built by encapsulation of go structs, each equivalent to an HTML tag.

### I. Build a hello component
```go
type HelloComponent struct {
	iu.Div // embedding a div as a base

	Greeting *iu.Text  // to keep an easy access on the greeting
	Input    *iu.Input // to keep an easy access on the input
}

func (comp *HelloComponent) OnInputChange(v string) {
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
```

### II. Build a page
```go
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
```

### III. Load it in the app
```go
func main() {
	var ctx = iu.EmptyContext{} // to be replaced by a Mac OSX window for eg.
	var view = NewHelloPage()

	ctx.Navigate(view)
}

```
