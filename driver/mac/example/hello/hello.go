package main

import "github.com/maxence-charriere/iu"

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
