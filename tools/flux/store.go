package flux

import "sync"

type Store interface {
	OnDispatch(a Action)

	ID() StoreID

	SetID(ID StoreID)
}

type StoreID int

type StoreBase struct {
	id             StoreID
	listeners      map[ListenerID]Listener
	lastListenerID ListenerID
	mtx            sync.Mutex
}

func (s *StoreBase) ID() StoreID {
	return s.id
}

func (s *StoreBase) SetID(ID StoreID) {
	s.id = ID
}

func (s *StoreBase) AddListener(l Listener) ListenerID {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.lastListenerID++
	id := s.lastListenerID

	s.listeners[id] = l
	return id
}

func (s *StoreBase) RemoveListener(ID ListenerID) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.listeners, ID)
}

// Emit sends an event to all registered listeners.
func (s *StoreBase) Emit(e Event) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for _, l := range s.listeners {
		l(e)
	}
}

// NewStoreBase creates a new instance of StoreBase.
func NewStoreBase() *StoreBase {
	return &StoreBase{
		listeners: map[ListenerID]Listener{},
	}
}
