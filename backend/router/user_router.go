package router

import (
	"geulSsi/app/controller"
	"geulSsi/app/repository"
	"geulSsi/app/service"
	"geulSsi/config"
	"geulSsi/config/db"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.Engine, toml *config.Config) {
	v1 := g.Group("/api/v1/users")
	{
		userController := userDependency()
		v1.POST("", userController.CreateUser)                           // 회원가입
		v1.PUT("", userController.ChangePasswordByNicknameAndSecretCode) // SecretCode로 비밀번호 변경
	}
}

func userDependency() controller.UserController {
	return controller.NewUserController(
		service.NewUserService(
			repository.NewUserRepository(*db.MySQL),
		),
	)
}
