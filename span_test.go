package iu

import "testing"

func spanTest() *Span {
	return &Span{
		Title:                "Test",
		Lang:                 "en",
		Class:                "MySpan",
		Dir:                  DirLtr,
		TabIndex:             1,
		ContentEditable:      true,
		Draggable:            true,
		Hidden:               true,
		Spellcheck:           true,
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
			&Span{},
		},
	}
}

func TestTestSpanSetDataContext(t *testing.T) {
	testComponentSetDataContext(t, spanTest())
}

func TestSpanInit(t *testing.T) {
	testComponentInit(t, spanTest())
}

func TestSpanClose(t *testing.T) {
	testComponentClose(t, spanTest())
}

func TestSpanRender(t *testing.T) {
	testComponentRender(t, spanTest())
}

func TestSpanSync(t *testing.T) {
	testComponentSync(t, spanTest())
}
