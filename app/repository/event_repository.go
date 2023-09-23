package repository

import "geulSsi/ent"

type EventRepository interface {
}

type eventRepositoryImpl struct {
	client ent.Client
}

func NewEventRepository(client ent.Client) EventRepository {
	return &eventRepositoryImpl{client: client}
}
