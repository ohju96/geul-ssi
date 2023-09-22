package validator

import (
	"geulSsi/app/dto/custom"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidate() {
	validate = validator.New()
}

// 유효성 검사 메서드
func Validator(req interface{}) *custom.Error {

	if err := validate.Struct(req); err != nil {
		return custom.RequestValidateFail
	}

	return nil
}

// 요청 객체 바인딩 메서드
func Binder(g *gin.Context, req interface{}) *custom.Error {

	if err := g.ShouldBind(req); err != nil {
		return custom.RequestBindFail
	}

	return nil
}

// 유효성 검사 및 요청 객체 바인딩 메서드
func BinderAndValidator(g *gin.Context, req interface{}) *custom.Error {

	if err := g.ShouldBind(req); err != nil {
		return custom.RequestBindFail
	}

	if err := validate.Struct(req); err != nil {
		return custom.RequestValidateFail
	}

	return nil
}
