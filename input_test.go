package iu

import "testing"

func inputTest() *Input {
	return &Input{
		Title:                "Test",
		Lang:                 "en",
		Class:                "MyInput",
		Max:                  "42",
		Min:                  "0",
		Step:                 "1",
		Value:                "input test",
		Dir:                  DirLtr,
		TabIndex:             1,
		ContentEditable:      true,
		Draggable:            true,
		Hidden:               true,
		Spellcheck:           true,
		Autofocus:            true,
		Checked:              true,
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
		OnSelect:             func() {},
		OnChanged:            func(value string) {},
	}
}

func TestTestInputSetDataContext(t *testing.T) {
	testComponentSetDataContext(t, inputTest())
}

func TestInputInit(t *testing.T) {
	testComponentInit(t, inputTest())
}

func TestInputInitConflictingOnChange(t *testing.T) {
	var inputOnChangeConflict = &Input{
		OnChanged: func(value string) {},
		OnChecked: func(checked bool) {},
	}
	defer func() { recover() }()

	testComponentInit(t, inputOnChangeConflict)
	t.Error("should have panic")
}

func TestInputClose(t *testing.T) {
	testComponentClose(t, inputTest())
}

func TestInputRender(t *testing.T) {
	testComponentRender(t, inputTest())
}

func TestInputRenderCheckable(t *testing.T) {
	var inputCheckableTest = &Input{
		Type:      InputCheckbox,
		OnChecked: func(checked bool) {},
	}

	testComponentRender(t, inputCheckableTest)
}

func TestInputSync(t *testing.T) {
	testComponentSync(t, inputTest())
}
