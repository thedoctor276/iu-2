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
<html lang="{{.Lang}}">
<head>
    <title>{{if .Title}}{{.Title}}{{else}}iu{{end}}</title>
    <meta charset="utf-8" /> 
    
    <style media="{{if .Media}}{{.Media}}{{else}}screen{{end}}" type="text/css">
        html {
            height: 100%;
            width: 100%;
            margin: 0pt;
            background-color: transparent;
        }
        
        body {
            height: 100%;
            width: 100%;
            margin: 0pt;
            font-family: "Helvetica Neue", "Segoe UI";
            font-size: 11pt;
            overflow-x: hidden;
            overflow-y: hidden;
            background-color: transparent;
        }
    </style>
    
{{range .CSS}}
    <link rel="stylesheet" href="{{.}}" />{{end}}
</head>
<body oncontextmenu="event.preventDefault()">
{{.MainView.Render}}

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
	Media string
	CSS   []string
	JS    []string
}

type Page struct {
	Context Context
	OnLoad  func()

	config   PageConfig
	mainView View
	template *template.Template
}

func (page *Page) Config() PageConfig {
	return page.config
}

func (page *Page) MainView() View {
	return page.mainView
}

func (page *Page) Render() string {
	var buffer bytes.Buffer
	var err error

	m := map[string]interface{}{
		"Title":       page.config.Title,
		"Lang":        page.config.Lang,
		"Media":       page.config.Media,
		"CSS":         page.config.CSS,
		"JS":          page.config.JS,
		"MainView":    compoM.Component(page.mainView),
		"FrameworkJS": FrameworkJS(),
	}

	if page.template == nil {
		if page.template, err = template.New("").Parse(pageTpl); err != nil {
			iulog.Panic(err)
		}
	}

	if err = page.template.Execute(&buffer, m); err != nil {
		iulog.Panic(err)
	}

	return buffer.String()
}

func (page *Page) Close() {
	ForRangeViews(page.mainView, func(v View) (err error) {
		err = UnregisterView(v)
		return
	})
}

func NewPage(mainView View, config PageConfig) *Page {
	p := &Page{
		config:   config,
		mainView: mainView,
	}

	ForRangeViews(mainView, func(v View) (err error) {
		if err = RegisterView(v); err != nil {
			return
		}

		c := compoM.Component(v)
		c.page = p
		return
	})

	return p
}
