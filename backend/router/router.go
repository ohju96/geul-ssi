package router

import (
	"geulSsi/config"
	"github.com/gin-gonic/gin"
)

func MainRouter(g *gin.Engine, toml *config.Config) {
	InitUserRouter(g, toml)
	InitSwagRouter(g)
	InitWiseSayingRouter(g)
	InitEventRouter(g, toml)
}
