package flux

import (
	"sync"

	"github.com/maxence-charriere/iu-log"
)

// DispatchToken identifies a callback wihthin a Dispatcher.
type DispatchToken uint

type dispatcher struct {
	callbacks      map[DispatchToken]*callback
	dispatching    bool
	lastID         DispatchToken
	pendingPayload Payload
	mtx            sync.Mutex
}

// Register registers a callback to be invoked with every dispatched payload.
// Returns a token that can be used with `flux.Dispatcher.WaitFor()`.
func (disp *dispatcher) Register(c Callback) DispatchToken {
	disp.mtx.Lock()
	defer disp.mtx.Unlock()

	disp.lastID++
	id := disp.lastID
	disp.callbacks[id] = newCallback(c)
	return id
}

// Unregister removes a callback based on its token.
func (disp *dispatcher) Unregister(ID DispatchToken) {
	_, ok := disp.callbacks[ID]
	if !ok {
		iulog.Warnf("%v does not map to a registered callback", ID)
	}

	delete(disp.callbacks, ID)
}

// WaitFor waits for the callbacks specified to be invoked before continuing execution
// of the current callback. This method should only be used by a callback in
// response to a dispatched payload.
func (disp *dispatcher) WaitFor(IDs ...DispatchToken) {
	if !disp.dispatching {
		iulog.Panic("must be invoked while dispatching")
	}

	for _, ID := range IDs {
		c, ok := disp.callbacks[ID]
		if !ok {
			iulog.Panicf("%v does not map to a registered callback", ID)
		}

		if c.Pending {
			iulog.Warn(c)
			if !c.Handled {
				iulog.Panicf("circular dependency detected while waiting for %v", ID)
			}

			continue
		}

		c.Call(disp.pendingPayload)
	}

}

// Dispatch dispatches a payload to all registered callbacks.
func (disp *dispatcher) Dispatch(p Payload) {
	if disp.dispatching {
		iulog.Panic("cannot dispatch in the middle of a dispatch")
	}

	disp.startDispatching(p)
	defer disp.stopDispatching()

	for _, c := range disp.callbacks {
		if c.Pending {
			continue
		}

		c.Call(disp.pendingPayload)
	}
}

func (disp *dispatcher) startDispatching(p Payload) {
	for _, c := range disp.callbacks {
		c.Init()
	}

	disp.pendingPayload = p
	disp.dispatching = true
}

func (disp *dispatcher) stopDispatching() {
	disp.pendingPayload = Payload{}
	disp.dispatching = false
}

func newDispatcher() *dispatcher {
	return &dispatcher{
		callbacks: map[DispatchToken]*callback{},
	}
}
