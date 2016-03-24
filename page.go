package iu

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/maxence-charriere/iu-log"
)

var (
	pageTpl = strings.Trim(`
<!doctype html>
<html lang="{{.Config.Lang}}">
<head>
    <title>{{if .Config.Title}}{{.Config.Title}}{{else}}iu{{end}}</title>
    <meta charset="utf-8" /> 
{{range .Config.CSS}}
    <link rel="stylesheet" href="{{.}}" />{{end}}
</head>
<body>
{{.MainComponent.Render}}

<script>
{{.FrameworkJS}}
</script>
{{range .Config.JS}}
<script src="{{.}}"></script>{{end}}
</body>
</html>
`, " \t\r\n")
)

type PageConfig struct {
	Title string
	Lang  string
	CSS   []string
	JS    []string
}

type Page struct {
	config        PageConfig
	mainComponent Component
	context       Context
	template      *template.Template
	components    map[string]Component
}

func (page *Page) Config() PageConfig {
	return page.config
}

func (page *Page) MainComponent() Component {
	return page.mainComponent
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
		iulog.Panicf("component with id = %v is not found in page %v", id, page.config.Title)
	}

	return
}

func (page *Page) RegisterComponent(component Component) {
	var id string

	if id = component.ID(); len(id) == 0 {
		iulog.Panicf("can't register component: %p is not initialized", component)
	}

	page.components[id] = component
}

func (page *Page) UnregisterComponent(component Component) {
	delete(page.components, component.ID())
}

func (page *Page) Init(ctx Context) {
	var err error

	if ctx == nil {
		iulog.Panic("ctx can't be nil")
	}

	if page.template, err = template.New("").Parse(pageTpl); err != nil {
		iulog.Panic(err)
	}

	page.context = ctx
	//PairViewComponent(page, page.MainComponent())

}

func (page *Page) Render() string {
	var buffer bytes.Buffer
	var err error

	if err = page.template.Execute(&buffer, page); err != nil {
		iulog.Panic(err)
	}

	return buffer.String()
}

func NewPage(mainComponent Component, config PageConfig) *Page {
	return &Page{
		config:        config,
		mainComponent: mainComponent,
		components:    map[string]Component{},
	}
}
