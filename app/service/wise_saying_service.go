package service

import (
	"context"
	"geulSsi/app/dto/custom"
	wisesayingDto "geulSsi/app/dto/wisesaying"
	"geulSsi/app/repository"
	"geulSsi/bcrypt"
	"geulSsi/ent"
)

type WiseSayingService interface {
	AddWiseSaying(ctx context.Context, addWiseSayingRequest wisesayingDto.AddWiseSayingRequest) (*custom.Response, *custom.Error) // 명언 추가
}

type wiseSayingService struct {
	userRepository       repository.UserRepository
	wiseSayingRepository repository.WiseSayingRepository
}

func (s wiseSayingService) AddWiseSaying(ctx context.Context, addWiseSayingRequest wisesayingDto.AddWiseSayingRequest) (*custom.Response, *custom.Error) {

	// 닉네임으로 유저 정보를 가져온다.
	user, err := s.userRepository.FindByNickname(ctx, &ent.User{Nickname: addWiseSayingRequest.Nickname})
	if err != nil {
		return nil, custom.InternalServerError
	}

	switch user.Nickname {
	case "root1":
		break
	case "root2":
		break
	case "root3":
		break
	default:
		return nil, custom.Unauthorized
	}

	// 비밀번호가 맞는지 체크한다.
	if err := bcrypt.Compare(user.Password, addWiseSayingRequest.Password); err != nil {
		return nil, custom.PasswdNotMatch
	}

	// 트랜잭션을 생성한다.
	tx, customError := custom.TxGenerate(ctx)
	if customError != nil {
		return nil, customError
	}

	// 명언을 추가한다.
	_, err = s.wiseSayingRepository.CreateWiseSaying(ctx, &ent.WiseSaying{
		CreatedBy:  addWiseSayingRequest.Nickname,
		WiseSaying: addWiseSayingRequest.WiseSaying,
	}, tx)
	if err != nil {
		custom.TxRollback(tx)
		return nil, custom.InternalServerError
	}

	// 커밋한다.
	if customError = custom.TxCommit(tx); customError != nil {
		return nil, customError
	}

	// 리턴
	return custom.NewSuccess(), nil

}

func NewWiseSayingService(userRepository repository.UserRepository, wiseSayingRepository repository.WiseSayingRepository) WiseSayingService {
	return &wiseSayingService{userRepository, wiseSayingRepository}
}
