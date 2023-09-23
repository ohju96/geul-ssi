package router

import (
	"geulSsi/app/controller"
	"github.com/gin-gonic/gin"
)

func InitWiseSayingRouter(g *gin.Engine) {
	v1 := g.Group("/api/v1/wise-sayings")
	{
		wiseSaying := wiseSayingDependency()
		v1.GET("/events", wiseSaying.PushWiseSaying)

	}
}

func wiseSayingDependency() controller.WiseSayingController {
	return controller.NewWiseSayingController()
}
