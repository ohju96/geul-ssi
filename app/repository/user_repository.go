package repository

import (
	"context"
	"geulSsi/ent"
	"geulSsi/ent/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *ent.User, tx *ent.Tx) (*ent.User, error) // 생성
	FindByUserNickname(ctx context.Context, req *ent.User) (*ent.User, error)     // 닉네임으로 조회
	FindBySecretCode(ctx context.Context, req *ent.User) (*ent.User, error)       // 시크릿코드로 조회
	ExistByUserNickname(ctx context.Context, req *ent.User) (bool, error)         // 닉네임으로 존재여부
	ExistBySecretCode(ctx context.Context, req *ent.User) (bool, error)           // 시크릿코드로 존재여부
}

type userRepositoryImpl struct {
	ent.Client
}

func (r userRepositoryImpl) ExistByUserNickname(ctx context.Context, req *ent.User) (bool, error) {
	return r.User.Query().
		Where(user.NicknameEQ(req.Nickname)).
		Exist(ctx)
}

func (r userRepositoryImpl) ExistBySecretCode(ctx context.Context, req *ent.User) (bool, error) {
	return r.User.Query().
		Where(user.SecretCodeEQ(req.SecretCode)).
		Exist(ctx)
}

func (r userRepositoryImpl) CreateUser(ctx context.Context, req *ent.User, tx *ent.Tx) (*ent.User, error) {
	return tx.User.Create().
		SetNickname(req.Nickname).
		SetSecretCode(req.SecretCode).
		SetPassword(req.Password).
		Save(ctx)
}

func (r userRepositoryImpl) FindByUserNickname(ctx context.Context, req *ent.User) (*ent.User, error) {
	return r.User.Query().
		Where(user.NicknameEQ(req.Nickname)).
		Only(ctx)
}

func (r userRepositoryImpl) FindBySecretCode(ctx context.Context, req *ent.User) (*ent.User, error) {
	return r.User.Query().
		Where(user.SecretCodeEQ(req.SecretCode)).
		Only(ctx)
}

func NewUserRepository(client ent.Client) UserRepository {
	return &userRepositoryImpl{
		Client: client,
	}
}
