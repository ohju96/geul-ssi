package repository

import (
	"context"
	"geulSsi/ent"
	"geulSsi/ent/heart"
)

type HeartRepository interface {
	CountHeartCountByEventId(ctx context.Context, req *ent.Heart) (int, error)
	GetHeartByEventIdAndNickName(ctx context.Context, req *ent.Heart) (*ent.Heart, error)
	CreateHeart(ctx context.Context, req *ent.Heart) (*ent.Heart, error)
	DeleteHeart(ctx context.Context, req *ent.Heart) (int, error)
}

type heartRepositoryImpl struct {
	ent.Client
}

func (r heartRepositoryImpl) DeleteHeart(ctx context.Context, req *ent.Heart) (int, error) {
	return r.Heart.Update().Where(
		heart.And(heart.EventIDEQ(req.EventID),
			heart.WriterEQ(req.Writer))).
		SetIsHeart(false).
		Save(ctx)
}

func (r heartRepositoryImpl) CreateHeart(ctx context.Context, req *ent.Heart) (*ent.Heart, error) {
	return r.Heart.Create().
		SetEventID(req.EventID).
		SetWriter(req.Writer).
		SetIsHeart(true).
		Save(ctx)
}

func (r heartRepositoryImpl) GetHeartByEventIdAndNickName(ctx context.Context, req *ent.Heart) (*ent.Heart, error) {
	return r.Heart.Query().Where(heart.And(heart.EventIDEQ(req.EventID), heart.WriterEQ(req.Writer))).First(ctx)
}

func (r heartRepositoryImpl) CountHeartCountByEventId(ctx context.Context, req *ent.Heart) (int, error) {
	return r.Heart.Query().Where(heart.And(heart.EventIDEQ(req.EventID))).Count(ctx)
}

func NewHeartRepository(client ent.Client) HeartRepository {
	return &heartRepositoryImpl{Client: client}
}
