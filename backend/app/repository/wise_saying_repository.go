package repository

import (
	"context"
	"geulSsi/ent"
)

type WiseSayingRepository interface {
	CreateWiseSaying(ctx context.Context, req *ent.WiseSaying, tx *ent.Tx) (*ent.WiseSaying, error) // 명언 추가
	CountAll(ctx context.Context) (int, error)
	FindAllWiseSaying(ctx context.Context) ([]string, error)
}

type wiseSayingRepository struct {
	*ent.Client
}

func (r wiseSayingRepository) FindAllWiseSaying(ctx context.Context) ([]string, error) {
	return r.WiseSaying.Query().Select("wise_saying").Strings(ctx)
}

func (r wiseSayingRepository) CountAll(ctx context.Context) (int, error) {
	return r.WiseSaying.Query().Count(ctx)
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
