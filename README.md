# iu
Package to create user interfaces with GO, HTML and CSS.

## Install
```
go get -u github.com/maxence-charriere/iu
```

## Available platforms
- OSX
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png) - very soon
- IOS
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)
- Android
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)
- Windows
![NO](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/No_icon_red.svg/16px-No_icon_red.svg.png)

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
