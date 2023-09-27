package repository

import (
	"context"
	"geulSsi/ent"
	"geulSsi/ent/event"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, req *ent.Event, tx *ent.Tx) (*ent.Event, error) // 생성
	GetEventListByNickname(ctx context.Context, req *ent.Event) ([]*ent.Event, error)
}

type eventRepositoryImpl struct {
	ent.Client
}

func (r eventRepositoryImpl) GetEventListByNickname(ctx context.Context, req *ent.Event) ([]*ent.Event, error) {
	return r.Event.Query().Where(event.WriterEQ(req.Writer)).All(ctx)
}

func (r eventRepositoryImpl) CreateEvent(ctx context.Context, req *ent.Event, tx *ent.Tx) (*ent.Event, error) {
	return tx.Event.Create().SetContents(req.Contents).SetWriter(req.Writer).Save(ctx)
}

func NewEventRepository(client ent.Client) EventRepository {
	return &eventRepositoryImpl{Client: client}
}
