package controller

import (
	"fmt"
	wiseSayingDto "geulSsi/app/dto/wisesaying"
	"geulSsi/app/service"
	"geulSsi/validator"
	"github.com/gin-gonic/gin"
	"time"
)

type WiseSayingController interface {
	PushWiseSaying(g *gin.Context) // SSE 전송
	AddWiseSaying(g *gin.Context)  // 명언 추가

}

type wiseSayingController struct {
	wiseSayingService service.WiseSayingService
}

// AddWiseSaying godoc
// @Summary 명언 추가
// @Description 관리자 계정으로 명언을 추가합니다.
// @Tags WiseSayings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param wiseSayingDto.wiseSaying body wiseSayingDto.AddWiseSayingRequest true "명언 추가 요청 객체"
// @Success 200 {object} custom.Response
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /wise-sayings [post]
func (c wiseSayingController) AddWiseSaying(g *gin.Context) {

	var addWiseSayingRequest wiseSayingDto.AddWiseSayingRequest
	if err := validator.BinderAndValidator(g, &addWiseSayingRequest); err != nil {
		g.JSON(err.StatusCode, err)
		return
	}

	ctx := g.Request.Context()

	res, err := c.wiseSayingService.AddWiseSaying(ctx, addWiseSayingRequest)
	if err != nil {
		g.JSON(err.StatusCode, err)
		return
	}

	g.JSON(res.StatusCode, res)

}

// PushWiseSaying godoc
// @Summary 명언 가져오기
// @Description 명언을 가져옵니다.
// @Tags WiseSayings
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /wise-sayings/events [get]
func (c wiseSayingController) PushWiseSaying(g *gin.Context) {

	g.Header("Content-Type", "text/event-stream")
	g.Header("Cache-Control", "no-cache")
	g.Header("Connection", "keep-alive")

	for {
		wiseSayingMsg, err := c.wiseSayingService.GetWiseSayingChannel()
		if err != nil {
			g.JSON(err.StatusCode, err)
			return
		}
		fmt.Println("wiseSayingMsg: ", *wiseSayingMsg)
		g.SSEvent("message", *wiseSayingMsg)
		g.Writer.Flush()
		time.Sleep(3 * time.Second)
	}
}

func NewWiseSayingController(wiseSayingService service.WiseSayingService) WiseSayingController {
	return &wiseSayingController{wiseSayingService: wiseSayingService}
}
