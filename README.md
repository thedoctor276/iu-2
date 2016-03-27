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
### I. Create a view
```go
type Hello struct {
	Greeting string
	Input    string
}

func (hello *Hello) Template() string {
	return `
<div id={{.ID}}>
    <h1>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>
    <input type="text" placeholder="What is your name?" onchange="{{.OnEvent "OnChange" "value"}}">
</div>
    `
}

func (hello *Hello) OnChange(name string) {
    hello.Greeting = name
    iu.SyncView(hello)
}

```

### II. Load it in the app
```go
func main() {
	ctx := New[osx|ios|android|windows]Context()

	page := iu.NewPage(&Hello{}, iu.PageConfig{
		Title: "Hello page",
		CSS:   []string{"hello.css"},
	})

	ctx.Navigate(page)
}
```
