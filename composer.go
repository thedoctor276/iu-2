package iu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

type Composer struct {
	id          string
	html        string
	dataContext interface{}
	view        View
	template    *template.Template
}

func (comp *Composer) ID() string {
	return comp.id
}

func (comp *Composer) View() View {
	return comp.view
}

func (comp *Composer) setView(v View) {
	comp.view = v
}

func (comp *Composer) OnEvent(eventName string, arg string) (handler string) {
	return fmt.Sprintf("CallEventHandler(this.id, '%v', %v)", eventName, arg)
}

func (comp *Composer) Dirty() {
	comp.html = ""
}

func (comp *Composer) Render(c Component) string {
	var buffer bytes.Buffer
	var err error

	if len(comp.html) == 0 {
		if err = comp.template.Execute(&buffer, c); err != nil {
			iulog.Panic(err)
		}

		comp.html = buffer.String()
	}

	return comp.html
}

func NewComposer(tpl string) *Composer {
	var comp = &Composer{}
	var err error

	comp.id = fmt.Sprintf("iu-%v", nextComponentId())

	if comp.template, err = template.New("").Parse(tpl); err != nil {
		iulog.Panic(err)
	}

	return comp
}

func NewComposerWithHTMLTemplate(filename string) *Composer {
	var tpl []byte
	var err error

	if tpl, err = ioutil.ReadFile(filename); err != nil {
		iulog.Panic(err)
	}

	return NewComposer(string(tpl))
}
