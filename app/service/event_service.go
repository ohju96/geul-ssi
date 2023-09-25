package service

import (
	"context"
	"geulSsi/app/dto/custom"
	eventDto "geulSsi/app/dto/event"
	"geulSsi/app/repository"
	"geulSsi/ent"
)

type EventService interface {
	CreateEvent(ctx context.Context, req *eventDto.Event) (*eventDto.CreateEventResponse, *custom.Error)
}

type eventServiceImpl struct {
	eventRepository repository.EventRepository
	userRepository  repository.UserRepository
}

func (s eventServiceImpl) CreateEvent(ctx context.Context, req *eventDto.Event) (*eventDto.CreateEventResponse, *custom.Error) {

	userReq := ent.User{
		Nickname:   req.User,
		SecretCode: req.Password,
	}

	_, err := s.userRepository.ExistBySecretCodeAndNickname(ctx, &userReq)
	if err != nil {
		return nil, custom.NotFound
	}

	userRes, err := s.userRepository.FindByNickname(ctx, &userReq)
	if err != nil {
		return nil, custom.NotFound
	}

	tx, customError := custom.TxGenerate(ctx)
	if customError != nil {
		return nil, customError
	}

	eventReq := ent.Event{
		Contents: req.User,
		Writer:   userRes.Nickname,
	}

	eventRes, err := s.eventRepository.CreateEvent(ctx, &eventReq, tx)
	if err != nil {
		custom.TxRollback(tx)
		return nil, custom.InternalServerError
	}

	if customError = custom.TxCommit(tx); customError != nil {
		return nil, customError
	}

	return &eventDto.CreateEventResponse{
		Payload: eventDto.CreateEventPayload{
			EventID:  eventRes.ID,
			NickName: eventRes.Writer,
		},
		Response: *custom.NewSuccess(),
	}, nil
}

func NewEventService(eventRepository repository.EventRepository, userRepository repository.UserRepository) EventService {
	return &eventServiceImpl{
		eventRepository: eventRepository,
		userRepository:  userRepository,
	}
}
