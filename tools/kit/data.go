package kit

import "github.com/maxence-charriere/iu"

// DataMapper represents a handler to determine which component will be created for each source element.
// The handler should create and initialize the component it returns.
type DataMapper func(d interface{}) iu.Component
