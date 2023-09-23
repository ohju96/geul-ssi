package service

import "geulSsi/app/repository"

type EventService interface {
}

type eventServiceImpl struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventServiceImpl{
		eventRepository: eventRepository,
	}
}
