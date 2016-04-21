package flux

// Payload represents a payload to be broadcasted by a dispatcher.
type Payload struct {
	Action Action
	Data   interface{}
}

// Action represents an identifier of the action to be performed by a store.
type Action string
