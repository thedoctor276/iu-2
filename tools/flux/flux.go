package flux

import "github.com/maxence-charriere/iu-log"

var (
	mainDispatcher = newDispatcher()
)

// RegisterStore registers a store to the flux dispatcher.
// The store listens for actions to be dispatched in OnDispatch(p flux.Payload).
func RegisterStore(s Store) {
	_, ok := mainDispatcher.callbacks[s.DispatchToken()]
	if ok {
		iulog.Warnf("%#v is already registered under token %v", s, s.DispatchToken())
		return
	}

	id := mainDispatcher.Register(s.OnDispatch)
	s.setDispatchToken(id)
}

// UnregisterStore removes a store from the flux dispatcher.
// The store stops listening dispatched actions.
func UnregisterStore(s Store) {
	_, ok := mainDispatcher.callbacks[s.DispatchToken()]
	if !ok {
		iulog.Warnf("%#v is already unregistered", s)
		return
	}

	mainDispatcher.Unregister(s.DispatchToken())
}

// WaitFor waits that the specified stores finish handling dispatch actions
// before continuing the current dispatch.
func WaitFor(s ...Store) {
	ids := make([]DispatchToken, len(s))

	for i := 0; i < len(s); i++ {
		ids[i] = s[i].DispatchToken()
	}

	mainDispatcher.WaitFor(ids...)
}

// Dispatch dispatches a payload..
func Dispatch(p Payload) {
	mainDispatcher.Dispatch(p)
}

// DispatchAction dispatches an action.
func DispatchAction(a Action) {
	mainDispatcher.Dispatch(Payload{
		Action: a,
	})

}

// DispatchActionWithData dispatches an action with data.
func DispatchActionWithData(a Action, d interface{}) {
	mainDispatcher.Dispatch(Payload{
		Action: a,
		Data:   d,
	})
}
