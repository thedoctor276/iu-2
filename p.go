package iu

type P struct {
	*ComponentBase
	Title                string
	Lang                 string
	Class                string
	Dir                  DirAttribute
	TabIndex             uint
	ContentEditable      bool
	Draggable            bool
	Hidden               bool
	Spellcheck           bool
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
	Content              []Component
}

func (comp *P) Tag() string {
	return "p"
}

func (comp *P) SetDataContext(dataCtx interface{}) {
	comp.dataContext = dataCtx
	comp.NotifyDataContextChanged()
}

func (comp *P) NotifyDataContextChanged() {
	if comp.OnDataContextChanged != nil {
		comp.OnDataContextChanged(comp.DataContext())
	}

	for _, component := range comp.Content {
		component.NotifyDataContextChanged()
	}
}

func (comp *P) Init(view View, parent Component) {
	comp.ComponentBase = NewComponentBase(view, CommonComponentTemplate())
	comp.Parent = parent
	view.RegisterComponent(comp)

	for _, component := range comp.Content {
		component.Init(view, comp)
	}
}

func (comp *P) Close() {
	comp.View().UnregisterComponent(comp)

	for _, component := range comp.Content {
		component.Close()
	}
}

func (comp *P) Render() string {
	return comp.ComponentBase.Render(comp)
}

func (comp *P) Sync() {
	var ctx = comp.View().Context()

	comp.Redraw = true
	ctx.InjectComponent(comp)
}
