package events

//import "sync"

type Handler interface{
	Handle(event DomainEvent) error
}

type Dispatcher interface{
	Emit(event DomainEvent)
	Register(eventType string , handler Handler)
}