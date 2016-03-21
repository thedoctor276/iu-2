package iu

import (
	"bytes"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

type Page struct {
	Title    string
	Lang     string
	CSS      []string
	Body     []Component
	OnLoaded func()

	context    Context
	loaded     bool
	template   *template.Template
	components map[string]Component
}

func (page *Page) Context() Context {
	return page.context
}

func (page *Page) FrameworkJS() string {
	return FrameworkJS()
}

func (page *Page) Component(id string) (component Component) {
	var ok bool

	if component, ok = page.components[id]; !ok {
		iulog.Panicf("component with id = %v is not found in page %v", id, page.Title)
	}

	return
}

func (page *Page) RegisterComponent(component Component) {
	var id string

	if id = component.ID(); len(id) == 0 {
		iulog.Panicf("can't register component: %v is not initialized", component.Tag())
	}

	page.components[id] = component
}

func (page *Page) UnregisterComponent(component Component) {
	delete(page.components, component.ID())
}

func (page *Page) Init(ctx Context) {
	var err error

	if len(page.Title) == 0 {
		iulog.Panic("page.Title should not be empty")
	}

	if ctx == nil {
		iulog.Panic("ctx can't be nil")
	}

	if page.template, err = template.New("").Parse(PageTemplate()); err != nil {
		iulog.Panic(err)
	}

	page.context = ctx
	page.components = make(map[string]Component)

	for _, component := range page.Body {
		component.Init(page, nil)
		page.RegisterComponent(component)
	}
}

func (page *Page) Render() string {
	var buffer bytes.Buffer
	var err error

	if err = page.template.Execute(&buffer, page); err != nil {
		iulog.Panic(err)
	}

	return buffer.String()
}
