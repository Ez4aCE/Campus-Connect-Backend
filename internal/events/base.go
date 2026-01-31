package events

import "time"

type BaseEvent struct{
	Occured time.Time
}

func (e BaseEvent) OccuredAt() time.Time{
	return e.Occured
}