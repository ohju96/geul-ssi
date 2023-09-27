package controller

import (
	eventDto "geulSsi/app/dto/event"
	"geulSsi/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EventController interface {
	GetEvent(g *gin.Context)
	CreateEvent(g *gin.Context)
	GetEventList(g *gin.Context)
	CreateHeart(g *gin.Context)
}

type eventControllerImpl struct {
	eventChannel chan eventDto.Event
	eventService service.EventService
}

// CreateHeart godoc
// @Summary 이벤트에 좋아요를 누릅니다.
// @Description 이벤트에 좋아요를 누릅니다.
// @Tags Events
// @Accept json
// @Produce json
// @Param eventDto.CreateHeartReq body eventDto.CreateHeartReq true "좋아요 생성 요청 객체"
// @Success 200 {object} eventDto.CreateHeartResponse
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /events/heart [post]
func (c eventControllerImpl) CreateHeart(g *gin.Context) {

	var req eventDto.CreateHeartReq

	// 클라이언트로부터 JSON 데이터를 파싱
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := g.Request.Context()

	res, customErr := c.eventService.CreateHeart(ctx, &req)
	if customErr != nil {
		g.JSON(customErr.StatusCode, customErr)
		return
	}

	g.JSON(res.StatusCode, res)
	return
}

// GetEventList godoc
// @Summary 자신이 생성한 이벤트를 조회합니다.
// @Description 자신이 생성한 이벤트를 조회합니다.
// @Tags Events
// @Accept json
// @Produce json
// @Param eventDto.GetEventReq body eventDto.GetEventReq true "이벤트 생성 요청 객체"
// @Success 200 {object} eventDto.GetEventListResponse
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /events/list [post]
func (c eventControllerImpl) GetEventList(g *gin.Context) {

	var req eventDto.GetEventReq

	// 클라이언트로부터 JSON 데이터를 파싱
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := g.Request.Context()

	res, customErr := c.eventService.GetEventList(ctx, &req)
	if customErr != nil {
		g.JSON(customErr.StatusCode, customErr)
		return
	}

	g.JSON(res.StatusCode, res)
	return
}

// CreateEvent godoc
// @Summary 이벤트 생성
// @Description 이벤트를 생성합니다.
// @Tags Events
// @Accept json
// @Produce json
// @Param eventDto.Event body eventDto.Event true "이벤트 생성 요청 객체"
// @Success 200 {object} custom.Response
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /events [post]
func (c eventControllerImpl) CreateEvent(g *gin.Context) {

	eventDto.Dummy()
	var event eventDto.Event

	// 클라이언트로부터 JSON 데이터를 파싱
	if err := g.ShouldBindJSON(&event); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 쿠키를 검증하여 이벤트 생성 제한 확인
	cookie, err := g.Request.Cookie("eventLimit")
	if err != nil {
		// 쿠키가 없으면 새로 생성하고 만료 시간 설정
		expiration := time.Now().Add(10 * time.Second) // 10초
		cookie = &http.Cookie{
			Name:    "eventLimit",
			Value:   "created",
			Expires: expiration,
		}
		http.SetCookie(g.Writer, cookie)
	} else {
		// 쿠키가 이미 존재하면 이벤트 생성 제한
		g.JSON(http.StatusForbidden, "10초 후에 이벤트를 생성할 수 있습니다.")
		return
	}

	ctx := g.Request.Context()

	res, customErr := c.eventService.CreateEvent(ctx, &event)
	if customErr != nil {
		g.JSON(customErr.StatusCode, customErr)
		return
	}

	event.EventId = res.Payload.EventID

	go func() {
		c.eventChannel <- event
	}()

	// 이벤트 생성 완료 메시지 반환
	g.JSON(http.StatusOK, res)
	return
}

// GetEvent godoc
// @Summary 이벤트 조회
// @Description 이벤트를 조회합니다.
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} custom.Response
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /events [get]
func (c eventControllerImpl) GetEvent(g *gin.Context) {

	select {

	// 채널에서 이벤트 데이터를 받아옴
	case event, ok := <-c.eventChannel:

		if !ok {
			// 채널이 닫혔을 때의 처리 (예: 종료 또는 에러 처리)
			g.JSON(http.StatusInternalServerError, "서버 오류 발생")
			return
		}

		if event.Content == "" {
			g.JSON(http.StatusOK, "회원 가입 후 메시지를 등록해주세요!")
			return
		}

		// 받아온 이벤트 데이터를 클라이언트에 반환
		g.JSON(http.StatusOK, gin.H{"user": event.User, "content": event.Content, "eventId": event.EventId})
		return

	default:
		g.JSON(http.StatusOK, "회원 가입 후 메시지를 등록해주세요!")
		return
	}
}

func NewEventController(eventService service.EventService) EventController {

	// 이벤트 데이터를 저장할 채널 생성
	eventChannel := make(chan eventDto.Event, 1) // 채널 버퍼 크기 조절 가능

	return &eventControllerImpl{
		eventService: eventService,
		eventChannel: eventChannel,
	}
}
