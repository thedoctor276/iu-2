package iu

// func TestPageContext(t *testing.T) {
// 	var context = &EmptyContext{}
// 	var page = &Page{Title: "Test"}

// 	page.Init(context)

// 	if ctx := page.Context(); ctx != context {
// 		t.Error("ctx should be %v: %v", context, ctx)
// 	}
// }

// func TestPageFrameworkJS(t *testing.T) {
// 	var page = &Page{Title: "Test"}

// 	if js := page.FrameworkJS(); js != FrameworkJS() {
// 		t.Error("js should be %v: %v", FrameworkJS(), js)
// 	}
// }

// func TestPageComponent(t *testing.T) {
// 	var div = &Div{}
// 	var ctx = &EmptyContext{}
// 	var page = &Page{
// 		Title: "Test",
// 		Body:  []Component{div},
// 	}

// 	page.Init(ctx)

// 	if comp := page.Component(div.ID()); comp != div {
// 		t.Errorf("comp: %p should be div: %p", comp, div)
// 	}
// }

// func TestPageNonexistentComponent(t *testing.T) {
// 	var ctx = &EmptyContext{}
// 	var page = &Page{
// 		Title: "Test",
// 	}

// 	defer func() { recover() }()

// 	page.Init(ctx)
// 	page.Component("42")
// 	t.Error("should have panic")
// }

// func TestPageRegisterNoneinitializedComponent(t *testing.T) {
// 	var div = &Div{ComponentBase: &ComponentBase{}}
// 	var page = &Page{
// 		Title: "Test",
// 	}

// 	defer func() { recover() }()

// 	page.RegisterComponent(div)
// 	t.Error("should have panic")
// }

// func TestPageInit(t *testing.T) {
// 	var context = &EmptyContext{}
// 	var page = &Page{Title: "Test"}

// 	page.Init(context)
// }

// func TestPageInitNoTitle(t *testing.T) {
// 	var context = &EmptyContext{}
// 	var page = &Page{}

// 	defer func() { recover() }()

// 	page.Init(context)
// 	t.Error("should have panic")
// }

// func TestPageInitNilContext(t *testing.T) {
// 	var page = &Page{Title: "Test"}

// 	defer func() { recover() }()

// 	page.Init(nil)
// 	t.Error("should have panic")
// }

// func TestPageRender(t *testing.T) {
// 	var context = &EmptyContext{}

// 	var page = &Page{
// 		Title: "Page",
// 		Lang:  "en",
// 		CSS: []string{
// 			"test.css",
// 			"test2.css",
// 		},
// 	}

// 	page.Init(context)
// 	t.Log(page.Render())
// }

// func TestPageRenderNoTemplate(t *testing.T) {
// 	var page = &Page{Title: "Test"}

// 	defer func() { recover() }()

// 	page.Render()
// 	t.Error("should have panic")
// }
