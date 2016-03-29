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
![Hello](https://www.dropbox.com/s/kagdq53o2j7ttr0/Screen%20Shot%202016-03-28%20at%2018.11.51.png?raw=1)

### I. Create a view
```go
type Hello struct {
	Greeting string
	Input    string
}

func (hello *Hello) Template() string {
	return `
<div class="content">
    <div id="{{.ID}}" class="hellobox">
        <h1>Hello, <span>{{if .Greeting}}{{.Greeting}}{{else}}World{{end}}</span></h1>
        <input type="text" 
               autofocus 
               value="{{if .Greeting}}{{.Greeting}}{{end}}" 
               placeholder="What is your name?" 
               onchange="{{.OnEvent "OnChange" "value"}}">
    </div>
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
var ctx *iuosx.Window

func main() {
	iuosx.OnLaunch = OnLaunch
	iuosx.OnReopen = OnReopen

	iuosx.Run()
}

func OnLaunch() {
	ctx = iuosx.NewWindow("hello", iuosx.WindowConfig{
		Width:  1280,
		Height: 720,
	}) // IOS, Android and Windows will come later

	p := iu.NewPage(&Hello{}, iu.PageConfig{
		Title: "Hello page",
		CSS:   []string{"hello.css"},
	})

	ctx.Navigate(p)
}

func OnReopen() {
	ctx.Show()
}
```
