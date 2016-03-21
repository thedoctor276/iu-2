package iu

import "github.com/maxence-charriere/iu-log"

const (
	InputText     InputType = "text"
	InputPassword InputType = "password"
	InputCheckbox InputType = "checkbox"
	InputRadio    InputType = "radio"
	InputNumber   InputType = "number"
	InputRange    InputType = "range"
)

type Input struct {
	*ComponentBase
	Title           string
	Lang            string
	Class           string
	Max             string
	Min             string
	Step            string
	Value           string
	Placeholder     string
	Type            InputType
	Dir             DirAttribute
	TabIndex        uint
	MaxLength       uint
	ContentEditable bool
	Draggable       bool
	Hidden          bool
	Spellcheck      bool
	Autofocus       bool
	Checked         bool
	Disabled        bool

	OnDataContextChanged func(dataCtx interface{})
	OnClick              func(event MouseEvent)
	OnContextMenu        func(event MouseEvent)
	OnDblClick           func(event MouseEvent)
	OnMouseDown          func(event MouseEvent)
	OnMouseEnter         func(event MouseEvent)
	OnMouseLeave         func(event MouseEvent)
	OnMouseMove          func(event MouseEvent)
	OnMouseOver          func(event MouseEvent)
	OnMouseOut           func(event MouseEvent)
	OnMouseUp            func(event MouseEvent)
	OnDrag               func(event MouseEvent)
	OnDragEnd            func(event MouseEvent)
	OnDragEnter          func(event MouseEvent)
	OnDragLeave          func(event MouseEvent)
	OnDragOver           func(event MouseEvent)
	OnDragStart          func(event MouseEvent)
	OnDrop               func(event MouseEvent)
	OnScroll             func()
	OnWheel              func(event WheelEvent)
	OnKeyDown            func(event KeyboardEvent)
	OnKeyPress           func(event KeyboardEvent)
	OnKeyUp              func(event KeyboardEvent)
	OnCopy               func()
	OnCut                func()
	OnPaste              func()
	OnBlur               func()
	OnFocus              func()
	OnSelect             func()
	OnChange             func(value string)
	OnCheck              func(check bool)
}

type InputType string

func (comp *Input) Tag() string {
	return "input"
}

func (comp *Input) SetDataContext(dataCtx interface{}) {
	comp.dataContext = dataCtx
	comp.NotifyDataContextChanged()
}

func (comp *Input) NotifyDataContextChanged() {
	if comp.OnDataContextChanged != nil {
		comp.OnDataContextChanged(comp.DataContext())
	}
}

func (comp *Input) Init(view View, parent Component) {
	comp.ComponentBase = NewComponentBase(view, InputTemplate())
	comp.Parent = parent
	view.RegisterComponent(comp)

	if len(comp.Type) == 0 {
		comp.Type = InputText
	}

	if comp.OnChange != nil && comp.OnCheck != nil {
		iulog.Panicf("%v with id = %v: can't set both OnChanged and OnChecked")
	}
}

func (comp *Input) Close() {
	comp.View().UnregisterComponent(comp)
}

func (comp *Input) Render() string {
	return comp.ComponentBase.Render(comp)
}

func (comp *Input) Sync() {
	var ctx = comp.View().Context()

	comp.Redraw = true
	ctx.InjectComponent(comp)
}
