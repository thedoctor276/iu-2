package flux

// Callback represents the func signature to be registered in the dispatcher.
type Callback func(p Payload)

type callback struct {
	Callback Callback
	Pending  bool
	Handled  bool
}

func (c *callback) Init() {
	c.Pending = false
	c.Handled = false
}

func (c *callback) Call(p Payload) {
	c.Pending = true
	defer func() { c.Handled = true }()

	c.Callback(p)
}

func newCallback(c Callback) *callback {
	return &callback{
		Callback: c,
	}
}
