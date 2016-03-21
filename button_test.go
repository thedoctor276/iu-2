package iu

import "testing"

func buttonTest() *Button {
	return &Button{
		Title:                "Test",
		Lang:                 "en",
		Class:                "MyButton",
		Dir:                  DirLtr,
		TabIndex:             1,
		ContentEditable:      true,
		Draggable:            true,
		Hidden:               true,
		Spellcheck:           true,
		Autofocus:            true,
		Disabled:             true,
		OnDataContextChanged: func(dataCtx interface{}) {},
		OnClick:              func(event MouseEvent) {},
		OnContextMenu:        func(event MouseEvent) {},
		OnDblClick:           func(event MouseEvent) {},
		OnMouseDown:          func(event MouseEvent) {},
		OnMouseEnter:         func(event MouseEvent) {},
		OnMouseLeave:         func(event MouseEvent) {},
		OnMouseMove:          func(event MouseEvent) {},
		OnMouseOver:          func(event MouseEvent) {},
		OnMouseOut:           func(event MouseEvent) {},
		OnMouseUp:            func(event MouseEvent) {},
		OnDrag:               func(event MouseEvent) {},
		OnDragEnd:            func(event MouseEvent) {},
		OnDragEnter:          func(event MouseEvent) {},
		OnDragLeave:          func(event MouseEvent) {},
		OnDragOver:           func(event MouseEvent) {},
		OnDragStart:          func(event MouseEvent) {},
		OnDrop:               func(event MouseEvent) {},
		OnScroll:             func() {},
		OnWheel:              func(event WheelEvent) {},
		OnKeyDown:            func(event KeyboardEvent) {},
		OnKeyPress:           func(event KeyboardEvent) {},
		OnKeyUp:              func(event KeyboardEvent) {},
		OnCopy:               func() {},
		OnCut:                func() {},
		OnPaste:              func() {},
		OnBlur:               func() {},
		OnFocus:              func() {},
		Content: []Component{
			&Button{},
		},
	}
}

func TestTestButtonSetDataContext(t *testing.T) {
	testComponentSetDataContext(t, buttonTest())
}

func TestButtonInit(t *testing.T) {
	testComponentInit(t, buttonTest())
}

func TestButtonClose(t *testing.T) {
	testComponentClose(t, buttonTest())
}

func TestButtonRender(t *testing.T) {
	testComponentRender(t, buttonTest())
}

func TestButtonSync(t *testing.T) {
	testComponentSync(t, buttonTest())
}
