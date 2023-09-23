package repository

import (
	"context"
	"geulSsi/ent"
)

type WiseSayingRepository interface {
	CreateWiseSaying(ctx context.Context, req *ent.WiseSaying, tx *ent.Tx) (*ent.WiseSaying, error) // 명언 추가

}

type wiseSayingRepository struct {
	*ent.Client
}

func (r wiseSayingRepository) CreateWiseSaying(ctx context.Context, req *ent.WiseSaying, tx *ent.Tx) (*ent.WiseSaying, error) {
	return tx.WiseSaying.Create().
		SetCreatedBy(req.CreatedBy).
		SetWiseSaying(req.WiseSaying).
		Save(ctx)
}

func NewWiseSayingRepository(client ent.Client) WiseSayingRepository {
	return &wiseSayingRepository{&client}
}
