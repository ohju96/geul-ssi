package service

import (
	"context"
	"geulSsi/app/dto/custom"
	wisesayingDto "geulSsi/app/dto/wisesaying"
	"geulSsi/app/repository"
	"geulSsi/bcrypt"
	"geulSsi/ent"
	"log"
	"sync"
)

var (
	wiseSayingMsg = make(chan string, 100)
	msgMutex      sync.Mutex
	wg            sync.WaitGroup
)

type WiseSayingService interface {
	AddWiseSaying(ctx context.Context, addWiseSayingRequest wisesayingDto.AddWiseSayingRequest) (*custom.Response, *custom.Error) // 명언 추가
	RegisterWiseSayingChannel(ctx context.Context)
	GetWiseSayingChannel() (*string, *custom.Error)
}

type wiseSayingService struct {
	userRepository       repository.UserRepository
	wiseSayingRepository repository.WiseSayingRepository
}

func (s wiseSayingService) GetWiseSayingChannel() (*string, *custom.Error) {

	defaultMsg := "채널에 등록된 명언이 없습니다."
	msgMutex.Lock()
	defer msgMutex.Unlock()

	select {
	case wiseSaying, ok := <-wiseSayingMsg:
		if !ok {
			return nil, custom.InternalServerError
		}

		if len(wiseSaying) == 0 {
			return &defaultMsg, nil
		}

		defaultMsg = wiseSaying

		go func() {
			wiseSayingMsg <- defaultMsg
		}()

		return &wiseSaying, nil
	default:
		return &defaultMsg, nil
	}
}

func (s wiseSayingService) RegisterWiseSayingChannel(ctx context.Context) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sayingList, err := s.wiseSayingRepository.FindAllWiseSaying(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		for _, saying := range sayingList {
			msgMutex.Lock()
			wiseSayingMsg <- saying
			msgMutex.Unlock()
		}
	}()
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

	// 채널에 추가한다.
	go func() {
		if len(wiseSayingMsg) >= 3 {
			<-wiseSayingMsg
		}

		msgMutex.Lock()
		wiseSayingMsg <- addWiseSayingRequest.WiseSaying
		defer msgMutex.Unlock()
	}()
	// 리턴
	return custom.NewSuccess(), nil

}

func NewWiseSayingService(userRepository repository.UserRepository, wiseSayingRepository repository.WiseSayingRepository) WiseSayingService {
	return &wiseSayingService{userRepository, wiseSayingRepository}
}
