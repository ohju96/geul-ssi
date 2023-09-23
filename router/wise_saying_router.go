package router

import (
	"geulSsi/app/controller"
	"geulSsi/app/repository"
	"geulSsi/app/service"
	"geulSsi/config/db"
	"github.com/gin-gonic/gin"
)

func InitWiseSayingRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1/wise-sayings")
	{
		wiseSaying := wiseSayingDependency()
		v1.GET("/events", wiseSaying.PushWiseSaying)
		v1.POST("", wiseSaying.AddWiseSaying)

	}
}

func wiseSayingDependency() controller.WiseSayingController {
	return controller.NewWiseSayingController(
		service.NewWiseSayingService(
			repository.NewUserRepository(*db.MySQL),
			repository.NewWiseSayingRepository(*db.MySQL),
		),
	)
}
