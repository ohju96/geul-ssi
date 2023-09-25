package repository

import (
	"context"
	"geulSsi/ent"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, req *ent.Event, tx *ent.Tx) (*ent.Event, error) // 생성
}

type eventRepositoryImpl struct {
	ent.Client
}

func (r eventRepositoryImpl) CreateEvent(ctx context.Context, req *ent.Event, tx *ent.Tx) (*ent.Event, error) {
	return tx.Event.Create().SetContents(req.Contents).SetWriter(req.Writer).Save(ctx)
}

func NewEventRepository(client ent.Client) EventRepository {
	return &eventRepositoryImpl{Client: client}
}
