package router

import (
	"geulSsi/app/controller"
	"geulSsi/app/repository"
	"geulSsi/app/service"
	"geulSsi/config"
	"geulSsi/config/db"
	"github.com/gin-gonic/gin"
)

func InitEventRouter(g *gin.Engine, toml *config.Config) {
	v1 := g.Group("/api/v1/events")
	{
		eventController := eventDependency()
		v1.POST("", eventController.CreateEvent)       // 이벤트 생성
		v1.GET("", eventController.GetEvent)           // 이벤트 조회
		v1.POST("/list", eventController.GetEventList) // 이벤트 리스트 조회
		v1.POST("/heart", eventController.CreateHeart) // 이벤트 좋아요 생성
	}
}

func eventDependency() controller.EventController {
	return controller.NewEventController(
		service.NewEventService(
			repository.NewEventRepository(*db.MySQL),
			repository.NewUserRepository(*db.MySQL),
			repository.NewHeartRepository(*db.MySQL),
		),
	)
}
