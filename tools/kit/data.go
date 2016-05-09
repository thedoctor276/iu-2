package kit

import "github.com/maxence-charriere/iu"

// DataContextComponent represents a component which have a data context.
type DataContextComponent interface {
	iu.Component

	DataContext() interface{}

	SetDataContext(ctx interface{}) d
}
