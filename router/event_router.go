package router

import (
	"geulSsi/app/controller"
	"geulSsi/config"
	"github.com/gin-gonic/gin"
)

func InitEventRouter(g *gin.Engine, toml *config.Config) {
	v1 := g.Group("/api/v1/events")
	{
		eventController := eventDependency()
		v1.POST("", eventController.CreateEvent) // 이벤트 생성
		v1.GET("", eventController.GetEvent)     // 이벤트 조회
	}
}

func eventDependency() controller.EventController {
	return controller.NewEventController()
}
