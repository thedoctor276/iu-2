# iu
Package to create user interfaces for iu apps.
## Install
```
go get -u github.com/maxence-charriere/iu
```
## Concept
**iu** is a tool to create apps using go programming language, HTML and CSS:
* **UI** where the user interface is designed (this package).

* **App** which is the container for the UI (currently a Mac OSX app, will come later on IOs, Android and Windows)

## Getting started
### I. Create a template
```go
const (
	HelloTpl = `
<div id={{.ID}}>
    <h1>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnInputChanged" "event"}}">
</div>
    `
)
```

### II. Create a component
```go
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

func NewHelloComponent() *HelloComponent {
	return &HelloComponent{
		Composer: iu.NewComposer(HelloTpl),
	}
}
```

### III. Load it in the app
```go
func main() {
	ctx := &iu.EmptyContext{}

	hello := iu.NewPage(NewHelloComponent(), iu.PageConfig{
		Title: "Hello page",
		CSS:   []string{"hello.css"},
	})

	ctx.Navigate(hello)
}
```
