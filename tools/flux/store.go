package flux

import (
	"sync"

	"github.com/maxence-charriere/iu-log"
)

// Store represents the interface.
type Store interface {
	DispatchToken() DispatchToken

	setDispatchToken(ID DispatchToken)

	OnDispatch(p Payload)

	AddListener(l Listener) ListenerToken

	RemoveListener(ID ListenerToken)

	Emit(e Event)
}

// Event identifies an event throwed by a store.
type Event string

// Listener represents the func signature to by listened by a store.
type Listener func(e Event)

// ListenerToken identifies a listener within a store.
type ListenerToken uint

// StoreBase represents the basic functionalities of a store.
// It should be embedded in your store implementations.
type StoreBase struct {
	dispatchToken DispatchToken
	lastID        ListenerToken
	listeners     map[ListenerToken]Listener
	mtx           sync.Mutex
}

// DispatchToken returns the token associated to the the store by the dispatcher.
func (s *StoreBase) DispatchToken() DispatchToken {
	return s.dispatchToken
}

func (s *StoreBase) setDispatchToken(ID DispatchToken) {
	s.dispatchToken = ID
}

// AddListener add a listener to listen emitted event by the store.
// It returns a token to be use for unregister the listener.
func (s *StoreBase) AddListener(l Listener) ListenerToken {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.lastID++
	id := s.lastID
	s.listeners[id] = l
	return id
}

// RemoveListener removes a listener.
func (s *StoreBase) RemoveListener(ID ListenerToken) {
	_, ok := s.listeners[ID]
	if !ok {
		iulog.Warnf("%v does not map to a registered listener", ID)
	}

	delete(s.listeners, ID)
}

// Emit emits an event.
func (s *StoreBase) Emit(e Event) {
	for _, l := range s.listeners {
		l(e)
	}
}

// NewStoreBase creates a StoreBase.
// StoreBase should be instantiated only with this function.
func NewStoreBase() *StoreBase {
	return &StoreBase{
		listeners: map[ListenerToken]Listener{},
	}
}
