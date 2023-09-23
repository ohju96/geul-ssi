package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

type WiseSayingController interface {
	PushWiseSaying(g *gin.Context) // SSE 전송
}

type wiseSayingController struct {
}

func (c wiseSayingController) PushWiseSaying(g *gin.Context) {

	g.Header("Content-Type", "text/event-stream")
	g.Header("Cache-Control", "no-cache")
	g.Header("Connection", "keep-alive")

	for {
		message := "현재 시간 : " + time.Now().Format("2006-01-02 15:04:05")
		g.SSEvent("message", message)
		g.Writer.Flush()
		time.Sleep(3 * time.Second)
	}
}

func NewWiseSayingController() WiseSayingController {
	return &wiseSayingController{}
}
