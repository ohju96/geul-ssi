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
	ChangePasswordByNicknameAndSecretCode(ctx context.Context, changePasswordRequest *userDto.ChangePasswordRequest) (*custom.Response, *custom.Error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (s userServiceImpl) ChangePasswordByNicknameAndSecretCode(ctx context.Context, changePasswordRequest *userDto.ChangePasswordRequest) (*custom.Response, *custom.Error) {

	// 닉네임이랑 시크릿코드랑 일치하는 유저가 있는지 확인
	existUser, err := s.userRepository.ExistBySecretCodeAndNickname(ctx, &ent.User{
		SecretCode: changePasswordRequest.SecretCode,
		Nickname:   changePasswordRequest.Nickname})
	if err != nil {
		return nil, custom.InternalServerError
	}

	// 유저가 없다면 에러
	if !existUser {
		return nil, custom.NotFound
	}

	// 비밀번호 암호화
	password, err := bcrypt.Generate(changePasswordRequest.Password)
	if err != nil {
		return nil, custom.BcryptGenerateFail
	}

	// 트랜잭션 생성
	tx, customError := custom.TxGenerate(ctx)
	if customError != nil {
		return nil, customError
	}

	// 비밀번호 변경
	if err := s.userRepository.UpdatePasswordByNicknameAndSecretCode(ctx, &ent.User{
		SecretCode: changePasswordRequest.SecretCode,
		Nickname:   changePasswordRequest.Nickname,
		Password:   password}, tx); err != nil {
		custom.TxRollback(tx)
	}

	// 트랜잭션 커밋
	if customError = custom.TxCommit(tx); customError != nil {
		return nil, customError
	}

	return custom.NewSuccess(), nil

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
		Response: *custom.NewSuccess(),
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
