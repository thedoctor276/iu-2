package iu

import (
	"fmt"
	"testing"
)

func testComponentSetDataContext(t *testing.T, comp Component) {
	var ctx = &EmptyContext{}
	var view = &Page{
		Title: "Test",
		Body: []Component{
			comp,
		},
	}

	view.Init(ctx)
	comp.SetDataContext("AnyDataContext")
}

func testComponentInit(t *testing.T, comp Component) {
	var ctx = &EmptyContext{}
	var view = &Page{
		Title: "Test",
		Body: []Component{
			comp,
		},
	}

	view.Init(ctx)
}

func testComponentClose(t *testing.T, comp Component) {
	var ctx = &EmptyContext{}
	var view = &Page{
		Title: "Test",
		Body: []Component{
			comp,
		},
	}

	view.Init(ctx)
	view.Body = nil
	comp.Close()
}

func testComponentRender(t *testing.T, comp Component) {
	var ctx = &EmptyContext{}
	var view = &Page{
		Title: "Test",
		Body: []Component{
			comp,
		},
	}

	view.Init(ctx)

	// render full
	t.Log(comp.Render())

	// render cache
	comp.Render()
}

func testComponentSync(t *testing.T, comp Component) {
	var ctx = &EmptyContext{}
	var view = &Page{
		Title: "Test",
		Body: []Component{
			comp,
		},
	}

	view.Init(ctx)
	comp.Sync()
}

func TestComponentBaseID(t *testing.T) {
	var view = &Page{}
	var component *ComponentBase
	var expected = fmt.Sprintf("iu-%v", 1)

	lastComponentID = 0
	defer func() { lastComponentID = 0 }()

	component = NewComponentBase(view, CommonComponentTemplate())

	if id := component.ID(); id != expected {
		t.Errorf("id should be %v: %v", expected, id)
	}
}

func TestComponentBaseDataContext(t *testing.T) {
	var div = divTest()

	testComponentSetDataContext(t, div)

	if dataCtx := div.Content[0].DataContext(); dataCtx != div.DataContext() {
		t.Errorf("dataCtx should be %v: %v", div.DataContext(), dataCtx)
	}
}

func TestComponentBaseView(t *testing.T) {
	var view = &Page{}
	var component = NewComponentBase(view, CommonComponentTemplate())
	var expected = "CallEventHandler(this.id, 'OnClick', event)"

	if onevent := component.OnEvent("OnClick", "event"); onevent != expected {
		t.Errorf("onevent should be %v: %v", expected, onevent)
	}
}

func TestComponentBaseOnEvent(t *testing.T) {
	var view = &Page{}
	var component = NewComponentBase(view, CommonComponentTemplate())

	if v := component.View(); v != view {
		t.Errorf("p should be %v: %v", view, v)
	}
}

func TestNewComponentBase(t *testing.T) {
	var view = &Page{}

	NewComponentBase(view, CommonComponentTemplate())
}

func TestNextComponentId(t *testing.T) {
	lastComponentID = 0
	defer func() { lastComponentID = 0 }()

	nextComponentId()

	if lastComponentID != 1 {
		t.Errorf("lastComponentID should be 1: %v", lastComponentID)
	}
}
