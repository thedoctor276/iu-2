package kit

import (
	"reflect"

	"github.com/maxence-charriere/iu"
	"github.com/maxence-charriere/iu-log"
)

// Repeater is a component that allow to render a slice of components based on a data source.
type Repeater struct {
	children []iu.Component
	mapper   DataMapper
}

// OnDismount is the action called when the repeater is about to be dismounted.
// This is for component implementation, should not be called by the package user.
func (r *Repeater) OnDismount() {
	r.clearChildren()
}

func (r *Repeater) clearChildren() {
	for _, c := range r.children {
		iu.DismountComponent(c)
	}

	r.children = nil
}

// Template returns the repeater template.
func (r *Repeater) Template() string {
	return `
<div class="Repeater" style="width:100%;height:100%">
    {{range .Children}}
        {{.Render}}
    {{end}}
</div>
`
}

// Children returns the slice of components which mirrors the source of the repeater.
func (r *Repeater) Children() []iu.Component {
	return r.children
}

// SetSource sets the source of the repeater.
// Will panic if the source is not a slice or an array.
// Should not contain components.
func (r *Repeater) SetSource(src interface{}) {
	v := reflect.ValueOf(src)

	if k := v.Kind(); k != reflect.Array && k != reflect.Slice {
		iulog.Panicf("src must be a slice or an array: %T", src)
	}

	r.clearChildren()
	d := iu.DriverByComponent(r)

	for i := 0; i < v.Len(); i++ {
		s := v.Index(i).Interface()
		c := r.mapper(s)
		iu.MountComponent(c, d)
		r.children = append(r.children, c)
	}
}

// NewRepeater creates an instance of repeater.
func NewRepeater(m DataMapper) *Repeater {
	return &Repeater{
		mapper: m,
	}
}
