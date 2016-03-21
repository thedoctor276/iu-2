package iu

type Text struct {
	*ComponentBase
	Value                string
	OnDataContextChanged func(dataCtx interface{})
}

func (comp *Text) Tag() string {
	return "text"
}

func (comp *Text) SetDataContext(dataCtx interface{}) {
	comp.dataContext = dataCtx
	comp.NotifyDataContextChanged()
}

func (comp *Text) NotifyDataContextChanged() {
	if comp.OnDataContextChanged != nil {
		comp.OnDataContextChanged(comp.DataContext())
	}
}

func (comp *Text) Init(view View, parent Component) {
	comp.ComponentBase = NewComponentBase(view, "{{.Value}}")
	comp.Parent = parent
	view.RegisterComponent(comp)
}

func (comp *Text) Close() {
	comp.View().UnregisterComponent(comp)
}

func (comp *Text) Render() string {
	return comp.ComponentBase.Render(comp)
}

func (comp *Text) Sync() {
}
