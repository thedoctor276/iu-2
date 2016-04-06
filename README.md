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

### Create a view
```go
// hello.go
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

### Configure the app
```go
// main.go

func main() {
	mac.SetMenu(mac.MenuQuit)
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
	win, err := mac.WindowByID("Main")

	if err != nil {
		win = newMainWindow()
	}

	win.Show()
}

func newMainWindow() *mac.Window {
	win := mac.CreateWindow("Main", mac.WindowConfig{
		Width:      1240,
		Height:     720,
		Background: mac.WindowBackgroundDark,
	})

	p := iu.NewPage(&Hello{}, iu.PageConfig{
		CSS: []string{"hello.css"},
	})

	win.Navigate(p)
	return win
}

```

### Stylize the result
```css
html {
    background-size: cover;
    background-repeat: no-repeat;
    background-position: 50% 50%;
    background-image: url("bg.jpg");
}

body {
    color: white;
    background-color: rgba(0, 0, 0, 0.1);    
}

h1 {
    font-weight: 300;
}

input {
    font-size: 11pt;
    border: 0pt;
    border-left: 2pt;
    border-color: lightgrey;
    background-color: transparent;
    color: white;
    border-style: solid;
    box-shadow: none;
    padding: 5pt;
    outline: none;
}

input:focus {
    border-color: dodgerblue;
}

.content {
    margin: 0pt;
}

.hellobox {
    position: absolute;
    top: 50%;
    left: 50%;
    padding: 50px;
    transform: translate(-50%, -50%);
}
```
