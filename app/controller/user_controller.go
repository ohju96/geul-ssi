package controller

import (
	userDto "geulSsi/app/dto/user"
	"geulSsi/app/service"
	"geulSsi/validator"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(g *gin.Context) // 회원가입
}

type userControllerImpl struct {
	userService service.UserService
}

// CreateUser godoc
// @Summary 회원가입
// @Description 회원가입을 하면 SecretCode를 발급합니다. SecretCode로 비밀번호 변경이나 회원 찾기에 사용되나 재발급은 되지 않으니 보관 안내가 필요합니다.
// @Tags Users
// @Accept json
// @Produce json
// @Param userDto.CreateUserRequest body userDto.CreateUserRequest true "회원가입 요청 객체"
// @Success 200 {object} userDto.CreateUserResponse
// @Failure 400 {object} custom.Fail400GetResponse
// @Failure 500 {object} custom.Fail500GetResponse
// @Router /api/v1/users [post]
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
