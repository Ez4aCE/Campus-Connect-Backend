package handlers

import (
	"campus-connect-backend/internal/events"
	"campus-connect-backend/internal/services"
	"encoding/json"

)

type NotificationHandler struct {
	notificationService *services.NotificationService
}

func NewNotificationHandler(ns *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: ns,
	}
}

func (h *NotificationHandler) Handle(event events.DomainEvent) error {

	switch e := event.(type) {
	case events.EventPublished:
		return h.handleEventPublished(e)
	
	case  events.EventCancelled:
		return h.handleEventCancelled(e)

	case events.MembershipApproved:
		return h.handleMembershipApproved(e)

	case events.ClubApproved:
		return h.handleClubApproved(e)

	case events.RegistrationConfirmed:
		return h.handleRegistrationConfirmed(e)

	}
	return nil
}

func (h *NotificationHandler) handleEventPublished(e events.EventPublished) error{
	metadata,_:=json.Marshal(map[string]interface{}{
		"event_id":e.EventID,
		"redirect":"/events/"+e.EventID.String(),
	})

	title:="New Event Published"
	message :="A new event has been published. check it out!"

	return h.notificationService.Broadcast(
		title,
		message,
		events.EventTypeEventPublished,
		metadata,
		&e.ActorID,

	)
}

func (h *NotificationHandler) handleEventCancelled(e events.EventCancelled) error{
	metadata ,_:= json.Marshal(map[string]interface{}{
		"event_id":e.EventID,
	})

	title:="Event Cancelled"
	message :="An event you were interested in has been cancelled."

	return h.notificationService.Broadcast(
		title,
		message,
		events.EventTypeEventCancelled,
		metadata,
		&e.ActorID,
	)
}

func (h *NotificationHandler) handleMembershipApproved(e events.MembershipApproved) error{
	metadata,_:=json.Marshal(map[string]interface{}{
		"clubd_id":e.ClubID,
	})

	title:= "Membership Aprroved"
	message:="Your request to join the club has been approved."

	return h.notificationService.Broadcast(
		title,
		message,
		events.EventTypeMembershipApproved,
		metadata,
		&e.UserID,
	)
}

func (h *NotificationHandler) handleClubApproved(e events.ClubApproved) error{
	metadata, _:=json.Marshal(map[string]interface{}{
		"club_id":e.ClubID,
	})

	title:="Club Approved"
	message:= "Your club has been approved by admin."

	return h.notificationService.CreateForUser(
		title,
		message,
		events.EventTypeClubApproved,
		metadata,
		&e.ApprovedBy,
		e.OwnerID,
	)
}


func (h *NotificationHandler) handleRegistrationConfirmed(e events.RegistrationConfirmed) error{
	metadata,_:=json.Marshal(map[string]interface{}{
		"event_id":e.EventID,
	})

	title:="Registration Confirmed"
	message:="Your have succesfully registered for the event."

	return h.notificationService.CreateForUser(
		title,
		message,
		events.EventTypeRegistrationConfirmed,
		metadata,
		nil,
		e.UserID,
	)
}


