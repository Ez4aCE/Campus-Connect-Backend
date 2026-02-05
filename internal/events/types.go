package events

import(
	//"time"
	"github.com/google/uuid"
)

const (
	EventTypeEventPublished = "EVENT_PUBLISHED"
	EventTypeEventCancelled = "EVENT_CANCELLED"
	EventTypeMembershipApproved ="MEMBERSHIP_APPROVED"
	EventTypeClubApproved = "CLUB_APPROVED"
	EventTypeRegistrationConfirmed="REGISTRATION_CONFIRMED"
)


//Event published
type EventPublished struct{
	BaseEvent
	EventID uuid.UUID
	ActorID uuid.UUID
}

func (e EventPublished) EventType() string{
	return EventTypeEventPublished
}


//Event cancelled
type EventCancelled struct{
	BaseEvent
	EventID uuid.UUID
	ActorID uuid.UUID
}

func (e EventCancelled) EventType() string{
	return EventTypeEventCancelled
}

//membership approved
type MembershipApproved struct{
	BaseEvent
	ClubID uuid.UUID
	UserID uuid.UUID
	ApprovedBy uuid.UUID
}

func (e MembershipApproved) EventType() string{
	return EventTypeMembershipApproved
}

//club approved
type ClubApproved struct{
	BaseEvent
	ClubID uuid.UUID
	OwnerID uuid.UUID
	ApprovedBy uuid.UUID
}

func (e ClubApproved) EventType() string{
	return EventTypeClubApproved
}

//Registration Confirmed
type RegistrationConfirmed struct{
	BaseEvent
	EventID uuid.UUID
	UserID uuid.UUID
}

func (e RegistrationConfirmed) EventType() string{
	return EventTypeRegistrationConfirmed
}



