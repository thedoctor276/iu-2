package iu

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

var (
	lastComponentID uint64
	componentMutex  sync.Mutex
)

type Component interface {
	ID() string

	Tag() string

	DataContext() interface{}

	SetDataContext(dataCtx interface{})

	NotifyDataContextChanged()

	View() View

	Init(view View, parent Component)

	Close()

	Render() string

	Sync()
}

type ComponentBase struct {
	Parent Component
	Redraw bool

	id          string
	html        string
	parent      Component
	dataContext interface{}
	view        View
	template    *template.Template
}

func (component *ComponentBase) ID() string {
	return component.id
}

func (component *ComponentBase) DataContext() (dataCtx interface{}) {
	if dataCtx = component.dataContext; dataCtx != nil {
		return
	}

	if component.Parent != nil {
		dataCtx = component.Parent.DataContext()
	}

	return
}

func (component *ComponentBase) View() View {
	return component.view
}

func (component *ComponentBase) OnEvent(eventName string, arg string) (handler string) {
	return fmt.Sprintf("CallEventHandler(this.id, '%v', %v)", eventName, arg)
}

func (component *ComponentBase) Render(c Component) string {
	var buffer bytes.Buffer
	var err error

	if !component.Redraw {
		return component.html
	}

	if err = component.template.Execute(&buffer, c); err != nil {
		iulog.Panic(err)
	}

	component.html = buffer.String()
	component.Redraw = false
	return component.html
}

func NewComponentBase(view View, templateString string) *ComponentBase {
	var component = &ComponentBase{}
	var err error

	component.Redraw = true
	component.id = fmt.Sprintf("iu-%v", nextComponentId())
	component.view = view

	if component.template, err = template.New("").Parse(templateString); err != nil {
		iulog.Panic(err)
	}

	return component
}

func nextComponentId() uint64 {
	componentMutex.Lock()
	defer componentMutex.Unlock()

	lastComponentID++
	return lastComponentID
}
