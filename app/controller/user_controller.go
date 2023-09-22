package controller

import (
	userDto "geulSsi/app/dto/user"
	"geulSsi/app/service"
	"geulSsi/validator"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(g *gin.Context)
}

type userControllerImpl struct {
	userService service.UserService
}

func (c userControllerImpl) CreateUser(g *gin.Context) {

	var createUserRequest userDto.CreateUserRequest
	if err := validator.BinderAndValidator(g, &createUserRequest); err != nil {
		g.JSON(err.StatusCode, err)
		return
	}

	ctx := g.Request.Context()

	res, customError := c.userService.CreateUser(ctx, &createUserRequest)
	if customError != nil {
		g.JSON(customError.StatusCode, customError)
		return
	}

	g.JSON(res.StatusCode, res)
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}
