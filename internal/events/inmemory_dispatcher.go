package events

import "sync"

type InMemoryDispatcher struct{
	handlers map[string][]Handler
	mu       sync.RWMutex
}

func NewInMemoryDispatcher() *InMemoryDispatcher{
	return &InMemoryDispatcher{
		handlers: make(map[string][]Handler),
	}
}

func (d *InMemoryDispatcher) Register(eventType string, handler Handler){
	d.mu.Lock()
	defer d.mu.Lock()
	d.handlers[eventType]=append(d.handlers[eventType], handler)

}

func (d *InMemoryDispatcher) Emit(event DomainEvent){
	d.mu.Lock()
	handlers := d.handlers[event.EventType()]
	d.mu.RUnlock()

	for _,h:= range handlers{
		_=h.Handle(event)
	}
}