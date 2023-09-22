package service

import (
	"context"
	"geulSsi/app/dto/custom"
	userDto "geulSsi/app/dto/user"
	"geulSsi/app/repository"
	"geulSsi/bcrypt"
	"geulSsi/ent"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, createUserRequest *userDto.CreateUserRequest) (*userDto.CreateUserResponse, *custom.Error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (s userServiceImpl) CreateUser(ctx context.Context, createUserRequest *userDto.CreateUserRequest) (*userDto.CreateUserResponse, *custom.Error) {

	existNickname, err := s.userRepository.ExistByUserNickname(ctx, &ent.User{Nickname: createUserRequest.Nickname})
	if err != nil {
		return nil, custom.InternalServerError
	}

	if existNickname {
		return nil, custom.AlreadyExist
	}

	secretCode, err := uuid.NewUUID()
	if err != nil {
		return nil, custom.UUIDGenerateFail
	}

	existSecretCode, err := s.userRepository.ExistBySecretCode(ctx, &ent.User{SecretCode: secretCode.String()})
	if err != nil {
		return nil, custom.InternalServerError
	}

	if existSecretCode {
		return nil, custom.AlreadyExist
	}

	password, err := bcrypt.Generate(createUserRequest.Password)
	if err != nil {
		return nil, custom.BcryptGenerateFail
	}

	tx, customError := custom.TxGenerate(ctx)
	if customError != nil {
		return nil, customError
	}

	user, err := s.userRepository.CreateUser(ctx, &ent.User{
		Nickname:   createUserRequest.Nickname,
		SecretCode: secretCode.String(),
		Password:   password,
	}, tx)
	if err != nil {
		custom.TxRollback(tx)
		return nil, custom.InternalServerError
	}

	if customError = custom.TxCommit(tx); customError != nil {
		return nil, customError
	}

	return &userDto.CreateUserResponse{
		Response: custom.NewSuccess(),
		Payload: userDto.CreateUserPayload{
			UserID:     user.ID,
			SecretCode: user.SecretCode,
		},
	}, nil

}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}
