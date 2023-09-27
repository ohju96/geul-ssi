package service

import (
	"context"
	"geulSsi/app/dto/custom"
	eventDto "geulSsi/app/dto/event"
	"geulSsi/app/repository"
	"geulSsi/bcrypt"
	"geulSsi/ent"
)

type EventService interface {
	CreateEvent(ctx context.Context, req *eventDto.Event) (*eventDto.CreateEventResponse, *custom.Error)
	GetEventList(ctx context.Context, req *eventDto.GetEventReq) (*eventDto.GetEventListResponse, *custom.Error)
	CreateHeart(ctx context.Context, req *eventDto.CreateHeartReq) (*eventDto.CreateHeartResponse, *custom.Error)
}

type eventServiceImpl struct {
	eventRepository repository.EventRepository
	userRepository  repository.UserRepository
	heartRepository repository.HeartRepository
}

func (s eventServiceImpl) CreateHeart(ctx context.Context, req *eventDto.CreateHeartReq) (*eventDto.CreateHeartResponse, *custom.Error) {

	userReq := ent.User{
		Nickname: req.NickName,
		Password: req.Password,
	}

	userRes, err := s.userRepository.FindByUserNickname(ctx, &userReq)
	if err != nil {
		return nil, custom.NotFound
	}

	// 비밀번호가 맞는지 체크한다.
	if err := bcrypt.Compare(userRes.Password, req.Password); err != nil {
		return nil, custom.PasswdNotMatch
	}

	heartReq := ent.Heart{
		EventID: req.EventId,
		Writer:  req.NickName,
	}

	heartRes, err := s.heartRepository.GetHeartByEventIdAndNickName(ctx, &heartReq)
	if !ent.IsNotFound(err) {
		return nil, custom.InternalServerError
	}

	if ent.IsNotFound(err) {
		heartRes, err = s.heartRepository.CreateHeart(ctx, &heartReq)
		if err != nil {
			return nil, custom.InternalServerError
		}

		return &eventDto.CreateHeartResponse{
			Response: *custom.NewSuccess(),
		}, nil
	} else {
		_, err = s.heartRepository.DeleteHeart(ctx, heartRes)
		if err != nil {
			return nil, custom.InternalServerError
		}

		return &eventDto.CreateHeartResponse{
			Response: *custom.NewSuccess(),
		}, nil
	}
}

func (s eventServiceImpl) GetEventList(ctx context.Context, req *eventDto.GetEventReq) (*eventDto.GetEventListResponse, *custom.Error) {

	userReq := ent.User{
		Nickname: req.NickName,
		Password: req.Password,
	}

	userRes, err := s.userRepository.FindByUserNickname(ctx, &userReq)
	if err != nil {
		return nil, custom.NotFound
	}

	// 비밀번호가 맞는지 체크한다.
	if err := bcrypt.Compare(userRes.Password, req.Password); err != nil {
		return nil, custom.PasswdNotMatch
	}

	eventReq := ent.Event{
		Writer: userRes.Nickname,
	}

	eventRes, err := s.eventRepository.GetEventListByNickname(ctx, &eventReq)
	if err != nil {
		return nil, custom.InternalServerError
	}

	var eventList []eventDto.GetEventListPayload

	if len(eventRes) != 0 {
		for _, event := range eventRes {

			heartReq := ent.Heart{EventID: event.ID}
			heartRes, err := s.heartRepository.CountHeartCountByEventId(ctx, &heartReq)
			if err != nil {
				return nil, custom.InternalServerError
			}

			eventList = append(eventList, eventDto.GetEventListPayload{
				EventID:       event.ID,
				EventContents: event.Contents,
				NickName:      event.Writer,
				HeartCount:    heartRes,
			})
		}
	}

	if len(eventList) == 0 {
		eventList = []eventDto.GetEventListPayload{}
	}

	return &eventDto.GetEventListResponse{
		GetEventListPayload: eventList,
		Response:            *custom.NewSuccess(),
	}, nil
}

func (s eventServiceImpl) CreateEvent(ctx context.Context, req *eventDto.Event) (*eventDto.CreateEventResponse, *custom.Error) {

	userReq := ent.User{
		Nickname: req.User,
		Password: req.Password,
	}

	userRes, err := s.userRepository.FindByUserNickname(ctx, &userReq)
	if err != nil {
		return nil, custom.NotFound
	}

	// 비밀번호가 맞는지 체크한다.
	if err := bcrypt.Compare(userRes.Password, req.Password); err != nil {
		return nil, custom.PasswdNotMatch
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

func NewEventService(eventRepository repository.EventRepository, userRepository repository.UserRepository,
	heartRepository repository.HeartRepository,
) EventService {
	return &eventServiceImpl{
		eventRepository: eventRepository,
		userRepository:  userRepository,
		heartRepository: heartRepository,
	}
}
