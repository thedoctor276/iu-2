package iu

type Br struct {
	*ComponentBase
}

func (comp *Br) Tag() string {
	return "Br"
}

func (comp *Br) SetDataContext(dataCtx interface{}) {
}

func (comp *Br) NotifyDataContextChanged() {
}

func (comp *Br) Init(view View, parent Component) {
	comp.ComponentBase = NewComponentBase(view, "<br>")
	comp.Parent = parent
	view.RegisterComponent(comp)
}

func (comp *Br) Close() {
	comp.View().UnregisterComponent(comp)
}

func (comp *Br) Render() string {
	return comp.ComponentBase.Render(comp)
}

func (comp *Br) Sync() {
}
