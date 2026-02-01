package events

import "time"


type DomainEvent interface{
	EventType() string
	OccuredAt() time.Time
}